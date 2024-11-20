package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// uniqueLines удаляет дубликаты строк
func uniqueLines(lines []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}
	return result
}

// equal сравнивает два слайса строк
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// numericSort сравнивает два значения с суффиксами (например, 1K < 1M)
func numericSort(colI, colJ string) bool {
	// Извлекаем числовое значение и суффикс из строк
	valI, suffixI := extractNumericValue(colI)
	valJ, suffixJ := extractNumericValue(colJ)

	// Применяем множители для учета суффиксов
	valI = applySuffixMultiplier(valI, suffixI)
	valJ = applySuffixMultiplier(valJ, suffixJ)

	// Сравниваем числовые значения после применения множителей
	return valI < valJ
}

// extractNumericValue извлекает числовое значение и суффикс из строки
func extractNumericValue(s string) (float64, string) {
	// Ищем числовую часть и суффикс в строке
	numberPart := ""
	suffix := ""

	for i, ch := range s {
		if ch >= '0' && ch <= '9' || ch == '.' {
			numberPart += string(ch) // Добавляем символ в числовую часть
		} else {
			suffix = s[i:] // Остальная часть строки — это суффикс
			break
		}
	}

	// Преобразуем числовую часть в float64
	value, _ := strconv.ParseFloat(numberPart, 64)
	return value, suffix
}

// applySuffixMultiplier применяет множитель в зависимости от суффикса
func applySuffixMultiplier(value float64, suffix string) float64 {
	switch suffix {
	case "K", "k":
		return value * 1024
	case "M", "m":
		return value * 1024 * 1024
	case "G", "g":
		return value * 1024 * 1024 * 1024
	case "T", "t":
		return value * 1024 * 1024 * 1024 * 1024
	default:
		return value // если суффикс не найден, возвращаем исходное значение
	}
}

// Месяцы в году для сортировки по названию месяца
var months = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func main() {
	// Определяем флаги
	k := flag.Int("k", 0, "Номер колонки для сортировки")
	n := flag.Bool("n", false, "Сортировка по числовому значению")
	r := flag.Bool("r", false, "Сортировка в обратном порядке")
	u := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	m := flag.Bool("m", false, "Сортировать по названию месяца")
	b := flag.Bool("b", false, "Игнорировать хвостовые пробелы")
	c := flag.Bool("c", false, "Проверить отсортированность данных")
	h := flag.Bool("h", false, "Сортировать по числовому значению с учетом суффиксов")

	flag.Parse()

	// Если не указан файл, завершаем работу
	if len(flag.Args()) < 1 {
		log.Fatal("Не указан файл для сортировки")
	}

	// Получаем имя файла
	fileName := flag.Args()[0]
	// Открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Ошибка при открытии файла: ", err)
	}
	defer file.Close()

	// Получаем строки из файла и кладём их в слайс
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Если указан флаг -b, игнорируем хвостовые пробелы при чтении
		if *b {
			line = strings.TrimRight(line, " ")
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Ошибка при чтении из файла: ", err)
	}

	// Если указан флаг -u, удаляем дубликаты
	if *u {
		lines = uniqueLines(lines)
	}

	// Сортировка строк
	sort.Slice(lines, func(i, j int) bool {
		// Извлекаем нужную колонку
		colsI := strings.Fields(lines[i])
		colsJ := strings.Fields(lines[j])

		// Если задан флаг -k
		if *k > 0 && *k <= len(colsI) && *k <= len(colsJ) {
			colI := colsI[*k-1]
			colJ := colsJ[*k-1]

			// Сортируем в обратном порядке, если указан флаг -r
			if *r {
				colI, colJ = colJ, colI
				// return colI > colJ
			}

			// Если задан флаг -M, сортируем по названию месяца
			if *m {
				// Извлекаем месяц из указанной колонки
				monthI := strings.Title(colI)
				monthJ := strings.Title(colJ)

				// Проверяем, что это месяцы
				monthNumI, okI := months[monthI]
				monthNumJ, okJ := months[monthJ]

				// Если оба значения это месяцы, сортируем их
				if okI && okJ {
					return monthNumI < monthNumJ
				}
			}

			// Сортируем по числам, если указан флаг -n
			if *n {
				ni, errI := strconv.Atoi(colI)
				nj, errJ := strconv.Atoi(colJ)
				if errI == nil && errJ == nil {
					return ni < nj
				}
			}

			// Сортировка по числам с учетом суффиксов
			if *h {
				return numericSort(colI, colJ)
			}

			return colI < colJ
		}

		// Если флаг -k не задан, сортируем по строкам
		if *r {
			return lines[i] > lines[j]
		}
		return lines[i] < lines[j]
	})

	// Проверка отсортированности данных, если указан флаг -c
	if *c {
		sortedLines := make([]string, len(lines))
		copy(sortedLines, lines)
		sort.Strings(sortedLines)
		if !equal(sortedLines, lines) {
			fmt.Println("Данные не отсортированы")
			return
		}
	}

	// Запись отсортированных строк в новый файл
	outputFileName := "sorted_" + fileName
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Ошибка при создании файла для записи:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	// Не забываем очистить буфер записи
	writer.Flush()
	fmt.Printf("Данные успешно отсортированы и записаны в файл: %s\n", outputFileName)
}
