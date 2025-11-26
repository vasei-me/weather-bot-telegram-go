package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken  string
	OpenWeatherKey string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("⚠️ فایل .env پیدا نشد! مطمئن شو در همان پوشه باشه")
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	openWeatherKey := os.Getenv("OPENWEATHER_KEY")

	if telegramToken == "" || openWeatherKey == "" {
		log.Fatal("⚠️ توکن تلگرام یا کلید OpenWeather خالیه! فایل .env رو چک کن")
	}

	return &Config{
		TelegramToken:  telegramToken,
		OpenWeatherKey: openWeatherKey,
	}
}