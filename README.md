# 🌤️ Telegram Weather Bot (Go)

A simple and efficient Telegram bot built with Go that provides detailed weather information in Persian. The bot fetches real-time weather data based on the user's location using the Telegram Bot API and the OpenWeatherMap API.

## ✨ Features

- 📍 **Location-Based**: Gets weather data using the user's live location sent via Telegram.
- 🌡️ **Detailed Info**: Displays current temperature, "feels like" temperature, and humidity.
- ☁️ **Persian Descriptions**: Translates weather conditions (e.g., "clear sky", "rain") into user-friendly Persian.
- 🌍 **Location Display**: Shows the city and country name for the provided coordinates.
- 🤖 **User-Friendly Interface**: Welcomes users with a clear message and an easy-to-use location button.
- 🛡️ **Error Handling**: Gracefully handles API errors and network issues with informative messages.

## 📸 Preview

(You can add a screenshot or GIF of the bot in action here)

## 🚀 Getting Started

Follow these steps to set up and run the bot locally.

### 1. Prerequisites

- **Go** (version 1.18 or newer)
- A **Bot Token** from Telegram's [@BotFather](https://t.me/botfather)
- An **API Key** from [OpenWeatherMap](https://openweathermap.org/api)

### 2. Get a Telegram Bot Token

1. Start a chat with [@BotFather](https://t.me/botfather) on Telegram.
2. Send the `/newbot` command.
3. Follow the instructions to choose a name and a username for your bot.
4. BotFather will provide you with a unique token. Copy it.

### 3. Get an OpenWeatherMap API Key

1. Sign up on the [OpenWeatherMap](https://home.openweathermap.org/users/sign_up) website.
2. Navigate to the "API keys" tab in your account dashboard.
3. Generate a new API key and copy it.
4. **Note:** It might take a few minutes for a new key to become active.

### 4. Clone and Run the Project

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/weather-telegram-bot.git
cd weather-telegram-bot

# Install necessary Go packages
go mod tidy

#Run the Bot
go run main.go

```

After running the program, you can work with the Telegram bot @SimpleMyWeatherBot
