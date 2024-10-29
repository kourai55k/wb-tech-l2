package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

// Unpack распаковывет строку
func Unpack(s string) (string, error) {
	r := []rune(s)
	var resRune []rune
	// Если строка пустая, возвращаем пустую строку
	if len(r) == 0 {
		log.Println("string is empty")
		return "", nil
	}
	// Если строка начинается с цифры, возвращаем ошибку
	if unicode.IsDigit(r[0]) {
		return "", fmt.Errorf("invalid string - line starts with number")
	}
	// Проходим в цикле по срезу рун
	for i := 0; i < len(r); i++ {
		// Если следующий символ - цифра, то добавляем в результат символ столько раз
		if i < len(r)-1 && unicode.IsDigit(r[i+1]) {
			// Превращаем руну в число, чтобы использовать его в цикле for
			sym, err := strconv.Atoi(string(r[i+1]))
			if err != nil {
				return "", fmt.Errorf("couldn't convert string to int")
			}
			// Добавляем символы в результат
			for j := 0; j < sym; j++ {
				resRune = append(resRune, r[i])
			}
			// Пропускаем цифру
			i++
		} else { // В противном случае добавляем символ один раз
			if unicode.IsDigit(r[i]) {
				return "", fmt.Errorf("invalid string - num after num")
			}
			resRune = append(resRune, r[i])
		}
	}
	// Превращаем слайс рун в строку
	res := string(resRune)
	// Возвращаем строку
	return res, nil
}

func main() {
	res, err := Unpack("a10b5")
	if err != nil {
		log.Fatalf("Couldn't unpack string: %v", err)
	}
	fmt.Println(res)
}
