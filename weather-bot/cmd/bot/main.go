package main

import (
	"log"
	"weather-bot/internal/bot"
	"weather-bot/internal/config"
	"weather-bot/internal/localization"
	"weather-bot/internal/weather"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize dependencies
	weatherClient := weather.NewOpenWeatherClient(cfg.OpenWeatherKey)
	translator := localization.NewPersianWeatherTranslator()
	weatherService := weather.NewWeatherService(weatherClient, translator)

	// Create bot handler
	handler := bot.NewWeatherMessageHandler(nil, weatherService)

	// Initialize bot
	telegramBot, err := bot.NewTelegramBot(cfg.TelegramToken, handler)
	if err != nil {
		log.Panic(err)
	}

	// Inject bot instance into handler
	handler.SetBot(telegramBot.GetBot())

	// Start the bot
	telegramBot.Start()
}