package pattern

import "fmt"

// Product: Пицца
type Pizza struct {
	Dough   string
	Sauce   string
	Topping string
}

// Builder: Интерфейс строителя пиццы
type PizzaBuilder interface {
	SetDough(dough string) PizzaBuilder
	SetSauce(sauce string) PizzaBuilder
	SetTopping(topping string) PizzaBuilder
	Build() *Pizza
}

// ConcreteBuilder: Реализация строителя для пиццы
type CustomPizzaBuilder struct {
	dough   string
	sauce   string
	topping string
}

// NewPizzaBuilder создает нового строителя
func NewPizzaBuilder() PizzaBuilder {
	return &CustomPizzaBuilder{}
}

func (b *CustomPizzaBuilder) SetDough(dough string) PizzaBuilder {
	b.dough = dough
	return b
}

func (b *CustomPizzaBuilder) SetSauce(sauce string) PizzaBuilder {
	b.sauce = sauce
	return b
}

func (b *CustomPizzaBuilder) SetTopping(topping string) PizzaBuilder {
	b.topping = topping
	return b
}

func (b *CustomPizzaBuilder) Build() *Pizza {
	return &Pizza{
		Dough:   b.dough,
		Sauce:   b.sauce,
		Topping: b.topping,
	}
}

// Director: Управляет процессом создания пиццы
type PizzaDirector struct {
	builder PizzaBuilder
}

func NewPizzaDirector(builder PizzaBuilder) *PizzaDirector {
	return &PizzaDirector{builder: builder}
}

func (d *PizzaDirector) ConstructMargherita() *Pizza {
	return d.builder.SetDough("Тонкое тесто").
		SetSauce("Томатный соус").
		SetTopping("Моцарелла").
		Build()
}

func (d *PizzaDirector) ConstructPepperoni() *Pizza {
	return d.builder.SetDough("Тонкое тесто").
		SetSauce("Томатный соус").
		SetTopping("Пепперони").
		Build()
}

// main демонстрирует работу паттерна "Строитель"
func Run2() {
	// Создаем строителя
	builder := NewPizzaBuilder()

	// Создаем директора и собираем пиццу "Маргарита"
	director := NewPizzaDirector(builder)
	margherita := director.ConstructMargherita()
	fmt.Printf("Маргарита: %+v\n", margherita)

	// Собираем пиццу "Пепперони"
	pepperoni := director.ConstructPepperoni()
	fmt.Printf("Пепперони: %+v\n", pepperoni)

	// Используем строителя для кастомной пиццы
	customPizza := builder.SetDough("Пышное тесто").
		SetSauce("Белый соус").
		SetTopping("Курица и грибы").
		Build()
	fmt.Printf("Кастомная пицца: %+v\n", customPizza)
}
