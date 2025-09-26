package client

import (
	"example/src/seminar3/tasks/weather/domain"
)

// WeatherProvider интерфейс для получения погоды
type WeatherProvider interface {
	GetWeather(city string) (*domain.WeatherData, error)
}

// WeatherService основной сервис
type WeatherService struct {
	provider WeatherProvider
}

func NewWeatherService(provider WeatherProvider) *WeatherService {
	return &WeatherService{provider: provider}
}

func (w *WeatherService) GetWeather(city string) (*domain.WeatherData, error) {
	return w.provider.GetWeather(city)
}
