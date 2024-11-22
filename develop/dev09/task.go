package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Функция для скачивания файла
func downloadFile(url, dest string) error {
	// Отправляем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создаем файл для сохранения
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	// Копируем данные из ответа в файл
	_, err = io.Copy(out, resp.Body)
	return err
}

// Функция для скачивания HTML страницы
func downloadPage(url, dest string) error {
	return downloadFile(url, dest)
}

// Функция для извлечения ссылок на ресурсы из HTML (очень простая версия)
func extractLinks(url string) ([]string, error) {
	// Загружаем HTML страницы
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var links []string
	var htmlContent []byte
	htmlContent, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Ищем ссылки на изображения, CSS и JS
	content := string(htmlContent)
	for _, tag := range []string{"src", "href"} {
		start := 0
		for {
			pos := strings.Index(content[start:], tag)
			if pos == -1 {
				break
			}
			start += pos + len(tag) + 2 // пропускаем атрибут и символы после
			end := strings.Index(content[start:], `"`)
			if end == -1 {
				break
			}
			links = append(links, content[start:start+end])
			start += end
		}
	}
	return links, nil
}

func main() {
	// Проверка на наличие аргумента URL
	if len(os.Args) < 2 {
		fmt.Println("Пожалуйста, укажите URL сайта как аргумент.")
		return
	}
	baseURL := os.Args[1]

	// Папка для сохранения
	dirName := "downloaded_site"
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Println("Ошибка при создании директории:", err)
		return
	}

	// Скачиваем HTML страницу
	htmlFile := filepath.Join(dirName, "index.html")
	err = downloadPage(baseURL, htmlFile)
	if err != nil {
		fmt.Printf("Ошибка при скачивании страницы: %v\n", err)
		return
	}
	fmt.Println("HTML страница скачана в", htmlFile)

	// Извлекаем все ссылки на ресурсы
	links, err := extractLinks(baseURL)
	if err != nil {
		fmt.Printf("Ошибка при извлечении ссылок: %v\n", err)
		return
	}

	// Скачиваем все ресурсы
	for _, link := range links {
		// Преобразуем относительные пути в абсолютные
		absoluteURL := link
		if !strings.HasPrefix(link, "http") {
			absoluteURL = baseURL + link
		}

		// Получаем имя файла для сохранения
		filename := filepath.Join(dirName, filepath.Base(link))
		err := downloadFile(absoluteURL, filename)
		if err != nil {
			fmt.Printf("Ошибка при скачивании файла %s: %v\n", absoluteURL, err)
		} else {
			fmt.Printf("Скачан файл %s\n", filename)
		}
	}

	fmt.Println("Загрузка завершена.")
}
