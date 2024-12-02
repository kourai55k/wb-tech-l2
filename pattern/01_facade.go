package pattern

import "fmt"

// Подсистема 1: Управление светом
type LightingSystem struct{}

func (l *LightingSystem) TurnOn() {
	fmt.Println("Свет включен.")
}

func (l *LightingSystem) TurnOff() {
	fmt.Println("Свет выключен.")
}

// Подсистема 2: Управление климатом
type ClimateControlSystem struct{}

func (c *ClimateControlSystem) SetTemperature(temp int) {
	fmt.Printf("Температура установлена на %d°C.\n", temp)
}

func (c *ClimateControlSystem) TurnOff() {
	fmt.Println("Система климат-контроля выключена.")
}

// Подсистема 3: Система безопасности
type SecuritySystem struct{}

func (s *SecuritySystem) Arm() {
	fmt.Println("Система безопасности активирована.")
}

func (s *SecuritySystem) Disarm() {
	fmt.Println("Система безопасности отключена.")
}

// Фасад: Унифицированный интерфейс для работы с подсистемами
type SmartHomeFacade struct {
	lighting    *LightingSystem
	climate     *ClimateControlSystem
	security    *SecuritySystem
	temperature int
}

// Конструктор для фасада
func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		lighting:    &LightingSystem{},
		climate:     &ClimateControlSystem{},
		security:    &SecuritySystem{},
		temperature: 22, // Значение по умолчанию
	}
}

// Упрощенный интерфейс для клиента
func (s *SmartHomeFacade) LeaveHome() {
	fmt.Println("Вы уходите из дома:")
	s.lighting.TurnOff()
	s.climate.TurnOff()
	s.security.Arm()
}

func (s *SmartHomeFacade) ComeHome() {
	fmt.Println("Вы приходите домой:")
	s.lighting.TurnOn()
	s.climate.SetTemperature(s.temperature)
	s.security.Disarm()
}

// Клиентский код
func Run1() {
	smartHome := NewSmartHomeFacade()

	// Выходим из дома
	smartHome.LeaveHome()

	// Возвращаемся домой
	smartHome.ComeHome()
}
