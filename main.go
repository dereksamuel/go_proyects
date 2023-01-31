package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func transformStr2Int(loopVal string) (int, error) {
	result, errorVal := strconv.Atoi(loopVal)

	return result, errorVal
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_value := scanner.Text()
	operator := "/"
	input_value_list := strings.Split(input_value, operator)
	var total int

	if operator == "-" || operator == "/" {
		rangeE := len(input_value_list) - 1
		for i := rangeE; i >= 0; i-- {
			loopVal, errorVal := transformStr2Int(input_value_list[i])
			isDone := false

			if errorVal != nil {
				fmt.Println("[error]:", errorVal)
				break
			}

			switch operator {
			case "-":
				if total == 0 && i == rangeE {
					total = loopVal
				} else {
					if total > loopVal {
						total = total - loopVal
					} else {
						total = loopVal - total
					}
				}
			case "/":
				total = total / loopVal
			default:
				fmt.Println("[error]: Does not supported that action")
				isDone = true
			}

			if isDone {
				break
			}
		}
	} else {
		for index, value := range input_value_list {
			loopVal, errorVal := transformStr2Int(value)
			isDone := false

			if errorVal != nil {
				fmt.Println("[error]:", errorVal)
				break
			}

			switch operator {
			case "+":
				total = total + loopVal
			case "*":
				if index == 0 {
					total = 1
				}
				total = total * loopVal
			default:
				fmt.Println("[error]: Does not supported that action")
				isDone = true
			}

			if isDone {
				break
			}
		}
	}

	fmt.Println("[total]:", total)
}
