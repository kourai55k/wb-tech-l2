package pattern

import (
	"fmt"
	"math"
)

// Visitor: Интерфейс посетителя
type Visitor interface {
	VisitCircle(circle *Circle)
	VisitSquare(square *Square)
	VisitRectangle(rectangle *Rectangle)
}

// Element: Интерфейс для геометрических фигур
type Shape interface {
	Accept(visitor Visitor)
}

// Circle: Конкретная фигура (Круг)
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

// Square: Конкретная фигура (Квадрат)
type Square struct {
	Side float64
}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitSquare(s)
}

// Rectangle: Конкретная фигура (Прямоугольник)
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}

// Concrete Visitor: Подсчет площади
type AreaCalculator struct {
	TotalArea float64
}

func (a *AreaCalculator) VisitCircle(circle *Circle) {
	area := math.Pi * circle.Radius * circle.Radius
	fmt.Printf("Площадь круга с радиусом %.2f: %.2f\n", circle.Radius, area)
	a.TotalArea += area
}

func (a *AreaCalculator) VisitSquare(square *Square) {
	area := square.Side * square.Side
	fmt.Printf("Площадь квадрата со стороной %.2f: %.2f\n", square.Side, area)
	a.TotalArea += area
}

func (a *AreaCalculator) VisitRectangle(rectangle *Rectangle) {
	area := rectangle.Width * rectangle.Height
	fmt.Printf("Площадь прямоугольника %.2fx%.2f: %.2f\n", rectangle.Width, rectangle.Height, area)
	a.TotalArea += area
}

// main демонстрирует работу паттерна Visitor
func Run3() {
	// Создаем список фигур
	shapes := []Shape{
		&Circle{Radius: 5},
		&Square{Side: 4},
		&Rectangle{Width: 3, Height: 6},
	}

	// Подсчет общей площади
	areaCalculator := &AreaCalculator{}
	for _, shape := range shapes {
		shape.Accept(areaCalculator)
	}
	fmt.Printf("Общая площадь всех фигур: %.2f\n", areaCalculator.TotalArea)
}
