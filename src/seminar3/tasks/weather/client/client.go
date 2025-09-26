package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"example/src/seminar3/tasks/weather/domain"
)

const (
	wttrInUrl  = "https://wttr.in/%s?format=j1"
	maxRetries = 3
	retryDelay = 2 * time.Second
)

// WttrInProvider реализация для wttr.in
type WttrInProvider struct {
	client  *http.Client
	baseURL string
}

func NewWttrInProvider() *WttrInProvider {
	return &WttrInProvider{
		client:  &http.Client{Timeout: 10 * time.Second},
		baseURL: wttrInUrl,
	}
}

// GetWeather получает данные о погоде с retry логикой
func (w *WttrInProvider) GetWeather(city string) (*domain.WeatherData, error) {
	if city == "" {
		return nil, fmt.Errorf("город не может быть пустым")
	}

	url := fmt.Sprintf(w.baseURL, city)

	var lastError error
	var weatherData *domain.WeatherData

	// Retry логика
	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			fmt.Printf("Повторная попытка %d/%d через %v...\n", attempt, maxRetries, retryDelay)
			time.Sleep(retryDelay)
		}

		body, err := w.makeRequest(url)
		if err != nil {
			lastError = err
			fmt.Printf("Попытка %d неудачна: %v\n", attempt+1, err)
			continue
		}

		weatherData, err = w.parseResponse(body, city)
		if err != nil {
			lastError = err
			fmt.Printf("Попытка %d: ошибка парсинга: %v\n", attempt+1, err)
			continue
		}
		fmt.Printf("Данные успешно получены (попытка %d)\n", attempt+1)
		return weatherData, nil
	}

	return nil, fmt.Errorf("не удалось получить данные после %d попыток: %w", maxRetries+1, lastError)
}

// makeRequest выполняет HTTP запрос с обработкой ошибок
func (w *WttrInProvider) makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	// Добавляем User-Agent чтобы быть хорошим гражданином интернета
	req.Header.Set("User-Agent", "WeatherCLI/1.0 (educational project)")

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("сервер вернул ошибку: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	return body, nil
}

// parseResponse парсит JSON ответ и преобразует в доменную модель
func (w *WttrInProvider) parseResponse(body []byte, requestedCity string) (*domain.WeatherData, error) {
	var wttrResponse domain.WttrInResponse
	if err := json.Unmarshal(body, &wttrResponse); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	return w.transformResponse(&wttrResponse, requestedCity)
}

// transformResponse преобразует сырые данные API в нашу доменную модель
func (w *WttrInProvider) transformResponse(
	response *domain.WttrInResponse,
	requestedCity string,
) (*domain.WeatherData, error) {
	condition := response.CurrentCondition[0]

	temp, err := parseFloat(condition.TempC)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга температуры: %w", err)
	}

	humidity, err := parseInt(condition.Humidity)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга влажности: %w", err)
	}

	windSpeed, err := parseFloat(condition.WindSpeedKmph)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга скорости ветра: %w", err)
	}

	feelsLike, err := parseFloat(condition.FeelsLikeC)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга ощущаемой температуры: %w", err)
	}

	cityName := w.getCityName(response.NearestArea, requestedCity)

	return &domain.WeatherData{
		City:        cityName,
		Temperature: temp,
		Humidity:    humidity,
		Description: condition.WeatherDesc[0].Value,
		WindSpeed:   windSpeed,
		FeelsLike:   feelsLike,
	}, nil
}

// getCityName извлекает название города из ответа
func (w *WttrInProvider) getCityName(nearestAreas []domain.NearestArea, requestedCity string) string {
	if len(nearestAreas) > 0 && len(nearestAreas[0].AreaName) > 0 {
		return nearestAreas[0].AreaName[0].Value
	}
	return requestedCity
}

// Вспомогательные функции для парсинга
func parseFloat(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
}

func parseInt(s string) (int, error) {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}
