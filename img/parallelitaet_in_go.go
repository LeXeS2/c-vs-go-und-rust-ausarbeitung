package main

import (
	"fmt"
	"time"
)

func arbeiter(arbeiterID int, aufgaben <-chan int, ergebnisse chan<- int) {
	for aufgabe := range aufgaben {
		fmt.Printf("Arbeiter %d bearbeitet Aufgabe %d\n", arbeiterID, aufgabe)
		time.Sleep(time.Second) // simuliert Arbeit
		ergebnisse <- aufgabe * 2
	}
}

func main() {
	const anzahlAufgaben = 5
	aufgaben := make(chan int, anzahlAufgaben)
	ergebnisse := make(chan int, anzahlAufgaben)

	for arbeiterNummer := 1; arbeiterNummer <= 2; arbeiterNummer++ {
		go arbeiter(arbeiterNummer, aufgaben, ergebnisse)
	}

	for aufgabenNummer := 1; aufgabenNummer <= anzahlAufgaben; aufgabenNummer++ {
		aufgaben <- aufgabenNummer
	}
	close(aufgaben)

	for i := 1; i <= anzahlAufgaben; i++ {
		ergebnis := <-ergebnisse
		fmt.Printf("Ergebnis: %d\n", ergebnis)
	}
}