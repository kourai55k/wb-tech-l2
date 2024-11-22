package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Разбираем аргументы командной строки
	timeout := flag.String("timeout", "10s", "Connection timeout duration")
	flag.Parse()

	host := flag.Args()[0]
	port := flag.Args()[1]

	// Преобразуем timeout в объект типа time.Duration
	timeoutDuration, err := time.ParseDuration(*timeout)
	if err != nil {
		fmt.Println("Invalid timeout duration:", err)
		return
	}

	// Формируем строку для подключения
	address := fmt.Sprintf("%s:%s", host, port)

	// Устанавливаем TCP соединение с таймаутом
	conn, err := net.DialTimeout("tcp", address, timeoutDuration)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Запускаем горутину для чтения данных из сокета и вывода в STDOUT
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println("Error reading from server:", err)
		}
	}()

	// Копируем данные из STDIN в сокет
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Println("Error sending data to server:", err)
	}
}
