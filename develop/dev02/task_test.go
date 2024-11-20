package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectError bool
	}{
		// Корректные строки
		{`a4bc2d5e`, `aaaabccddddde`, false},
		{`a4b12`, `aaaabbbbbbbbbbbb`, false},
		{`abcd`, `abcd`, false},
		{`a4\3b2`, `aaaa3bb`, false},
		{`a4\\2b`, `aaaa\\b`, false},
		{`qwe\4\5`, `qwe45`, false},
		{`qwe\45`, `qwe44444`, false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{`a4\b20`, `aaaabbbbbbbbbbbbbbbbbbbb`, false},
		{`a1b2c3`, `abbccc`, false},
		{``, ``, false},

		// Некорректные строки
		{`45`, ``, true},     // Начало строки с цифры
		{`a4c5e\`, ``, true}, // Строка заканчивается \
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := Unpack(test.input)
			if (err != nil) != test.expectError {
				t.Errorf("unexpected error status: got %v, want error: %v", err, test.expectError)
			}
			if result != test.expected {
				t.Errorf("unexpected result: got %q, want %q", result, test.expected)
			}
		})
	}
}
