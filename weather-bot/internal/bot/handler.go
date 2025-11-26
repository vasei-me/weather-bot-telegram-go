package bot

import (
	"fmt"

	"weather-bot/internal/weather"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler interface {
	HandleLocation(update tgbotapi.Update) error
	HandleMessage(update tgbotapi.Update) error
	SetBot(bot *tgbotapi.BotAPI)
}

type WeatherMessageHandler struct {
	bot           *tgbotapi.BotAPI
	weatherService *weather.WeatherService
}

func NewWeatherMessageHandler(bot *tgbotapi.BotAPI, weatherService *weather.WeatherService) *WeatherMessageHandler {
	return &WeatherMessageHandler{
		bot:           bot,
		weatherService: weatherService,
	}
}

func (h *WeatherMessageHandler) SetBot(bot *tgbotapi.BotAPI) {
	h.bot = bot
}

func (h *WeatherMessageHandler) HandleLocation(update tgbotapi.Update) error {
	lat := update.Message.Location.Latitude
	lon := update.Message.Location.Longitude

	weatherInfo, err := h.weatherService.GetWeatherInfo(lat, lon)
	if err != nil {
		errorMsg := tgbotapi.NewMessage(
			update.Message.Chat.ID, 
			"❌ خطا در دریافت آب و هوا\nاینترنت یا VPN رو چک کن و دوباره امتحان کن!",
		)
		_, err := h.bot.Send(errorMsg)
		return err
	}

	locationName := weatherInfo.Location
	if weatherInfo.Country != "" {
		locationName = fmt.Sprintf("%s، %s", locationName, weatherInfo.Country)
	}

	text := fmt.Sprintf(`🌍 آب و هوای *%s*

🌡️ دما: *%.1f°C*
🤒 حس واقعی: *%.1f°C*
💧 رطوبت: *%d%%*
☁️ وضعیت: %s`,
		locationName,
		weatherInfo.Temperature,
		weatherInfo.FeelsLike,
		weatherInfo.Humidity,
		weatherInfo.Description,
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "Markdown"
	if _, err := h.bot.Send(msg); err != nil {
		return err
	}

	// ارسال دوباره لوکیشن
	locMsg := tgbotapi.NewLocation(update.Message.Chat.ID, lat, lon)
	_, err = h.bot.Send(locMsg)
	return err
}

func (h *WeatherMessageHandler) HandleMessage(update tgbotapi.Update) error {
	welcome := `سلام دوست عزیز! 👋

من ربات آب و هوا هستم 🌤️
لوکیشن خودت رو برام بفرست تا بگم الان هوا چطوره!`

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonLocation("📍 ارسال موقعیت مکانی"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcome)
	msg.ReplyMarkup = keyboard
	_, err := h.bot.Send(msg)
	return err
}