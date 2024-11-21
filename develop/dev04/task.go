package main

import (
	"sort"
	"strings"
)

func FindAnagrams(words *[]string) *map[string][]string {
	// Мапа для группировки слов по анаграммам
	anagrams := make(map[string][]string)

	// Мапа для проверки уникальности слов
	unique := make(map[string]bool)

	// Идём по слайсу строк
	for i := 0; i < len(*words); i++ {
		wrds := *words
		word := wrds[i]

		// Приводим слово к нижнему регистру
		word = strings.ToLower(word)

		// Проверяем на уникальность
		if unique[word] {
			continue
		}
		unique[word] = true

		// Получаем ключ анаграмы с помощью сортировки букв в слове
		sortedWord := sortString(word)
		// Добавляем слово в мапу анаграмм
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	// Создаем мапу, которая будет выполнять требования для результата
	res := make(map[string][]string)
	for _, group := range anagrams {
		// Множества из 1 слова не должны попасть в результат
		if len(group) > 1 {
			key := group[0]
			// Сортируем слова в группе анаграмм
			sort.Strings(group)
			// Добавляем группу анаграмм в результат
			// Ключ - первое слово из группы
			res[key] = group
		}
	}

	return &res
}

// Вспомогательная функция для сортировки символов в строке
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
