package domain

import (
	"fmt"
	"time"
)

type WttrInResponse struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
	NearestArea      []NearestArea      `json:"nearest_area"`
}

type CurrentCondition struct {
	TempC         string        `json:"temp_C"`
	Humidity      string        `json:"humidity"`
	WeatherDesc   []WeatherDesc `json:"weatherDesc"`
	WindSpeedKmph string        `json:"windspeedKmph"`
	FeelsLikeC    string        `json:"FeelsLikeC"`
}

type WeatherDesc struct {
	Value string `json:"value"`
}

type NearestArea struct {
	AreaName []AreaName `json:"areaName"`
}

type AreaName struct {
	Value string `json:"value"`
}

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
	Description string  `json:"description"`
	WindSpeed   float64 `json:"wind_speed"`
	FeelsLike   float64 `json:"feels_like"`
}

// Display отображает погоду в консоли
func (w *WeatherData) Display() {
	fmt.Printf("\n🌤️  Погода в %s\n", w.City)
	fmt.Printf("🌡️  Температура: %.1f°C\n", w.Temperature)
	fmt.Printf("🤔 Ощущается как: %.1f°C\n", w.FeelsLike)
	fmt.Printf("💧 Влажность: %d%%\n", w.Humidity)
	fmt.Printf("💨 Скорость ветра: %.1f км/ч\n", w.WindSpeed)
	fmt.Printf("📝 Описание: %s\n", w.Description)
	fmt.Printf("🕒 Время запроса: %s\n", time.Now().Format("15:04:05"))
}
