package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Главная функция программы
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Отображаем приглашение
		fmt.Print("my_shell> ")

		// Считываем ввод пользователя
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
			continue
		}

		// Убираем лишние пробелы и символы новой строки
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		// Проверяем команду на выход
		if input == "\\quit" {
			fmt.Println("Выход из шелла.")
			break
		}

		// Разделение команды на части (обработка пайпов)
		commands := strings.Split(input, "|")
		if len(commands) > 1 {
			handlePipes(commands)
			continue
		}

		// Обработка одиночной команды
		handleCommand(input)
	}
}

// Функция обработки одиночной команды
func handleCommand(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		changeDirectory(args)
	case "pwd":
		printWorkingDirectory()
	case "echo":
		echo(args)
	case "kill":
		killProcess(args)
	case "ps":
		listProcesses()
	default:
		executeCommand(args)
	}
}

// Команда `cd`
func changeDirectory(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: отсутствует аргумент для cd")
		return
	}
	err := os.Chdir(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка смены директории:", err)
	}
}

// Команда `pwd`
func printWorkingDirectory() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка получения текущей директории:", err)
		return
	}
	fmt.Println(cwd)
}

// Команда `echo`
func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

// Команда `kill`
func killProcess(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: необходимо указать PID")
		return
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка: неверный PID")
		return
	}

	// Находим процесс по PID
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка: процесс не найден:", err)
		return
	}

	// Завершаем процесс
	err = process.Kill()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка завершения процесса:", err)
	} else {
		fmt.Printf("Процесс %d успешно завершён.\n", pid)
	}
}

// Команда `ps`
func listProcesses() {
	cmd := exec.Command("ps", "-e")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка выполнения команды ps:", err)
	}
}

// Выполнение произвольной команды
func executeCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка выполнения команды:", err)
	}
}

// Обработка пайпов
func handlePipes(commands []string) {
	var processes []*exec.Cmd

	// Создание списка процессов
	for _, command := range commands {
		args := strings.Fields(strings.TrimSpace(command))
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Ошибка: пустая команда в пайпе")
			return
		}
		processes = append(processes, exec.Command(args[0], args[1:]...))
	}

	// Настройка пайпов
	for i := 0; i < len(processes)-1; i++ {
		pipe, err := processes[i].StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка создания пайпа:", err)
			return
		}
		processes[i+1].Stdin = pipe
	}

	// Настройка последнего процесса для вывода
	processes[len(processes)-1].Stdout = os.Stdout

	// Запуск всех процессов
	for _, process := range processes {
		err := process.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка запуска команды:", err)
			return
		}
	}

	// Ожидание завершения процессов
	for _, process := range processes {
		err := process.Wait()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка выполнения команды:", err)
		}
	}
}
