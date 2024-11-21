package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Флаги
	fieldsFlag := flag.String("f", "", "Выбрать поля (например, '1,3,5')")
	delimiterFlag := flag.String("d", "\t", "Разделитель (по умолчанию TAB)")
	separatedFlag := flag.Bool("s", false, "Выводить только строки с разделителями")
	flag.Parse()

	// Проверка, что аргументы переданы (имя файла должно быть первым аргументом)
	if len(flag.Args()) == 0 {
		fmt.Println("Ошибка: не указан файл. Укажите имя файла как первый аргумент.")
		return
	}

	// Имя файла — это первый аргумент
	fileName := flag.Args()[0]

	// Получаем строку с полями
	var selectedFields []int
	if *fieldsFlag != "" {
		for _, f := range strings.Split(*fieldsFlag, ",") {
			fieldNum, err := strconv.Atoi(f)
			if err != nil {
				fmt.Println("Ошибка в указанных полях:", err)
				return
			}
			selectedFields = append(selectedFields, fieldNum)
		}
	}

	// Открытие файла для чтения
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Чтение данных из файла
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Разделяем строку на колонки
		columns := strings.Split(line, *delimiterFlag)

		// Если задан флаг -s, проверяем, есть ли разделитель в строке
		if *separatedFlag && len(columns) <= 1 {
			continue
		}

		// Если поля не указаны, выводим всю строку
		if len(selectedFields) == 0 {
			fmt.Println(line)
			continue
		}

		// Выводим только те поля, которые были указаны в флаге -f
		var output []string
		for _, field := range selectedFields {
			// Проверяем, что индекс поля не выходит за границы
			if field-1 < len(columns) {
				output = append(output, columns[field-1])
			}
		}

		// Если есть вывод, выводим строку
		if len(output) > 0 {
			fmt.Println(strings.Join(output, *delimiterFlag))
		}
	}

	// Проверяем на ошибки ввода
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения данных:", err)
	}
}
