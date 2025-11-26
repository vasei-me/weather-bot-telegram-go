package weather

import "strings"

type WeatherService struct {
	client     WeatherClient
	translator WeatherTranslator
}

type WeatherTranslator interface {
	Translate(description string) string
}

type WeatherInfo struct {
	Location    string
	Country     string
	Temperature float64
	FeelsLike   float64
	Humidity    int
	Description string
}

func NewWeatherService(client WeatherClient, translator WeatherTranslator) *WeatherService {
	return &WeatherService{
		client:     client,
		translator: translator,
	}
}

func (s *WeatherService) GetWeatherInfo(lat, lon float64) (*WeatherInfo, error) {
	weather, err := s.client.GetWeather(lat, lon)
	if err != nil {
		return nil, err
	}

	locationName := weather.Name
	if locationName == "" {
		locationName = "مکان شما"
	}

	description := "نامشخص"
	if len(weather.Weather) > 0 {
		desc := strings.ToLower(weather.Weather[0].Description)
		description = s.translator.Translate(desc)

		if description == "نامشخص" && weather.Weather[0].Main != "" {
			description = s.translator.Translate(strings.ToLower(weather.Weather[0].Main))
		}
	}

	tempC := weather.Main.Temp - 273.15
	feelsC := weather.Main.FeelsLike - 273.15

	return &WeatherInfo{
		Location:    locationName,
		Country:     weather.Sys.Country,
		Temperature: tempC,
		FeelsLike:   feelsC,
		Humidity:    weather.Main.Humidity,
		Description: description,
	}, nil
}