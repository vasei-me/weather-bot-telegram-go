package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
    Main struct {
        Temp      float64 `json:"temp"`
        FeelsLike float64 `json:"feels_like"`
        Humidity  int     `json:"humidity"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
        Main        string `json:"main"`
    } `json:"weather"`
    Name string `json:"name"`
    Sys  struct {
        Country string `json:"country"`
    } `json:"sys"`
}

func main() {
    // بارگذاری فایل .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("⚠️ فایل .env پیدا نشد! مطمئن شو در همان پوشه باشه")
    }

    TelegramToken := os.Getenv("TELEGRAM_TOKEN")
    OpenWeatherKey := os.Getenv("OPENWEATHER_KEY")

    if TelegramToken == "" || OpenWeatherKey == "" {
        log.Fatal("⚠️ توکن تلگرام یا کلید OpenWeather خالیه! فایل .env رو چک کن")
    }

    bot, err := tgbotapi.NewBotAPI(TelegramToken)
    if err != nil {
        log.Panic(err)
    }

    log.Printf("🤖 ربات با موفقیت راه‌اندازی شد: @%s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        // وقتی کاربر لوکیشن می‌فرسته
        if update.Message.Location != nil {
            lat := update.Message.Location.Latitude
            lon := update.Message.Location.Longitude

            weather, cityName, err := getWeather(lat, lon, OpenWeatherKey)
            if err != nil {
                errorMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "❌ خطا در دریافت آب و هوا\nاینترنت یا VPN رو چک کن و دوباره امتحان کن!")
                if _, err := bot.Send(errorMsg); err != nil {
                    log.Printf("Error sending error message to user: %v", err)
                }
                continue
            }

            // محافظت از آرایه خالی
            description := "نامشخص"
            if len(weather.Weather) > 0 {
                desc := strings.ToLower(weather.Weather[0].Description)
                description = toPersianWeather(desc)

                // اگر ترجمه نشد، از گروه اصلی (main) استفاده کن
                if description == "نامشخص" && weather.Weather[0].Main != "" {
                    description = toPersianWeather(strings.ToLower(weather.Weather[0].Main))
                }
            }

            tempC := weather.Main.Temp - 273.15
            feelsC := weather.Main.FeelsLike - 273.15

            locationName := cityName
            if locationName == "" {
                locationName = "مکان شما"
            }
            if weather.Sys.Country != "" {
                locationName = fmt.Sprintf("%s، %s", locationName, weather.Sys.Country)
            }

            text := fmt.Sprintf(`🌍 آب و هوای *%s*

🌡️ دما: *%.1f°C*
🤒 حس واقعی: *%.1f°C*
💧 رطوبت: *%d%%*
☁️ وضعیت: %s`,
                locationName,
                tempC,
                feelsC,
                weather.Main.Humidity,
                description,
            )

            msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
            msg.ParseMode = "Markdown"
            if _, err := bot.Send(msg); err != nil {
                log.Printf("Error sending weather message to user: %v", err)
            }

            // ارسال دوباره لوکیشن (قشنگ‌تره)
            locMsg := tgbotapi.NewLocation(update.Message.Chat.ID, lat, lon)
            if _, err := bot.Send(locMsg); err != nil {
                log.Printf("Error sending location message to user: %v", err)
            }
            continue
        }

        // پیام خوش‌آمدگویی + دکمه لوکیشن
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
        if _, err := bot.Send(msg); err != nil {
            log.Printf("Error sending welcome message to user: %v", err)
        }
    }
}

// ترجمه فارسی کامل (بدون هیچ کلید تکراری!)
func toPersianWeather(text string) string {
    m := map[string]string{
        // انگلیسی
        "clear sky":             "آسمان صاف",
        "few clouds":            "کمی ابری",
        "scattered clouds":      "ابری پراکنده",
        "broken clouds":         "ابری",
        "overcast clouds":       "کاملاً ابری",
        "light rain":            "باران سبک",
        "moderate rain":         "باران",
        "heavy intensity rain":  "باران شدید",
        "shower rain":           "رگبار",
        "thunderstorm":          "رعد و برق",
        "snow":                  "برف",
        "mist":                  "مه",
        "fog":                   "غبار مه",
        "haze":                  "غبار",
        "drizzle":               "نم‌نم باران",
        "smoke":                 "دود",
        "dust":                  "گرد و غبار",

        // فارسی (وقتی lang=fa استفاده می‌شه)
        "آسمان صاف":              "آسمان صاف",
        "غیوم قليلة":             "کمی ابری",
        "غیوم متفرقة":            "ابری پراکنده",
        "غیوم مكسرة":             "ابری",
        "غیوم كثيفة":             "کاملاً ابری",
        "مطر خفیف":              "باران سبک",
        "مطر":                   "باران",
        "مطر غزير":              "باران شدید",
        "زخات مطر":              "رگبار",
        "عاصفة رعدية":            "رعد و برق",
        "ثلج":                   "برف",
        "ضباب":                  "مه",
        "ضباب خفیف":             "غبار مه",

        // گروه‌های اصلی (fallback)
        "clear":                 "آسمان صاف",
        "clouds":                "ابری",
        "rain":                  "باران",
    }

    if val, ok := m[text]; ok {
        return val
    }
    return "نامشخص"
}

func getWeather(lat, lon float64, apiKey string) (*WeatherResponse, string, error) {
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&lang=fa", lat, lon, apiKey)

    client := &http.Client{Timeout: 12 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        return nil, "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, "", err
    }

    var weather WeatherResponse
    if err := json.Unmarshal(body, &weather); err != nil {
        return nil, "", err
    }

    cityName := weather.Name
    if cityName == "" {
        cityName = "مکان شما"
    }

    return &weather, cityName, nil
}