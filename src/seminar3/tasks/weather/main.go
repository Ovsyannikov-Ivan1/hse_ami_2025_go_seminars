package main

import (
	"fmt"
	"os"

	"example/src/seminar3/tasks/weather/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: weather <город>")
		fmt.Println("Пример: weather Moscow")
		fmt.Println("Пример: weather \"New York\"")
		fmt.Println("Пример: weather Лондон")
		os.Exit(1)
	}

	city := os.Args[1]

	provider := client.NewWttrInProvider()
	service := client.NewWeatherService(provider)

	fmt.Printf("Запрашиваю погоду для города: %s\n", city)

	data, err := service.GetWeather(city)
	if err != nil {
		fmt.Printf("❌ Ошибка: %v\n", err)
		fmt.Println("\nПодсказки:")
		fmt.Println("- Проверьте название города")
		fmt.Println("- Попробуйте английское название для международных городов")
		fmt.Println("- Убедитесь, что есть интернет-соединение")
		os.Exit(1)
	}

	data.Display()
}
