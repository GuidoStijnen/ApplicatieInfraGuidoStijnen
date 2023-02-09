package main

import (
	"fmt"
	"time"
)

func get_time() func() string {
	welkom_message := ""
	daytime := time.Now()

	return func() string {
		switch {
		case daytime.Hour() >= 7 && daytime.Hour() < 12:
			welkom_message = "goeie morgen bezoeker"
		case daytime.Hour() >= 12 && daytime.Hour() < 16:
			welkom_message = "goeie middag bezoeker"
		case daytime.Hour() >= 16 && daytime.Hour() < 23:
			welkom_message = "goeie avond bezoeker"
		default:
			fmt.Println("")
		}
		return welkom_message
	}
}
func main() {
	time := get_time()
	fmt.Println(time())
}
