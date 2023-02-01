package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type calc struct {
}

func (calc) add(x float64, y float64) float64 {
	return x + y
}

func (calc) substract(x float64, y float64) float64 {
	return x - y
}

func (calc) multiply(x float64, y float64) float64 {
	return x * y
}

func (calc) divide(x float64, y float64) float64 {
	return x / y
}

func main() {
	loopCount := 0
	result := 0.0
	fmt.Println("Hello, Welcome!")
	function := ""
	calculator := calc{}
	for {
		if loopCount == 0 {
			number, _ := getNumber()
			result = number
		} else if loopCount%2 == 0 {
			number, _ := getNumber()
			switch function {
			case "+":
				result = calculator.add(result, number)
			case "-":
				result = calculator.substract(result, number)
			case "*":
				result = calculator.multiply(result, number)
			case "/":
				result = calculator.divide(result, number)
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
