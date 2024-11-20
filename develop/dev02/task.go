package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// processCountStr обрабатывает накопленные повторения
// Преобразует строку с числовыми символами в количество повторений
// и добавляет повторяющиеся символы в результат.
func processCountStr(countStr string, lastChar rune, result *strings.Builder) (string, error) {
	if countStr == "" {
		return "", nil // Если нет числа, сразу возвращаем пустое значение
	}

	// Преобразуем строку с числом в целое число
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return "", errors.New("invalid string: failed to parse repeat count") // Если не удалось преобразовать, ошибка
	}

	// Добавляем символ в результат требуемое количество раз
	for j := 0; j < count-1; j++ {
		result.WriteRune(lastChar)
	}

	return "", nil // Возвращаем пустую строку после обработки числа
}

// Unpack выполняет распаковку строки с учётом повторений и экранирования
func Unpack(input string) (string, error) {
	if input == "" {
		return "", nil // Если строка пуста, возвращаем ошибку
	}

	// Используем strings.Builder для эффективной работы с результатом
	var result strings.Builder
	// Флаг для отслеживания экранирования
	var escape bool
	// Строка для хранения числового значения повторений
	var countStr string
	// Переменная для хранения последнего добавленного символа
	var lastChar rune
	// Флаг, который указывает, был ли символ перед числом
	var hasPrev bool

	// Перебираем символы входной строки
	for _, r := range input {
		switch {
		case escape: // Обрабатываем экранированный символ
			// Применяем предыдущие накопленные повторения перед экранированным символом
			countStr, _ = processCountStr(countStr, lastChar, &result)
			// Добавляем экранированный символ в результат
			result.WriteRune(r)
			lastChar = r
			hasPrev = true
			escape = false // Сброс флага экранирования
		case r == '\\': // Включаем режим экранирования
			// Применяем предыдущие накопленные повторения перед экранированием
			countStr, _ = processCountStr(countStr, lastChar, &result)
			escape = true // Включаем экранирование для следующего символа
		case unicode.IsDigit(r): // Если символ — цифра
			if !hasPrev {
				return "", errors.New("invalid string: digit without preceding character") // Ошибка, если число стоит перед символом
			}
			countStr += string(r) // Добавляем цифру к строке числа
		default: // Обычный символ
			// Применяем предыдущие накопленные повторения перед добавлением нового символа
			countStr, _ = processCountStr(countStr, lastChar, &result)
			result.WriteRune(r) // Добавляем символ в результат
			lastChar = r        // Обновляем последний символ
			hasPrev = true      // Устанавливаем флаг, что был предыдущий символ
		}
	}

	// После завершения перебора строки проверяем, нужно ли применить повторения для последнего символа
	_, err := processCountStr(countStr, lastChar, &result)
	if err != nil {
		return "", err
	}

	// Проверяем, если строка заканчивается на экранированный символ, возвращаем ошибку
	if escape {
		return "", errors.New("invalid string: ends with escape character")
	}

	return result.String(), nil // Возвращаем результат
}
