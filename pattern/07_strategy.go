package pattern

import "fmt"

// Strategy - интерфейс стратегии.
type Strategy interface {
	CalculatePrice(distance int, weight int) int
}

// StandardDelivery - конкретная стратегия для стандартной доставки.
type StandardDelivery struct{}

func (s *StandardDelivery) CalculatePrice(distance int, weight int) int {
	// Цена = расстояние * вес * 1 (для стандартной доставки)
	return distance * weight * 1
}

// ExpressDelivery - конкретная стратегия для экспресс-доставки.
type ExpressDelivery struct{}

func (e *ExpressDelivery) CalculatePrice(distance int, weight int) int {
	// Цена = расстояние * вес * 2 (для экспресс-доставки)
	return distance * weight * 2
}

// Context - контекст, который использует стратегию.
type DeliveryContext struct {
	Strategy Strategy
}

func (d *DeliveryContext) SetStrategy(strategy Strategy) {
	d.Strategy = strategy
}

func (d *DeliveryContext) GetDeliveryPrice(distance int, weight int) int {
	return d.Strategy.CalculatePrice(distance, weight)
}

func Run7() {
	// Создаем контекст и устанавливаем стратегию стандартной доставки
	standard := &StandardDelivery{}
	context := &DeliveryContext{}
	context.SetStrategy(standard)

	// Рассчитываем цену для стандартной доставки
	price := context.GetDeliveryPrice(100, 5) // 100 км, 5 кг
	fmt.Printf("Standard Delivery Price: %d\n", price)

	// Меняем стратегию на экспресс-доставку
	express := &ExpressDelivery{}
	context.SetStrategy(express)

	// Рассчитываем цену для экспресс-доставки
	price = context.GetDeliveryPrice(100, 5) // 100 км, 5 кг
	fmt.Printf("Express Delivery Price: %d\n", price)
}
