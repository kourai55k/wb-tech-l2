package pattern

import (
	"fmt"
)

// Handler - интерфейс обработчика.
type Handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerA реализует обработчик "A".
type ConcreteHandlerA struct {
	next Handler
}

// Реализация SendRequest для ConcreteHandlerA.
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Я обработчик 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerB реализует обработчик "B".
type ConcreteHandlerB struct {
	next Handler
}

// Реализация SendRequest для ConcreteHandlerB.
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Я обработчик 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerC реализует обработчик "C".
type ConcreteHandlerC struct {
	next Handler
}

// Реализация SendRequest для ConcreteHandlerC.
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Я обработчик 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// main функция, демонстрирующая паттерн Chain of Responsibility.
func Run5() {
	// Создаем обработчиков
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	// Строим цепочку обработчиков
	handlerA.next = handlerB
	handlerB.next = handlerC

	// Пробуем отправить запрос с разными значениями
	messages := []int{1, 2, 3, 4}

	for _, message := range messages {
		result := handlerA.SendRequest(message)
		if result != "" {
			fmt.Printf("Message %d: %s\n", message, result)
		} else {
			fmt.Printf("Message %d: Не обработано\n", message)
		}
	}
}
