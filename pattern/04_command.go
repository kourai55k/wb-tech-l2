package pattern

import "fmt"

// Command предоставляет интерфейс команды.
type Command interface {
	Execute() string // Метод для выполнения команды
}

// ToggleOnCommand реализует интерфейс Command для включения.
type ToggleOnCommand struct {
	receiver *Receiver // Получатель, который выполняет операцию
}

// Выполняет команду включения.
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn() // Вызываем метод ToggleOn у получателя
}

// ToggleOffCommand реализует интерфейс Command для выключения.
type ToggleOffCommand struct {
	receiver *Receiver // Получатель, который выполняет операцию
}

// Выполняет команду выключения.
func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff() // Вызываем метод ToggleOff у получателя
}

// Receiver представляет собой класс, который знает, как выполнять операции.
type Receiver struct {
}

// Метод для включения.
func (r *Receiver) ToggleOn() string {
	return "Toggle On" // Сообщение, что операция включения выполнена
}

// Метод для выключения.
func (r *Receiver) ToggleOff() string {
	return "Toggle Off" // Сообщение, что операция выключения выполнена
}

// Invoker представляет собой класс, который вызывает выполнение команд.
type Invoker struct {
	commands []Command // Список команд, которые будут выполнены
}

// StoreCommand добавляет команду в список.
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command) // Добавляем команду в список
}

// UnStoreCommand удаляет последнюю команду из списка.
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1] // Удаляем последнюю команду
	}
}

// Execute выполняет все команды, которые находятся в списке.
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n" // Выполняем каждую команду и добавляем результат
	}
	return result // Возвращаем все результаты выполнения команд
}

// Run4 демонстрирует работу паттерна Command.
func Run4() {
	// Создаем получателя (Receiver), который будет выполнять действия
	receiver := &Receiver{}

	// Создаем конкретные команды для включения и выключения
	toggleOnCommand := &ToggleOnCommand{receiver: receiver}
	toggleOffCommand := &ToggleOffCommand{receiver: receiver}

	// Создаем инвокер, который будет хранить команды
	invoker := &Invoker{}

	// Добавляем команды в инвокер
	invoker.StoreCommand(toggleOnCommand)
	invoker.StoreCommand(toggleOnCommand)
	invoker.StoreCommand(toggleOffCommand)
	invoker.StoreCommand(toggleOffCommand)

	// Выполняем все команды, которые хранятся в инвокере
	fmt.Println("Executing commands:")
	fmt.Print(invoker.Execute()) // Выводим результат выполнения всех команд

	// Удаляем последнюю команду
	invoker.UnStoreCommand()

	// Выполняем оставшиеся команды
	fmt.Println("\nExecuting remaining commands after undoing last command:")
	fmt.Print(invoker.Execute()) // Выводим результат выполнения оставшихся команд
}
