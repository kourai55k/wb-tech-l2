// Пакет pattern демонстрирует пример паттерна "Фабричный метод".
package pattern

import (
	"fmt"
	"log"
)

// action определяет доступные действия для клиентов.
type action string

const (
	A action = "A" // Действие "A"
	B action = "B" // Действие "B"
	C action = "C" // Действие "C"
)

// Creator предоставляет интерфейс фабрики.
type Creator interface {
	CreateProduct(action action) Product // Фабричный метод
}

// Product предоставляет интерфейс для всех продуктов.
// Все продукты, возвращаемые фабрикой, должны реализовывать этот интерфейс.
type Product interface {
	Use() string // Все продукты должны быть "используемыми"
}

// ConcreteCreator реализует интерфейс Creator.
type ConcreteCreator struct{}

// NewCreator конструктор для ConcreteCreator.
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// CreateProduct реализует фабричный метод.
// Создает объект конкретного продукта в зависимости от действия.
func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	case C:
		product = &ConcreteProductC{string(action)}
	default:
		log.Fatalln("Неизвестное действие")
	}

	return product
}

// ConcreteProductA представляет продукт "A".
type ConcreteProductA struct {
	action string
}

// Use возвращает действие для продукта "A".
func (p *ConcreteProductA) Use() string {
	return p.action
}

// ConcreteProductB представляет продукт "B".
type ConcreteProductB struct {
	action string
}

// Use возвращает действие для продукта "B".
func (p *ConcreteProductB) Use() string {
	return p.action
}

// ConcreteProductC представляет продукт "C".
type ConcreteProductC struct {
	action string
}

// Use возвращает действие для продукта "C".
func (p *ConcreteProductC) Use() string {
	return p.action
}

// Функция main демонстрирует использование фабричного метода.
func Run6() {
	// Создаем фабрику
	creator := NewCreator()

	// Создаем продукт A
	productA := creator.CreateProduct(A)
	fmt.Printf("Продукт A: %s\n", productA.Use())

	// Создаем продукт B
	productB := creator.CreateProduct(B)
	fmt.Printf("Продукт B: %s\n", productB.Use())

	// Создаем продукт C
	productC := creator.CreateProduct(C)
	fmt.Printf("Продукт C: %s\n", productC.Use())

	// Попытка создать неизвестный продукт (выдаст ошибку)
	// creator.CreateProduct("D")
}
