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
			number, _ := getNumber()
			result = number
		} else if loopCount%2 == 0 {
			number, _ := getNumber()
			switch function {
			case "+":
				result = add(result, number)
			case "-":
				result = substract(result, number)
			case "*":
				result = multiply(result, number)
			case "/":
				result = divide(result, number)
			default:
				break
			}
		} else {
			fmt.Println("Write your function please(add = +, substract = -, multiply = *, divide = / or done = 0)")
			scannerFunction := bufio.NewScanner(os.Stdin)
			scannerFunction.Scan()
			function = scannerFunction.Text()
			if function == "0" {
				break
			}
		}
		loopCount++
	}
	fmt.Println("[done]: GoodBye, see you soon!")
	fmt.Println("[result]:", result)
}

func getNumber() (float64, error) {
	fmt.Println("Write your NUMBER please")
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
