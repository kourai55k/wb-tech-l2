package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Simple anagrams",
			input: []string{"пятак", "тяпка", "пятка", "листок", "столик", "слиток"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "Single word groups excluded",
			input: []string{"дом", "мод", "кот", "лес"},
			expected: map[string][]string{
				"дом": {"дом", "мод"},
			},
		},
		{
			name:  "Case insensitive handling",
			input: []string{"Пятак", "пятка", "ТЯПКА", "листок", "слиток", "СТОЛИК"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "Repeating words ignored",
			input: []string{"пятак", "пятак", "тяпка", "пятка", "тяпка"},
			expected: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "Words without anagrams",
			input:    []string{"один", "два", "три", "четыре"},
			expected: map[string][]string{},
		},
		{
			name:  "Complex anagram groups",
			input: []string{"море", "ремо", "роме", "перо", "ропе", "ероп", "дом", "мод"},
			expected: map[string][]string{
				"море": {"море", "ремо", "роме"},
				"перо": {"ероп", "перо", "ропе"},
				"дом":  {"дом", "мод"},
			},
		},
		{
			name:  "Words with spaces or invalid strings",
			input: []string{"мир", "рим", "м и р", "рим", "ми р", "мри"},
			expected: map[string][]string{
				"мир": {"мир", "мри", "рим"},
			},
		},
		{
			name: "Large dataset",
			input: []string{
				"листок", "слиток", "столик", "кот", "ток", "окт",
				"дом", "мод", "дмо", "пятак", "тяпка", "пятка", "ероп", "перо", "ропе",
				"город", "дорог", "годор", "догор", "пример", "ремипр",
			},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"кот":    {"кот", "окт", "ток"},
				"дом":    {"дмо", "дом", "мод"},
				"пятак":  {"пятак", "пятка", "тяпка"},
				"ероп":   {"ероп", "перо", "ропе"},
				"город":  {"годор", "город", "догор", "дорог"},
				"пример": {"пример", "ремипр"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindAnagrams(&tt.input)
			if !reflect.DeepEqual(*result, tt.expected) {
				t.Errorf("Test %s failed: got %v, want %v", tt.name, *result, tt.expected)
			}
		})
	}
}
