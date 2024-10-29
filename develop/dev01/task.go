package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		// выводим ошибку в STDERR
		fmt.Fprintf(os.Stderr, "Couldn't get the time: %v\n", err)
	}
	sTime := fmt.Sprintf("%v:%v:%v", t.Hour(), t.Minute(), t.Second())
	fmt.Printf("Время сейчас - %v\n", sTime)
}
