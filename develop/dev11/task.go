package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Event представляет событие в календаре
type Event struct {
	ID        string    `json:"id"`                   // Уникальный идентификатор события
	UserID    int       `json:"user_id"`              // ID пользователя
	Title     string    `json:"title"`                // Название события
	Date      time.Time `json:"date"`                 // Дата события
	UpdatedAt time.Time `json:"updated_at,omitempty"` // Дата последнего обновления (если было)
}

// Storage для хранения событий
type Storage struct {
	events map[string]Event // Мапа для хранения событий, ключ - ID события
}

// Config структура для хранения конфигурации
type Config struct {
	Port string `json:"port"` // Порт для запуска HTTP-сервера
}

// NewStorage инициализирует пустое хранилище
func NewStorage() *Storage {
	return &Storage{
		events: make(map[string]Event), // Инициализация пустой мапы
	}
}

// AddEvent добавляет новое событие
func (s *Storage) AddEvent(event Event) error {
	if _, exists := s.events[event.ID]; exists {
		return errors.New("event already exists") // Проверка на существование события
	}
	s.events[event.ID] = event // Сохранение события
	return nil
}

// UpdateEvent обновляет событие
func (s *Storage) UpdateEvent(event Event) error {
	if _, exists := s.events[event.ID]; !exists {
		return errors.New("event not found") // Проверка, что событие существует
	}
	event.UpdatedAt = time.Now() // Устанавливаем время обновления
	s.events[event.ID] = event   // Сохраняем обновление
	return nil
}

// DeleteEvent удаляет событие
func (s *Storage) DeleteEvent(eventID string) error {
	if _, exists := s.events[eventID]; !exists {
		return errors.New("event not found") // Проверка, что событие существует
	}
	delete(s.events, eventID) // Удаление события
	return nil
}

// GetEventsForDate возвращает события за конкретную дату
func (s *Storage) GetEventsForDate(date time.Time) []Event {
	events := []Event{}
	for _, event := range s.events {
		// Сравниваем только дату, игнорируя время
		if event.Date.Format("2006-01-02") == date.Format("2006-01-02") {
			events = append(events, event)
		}
	}
	return events
}

// GetEventsForRange возвращает события за диапазон
func (s *Storage) GetEventsForRange(start, end time.Time) []Event {
	events := []Event{}
	for _, event := range s.events {
		// Проверяем, находится ли дата события в заданном диапазоне
		if event.Date.After(start) && event.Date.Before(end) {
			events = append(events, event)
		}
	}
	return events
}

// parseDate парсит дату из строки
func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr) // Формат даты в формате ISO (YYYY-MM-DD)
}

// respondJSON отправляет JSON-ответ
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)           // Устанавливаем HTTP-статус
	json.NewEncoder(w).Encode(data) // Кодируем данные в JSON и отправляем клиенту
}

// parseAndValidateParams парсит параметры из формы
func parseAndValidateParams(r *http.Request, keys ...string) (map[string]string, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err // Ошибка при парсинге параметров
	}
	params := make(map[string]string)
	for _, key := range keys {
		value := r.FormValue(key)
		if value == "" {
			return nil, fmt.Errorf("missing parameter: %s", key) // Отсутствует обязательный параметр
		}
		params[key] = value
	}
	return params, nil
}

// createEventHandler создает новое событие
func createEventHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseAndValidateParams(r, "user_id", "title", "date")
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		userID, err := strconv.Atoi(params["user_id"]) // Преобразуем user_id в число
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid user_id"})
			return
		}

		date, err := parseDate(params["date"]) // Парсим дату
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid date"})
			return
		}

		// Создаем объект события
		event := Event{
			ID:     fmt.Sprintf("%d-%s", userID, date.Format("20060102")), // Генерация уникального ID
			UserID: userID,
			Title:  params["title"],
			Date:   date,
		}

		// Сохраняем событие в хранилище
		if err := storage.AddEvent(event); err != nil {
			respondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
			return
		}

		respondJSON(w, http.StatusOK, map[string]string{"result": "event created"})
	}
}

// updateEventHandler обновляет событие
func updateEventHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseAndValidateParams(r, "id", "title", "date")
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		date, err := parseDate(params["date"]) // Парсим дату
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid date"})
			return
		}

		// Создаем объект обновленного события
		event := Event{
			ID:    params["id"],
			Title: params["title"],
			Date:  date,
		}

		// Обновляем событие
		if err := storage.UpdateEvent(event); err != nil {
			respondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
			return
		}

		respondJSON(w, http.StatusOK, map[string]string{"result": "event updated"})
	}
}

// deleteEventHandler удаляет событие
func deleteEventHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseAndValidateParams(r, "id")
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err := storage.DeleteEvent(params["id"]); err != nil {
			respondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
			return
		}

		respondJSON(w, http.StatusOK, map[string]string{"result": "event deleted"})
	}
}

// eventsForDayHandler возвращает события на конкретную дату
func eventsForDayHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseAndValidateParams(r, "date")
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		date, err := parseDate(params["date"]) // Парсим дату
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid date"})
			return
		}

		// Получаем события на день
		events := storage.GetEventsForDate(date)
		respondJSON(w, http.StatusOK, events)
	}
}

// eventsForWeekHandler возвращает события на неделю
func eventsForWeekHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseAndValidateParams(r, "start")
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		start, err := parseDate(params["start"]) // Парсим начальную дату
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid start date"})
			return
		}

		// Вычисляем конец недели
		end := start.AddDate(0, 0, 7)
		events := storage.GetEventsForRange(start, end)
		respondJSON(w, http.StatusOK, events)
	}
}

// eventsForMonthHandler возвращает события на месяц
func eventsForMonthHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseAndValidateParams(r, "start")
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		start, err := parseDate(params["start"]) // Парсим начальную дату
		if err != nil {
			respondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid start date"})
			return
		}

		// Вычисляем конец месяца
		end := start.AddDate(0, 1, 0)
		events := storage.GetEventsForRange(start, end)
		respondJSON(w, http.StatusOK, events)
	}
}

// loggingMiddleware добавляет логирование всех запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s, URL: %s", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}

// loadConfig читает конфигурацию из файла
func loadConfig() (Config, error) {
	file, err := ioutil.ReadFile("config.json") // Чтение файла конфигурации
	if err != nil {
		return Config{}, fmt.Errorf("could not read config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, fmt.Errorf("could not unmarshal config: %v", err)
	}

	// Проверка наличия порта в конфигурации
	if config.Port == "" {
		return Config{}, errors.New("port not specified in config")
	}

	return config, nil
}

// Основной запуск сервера
func main() {
	storage := NewStorage() // Инициализация хранилища событий

	// Попытка загрузить конфигурацию
	config, err := loadConfig()
	if err != nil {
		// Если ошибка при загрузке конфигурации, используем дефолтный порт
		log.Printf("Error loading config: %v. Using default port :8080", err)
		config.Port = ":8080"
	}

	// Создаем маршруты
	mux := http.NewServeMux()
	mux.Handle("/create_event", createEventHandler(storage))
	mux.Handle("/update_event", updateEventHandler(storage))
	mux.Handle("/delete_event", deleteEventHandler(storage))
	mux.Handle("/events_for_day", eventsForDayHandler(storage))
	mux.Handle("/events_for_week", eventsForWeekHandler(storage))
	mux.Handle("/events_for_month", eventsForMonthHandler(storage))

	// Добавляем логирование
	loggedMux := loggingMiddleware(mux)

	log.Printf("Server running on %s", config.Port)
	if err := http.ListenAndServe(config.Port, loggedMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
