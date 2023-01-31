package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	loopCount := 0
	result := 0.0
	fmt.Println("Hello, Welcome!")
	function := ""
	for {
		if loopCount == 0 {
			number, _ := getNumber(loopCount)
			result = number
		} else if loopCount%2 == 0 {
			number, _ := getNumber(loopCount)
			switch function {
			case "1":
				result = add(result, number)
			case "2":
				result = multiply(result, number)
			case "3":
				result = substract(result, number)
			case "4":
				result = divide(result, number)
			default:
				break
			}
		} else {
			fmt.Println("Write your function please(add = 1, substract = 2, multiply = 3, divide = 4 or done = 0)")
			scannerFunction := bufio.NewScanner(os.Stdin)
			scannerFunction.Scan()
			function = scannerFunction.Text()
			if function == "0" {
				// time.Sleep(2 * time.Second)
				break
			}
		}
		loopCount++
	}
	fmt.Println("[done]: GoodBye, see you soon!")
	fmt.Println("[result]:", result)
}

func getNumber(loopCount int) (float64, error) {
	fmt.Println("Write your NUMBER", loopCount, "please")
	scannerNumber := bufio.NewScanner(os.Stdin)
	scannerNumber.Scan()
	return strconv.ParseFloat(scannerNumber.Text(), 64)
}

func add(x float64, y float64) float64 {
	return x + y
}

func substract(x float64, y float64) float64 {
	return x - y
}

func multiply(x float64, y float64) float64 {
	return x * y
}

func divide(x float64, y float64) float64 {
	return x / y
}
