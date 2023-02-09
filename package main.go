package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func contains(licence_plates [5]string, plates string) bool {
	for v := 0; v < len(licence_plates); v++ {
		if licence_plates[v] == plates {
			return true
		}
	}
	return false

}

func get_time() func() string {
	welkom_message := ""
	daytime := time.Now()

	return func() string {
		switch {
		case daytime.Hour() >= 7 && daytime.Hour() < 12:
			welkom_message = "goeie morgen"
		case daytime.Hour() >= 12 && daytime.Hour() < 16:
			welkom_message = "goeie middag"
		case daytime.Hour() >= 16 && daytime.Hour() < 23:
			welkom_message = "goeie avond"
		default:
			fmt.Println("")
		}
		return welkom_message
	}
}
func main() {
	daytime := time.Now()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the licenseplate: ")
	input, _ := reader.ReadString('\n')
	licence_plates := [5]string{"65GPFK\r\n", "DSR17\r\n", "6969\r\n", "DSK815\r\n", "8888\r\n"} //\r\n gay shit
	result := contains(licence_plates, input)
	time := get_time()
	switch {
	case daytime.Hour() >= 23 && daytime.Hour() < 7:
		fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten")
	default:
		if !result {
			fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
		} else {
			fmt.Println(time(), "Welkom bij fonteyn vakantieparken")
		}
	}

}
