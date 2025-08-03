package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Nachricht von Kanal 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Nachricht von Kanal 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			fmt.Println("Keine Nachrichten verfuegbar")
			time.Sleep(500 * time.Millisecond)
			i-- // Damit die Schleife nicht zu früh endet
		}
	}
}