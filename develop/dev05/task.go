package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Определение флагов командной строки
	after := flag.Int("A", 0, "Print N lines after the match")
	before := flag.Int("B", 0, "Print N lines before the match")
	context := flag.Int("C", 0, "Print N lines before and after the match")
	count := flag.Bool("c", false, "Count the matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert the match")
	fixed := flag.Bool("F", false, "Exact string match (no regex)")
	lineNum := flag.Bool("n", false, "Print line numbers")

	// Парсим флаги
	flag.Parse()

	// Позиционные аргументы
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: <pattern> <file>")
		return
	}

	// Шаблон поиска — первый позиционный аргумент
	pattern := args[0]
	// Имя файла — второй позиционный аргумент
	filePath := args[1]

	// Открытие файла для чтения
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Чтение строк из файла
	scanner := bufio.NewScanner(file)
	var lineNumCounter int
	var matchingLines []string

	// Чтение и обработка каждой строки
	for scanner.Scan() {
		lineNumCounter++
		line := scanner.Text()

		// Проверка совпадения
		matched := matchLine(line, pattern, *fixed, *ignoreCase)

		// Если инвертирование поиска, то ищем строки, не совпавшие
		if *invert {
			matched = !matched
		}

		// Если совпадение, обрабатываем вывод
		if matched {
			matchingLines = append(matchingLines, line)

			// Если флаг count
			if *count {
				continue
			}

			// Напечатаем номер строки, если нужно
			if *lineNum {
				fmt.Printf("%d: %s\n", lineNumCounter, line)
			} else {
				fmt.Println(line)
			}

			// Печать контекста (строки до и после совпадения)
			printContext(scanner, lineNumCounter, *before, *after, *context)
		}
	}

	// Если требуется подсчитать количество строк
	if *count {
		fmt.Printf("Matching lines: %d\n", len(matchingLines))
	}
}

// Функция для проверки совпадения строки с паттерном
func matchLine(line, pattern string, fixed bool, ignoreCase bool) bool {
	// Если нужно игнорировать регистр
	if ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	if fixed {
		return line == pattern
	}
	return strings.Contains(line, pattern)
}

// Печать контекста строк до и после совпадения
func printContext(scanner *bufio.Scanner, lineNumCounter, before, after, context int) {
	// var contextLines []string

	// Печать строк до совпадения
	if before > 0 {
		for i := 0; i < before && scanner.Scan(); i++ {
			fmt.Println(scanner.Text())
		}
	}

	// Печать строки с совпадением
	if context > 0 {
		for i := 0; i < context && scanner.Scan(); i++ {
			fmt.Println(scanner.Text())
		}
	}

	// Печать строк после совпадения
	if after > 0 {
		for i := 0; i < after && scanner.Scan(); i++ {
			fmt.Println(scanner.Text())
		}
	}
}
