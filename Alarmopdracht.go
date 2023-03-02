package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:] // krijg de argumenten vanaf de command line

	if len(args) == 0 { // als er geen argumenten zijn gegeven
		fmt.Println("Input aantal")
		return
	}

	iterations, err := strconv.Atoi(args[0]) // probeer het argument te converteren naar een integer

	if err != nil || iterations <= 0 { // als het argument geen positief geheel getal is return "text"
		fmt.Println("Fout no positive number")
		return
	}

	for i := 1; i <= iterations; i++ { // weergeef de tekst "Alarm x !" zo vaak als opgegeven
		fmt.Printf("Alarm %d !\n", i)
	}
}
