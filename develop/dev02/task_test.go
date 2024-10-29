package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		// Тесты на корректные строки
		{"a2bcd3", "aabcccd", false},   // Несколько символов с цифрами
		{"abcd", "abcd", false},        // Строка без цифр
		{"a3b2", "aaabb", false},       // Одинаковые символы
		{"a4b3c2", "aaaabbbcc", false}, // Длинные последовательности

		// Тесты на граничные случаи
		{"", "", false},   // Пустая строка
		{"a", "a", false}, // Одна буква без цифры

		// Ошибочные случаи
		{"3abc", "", true},             // Строка начинается с цифры
		{"a3b2c1d0", "aaabbcd", false}, // Ноль повторений, не добавляем символ
		{"a10b", "aaaaaaaaaab", false}, // Большое число повторений
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := Unpack(tt.input)

			// Проверяем наличие ошибки, если она ожидается
			if (err != nil) != tt.hasError {
				t.Errorf("Unpack(%q) error = %v, expected error = %v", tt.input, err, tt.hasError)
			}

			// Проверяем правильность результата
			if result != tt.expected {
				t.Errorf("Unpack(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
