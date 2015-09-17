package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Printf("Now you have %g problems. %f", math.Nextafter(2, 3), math.Nextafter(5, 5))

	// reads user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter your first number: ")
	first, _ := reader.ReadString('\n')

	// remove new line chars
	var firsttrim = trim(first, "\r\n")
	num1, _ := strconv.ParseFloat(firsttrim, 64)

	fmt.Printf("Enter your operand: ")
	op, _ := reader.ReadString('\n')

	var optrim = trim(op, "\r\n")
	var vo = validOp(optrim)
	for !vo {
		fmt.Printf("Invalid operand, enter a new one: ")
		op, _ := reader.ReadString('\n')
		optrim = trim(op, "\r\n")
		vo = validOp(optrim)
	}

	fmt.Printf("Enter your second number: ")
	second, _ := reader.ReadString('\n')

	// remove  new line chars
	var secondtrim = strings.Trim(second, "\r\n")
	num2, _ := strconv.ParseFloat(secondtrim, 64)

	fmt.Printf("Result: " + callCalculation(optrim, num1, num2))
	//fmt.Printf("Result: " + strconv.FormatFloat(result, 'f', -1, 64))

	/*fmt.Printf("Add: " + strconv.FormatFloat(addition(2, 5), 'f', -1, 64))
	fmt.Printf("\nSub: " + strconv.FormatFloat(subtract(5, 2), 'f', -1, 64))
	fmt.Printf("\nMult: " + strconv.FormatFloat(multiply(5, 2), 'f', -1, 64))
	fmt.Printf("\nDiv: " + strconv.FormatFloat(divide(10, 2), 'f', -1, 64))
	fmt.Printf("\nMod: " + strconv.FormatInt(modulo(10, 2), 10))*/
}

func callCalculation(op string, a, b float64) string {
	if op == "+" {
		return convertFloat64ToString(addition(a, b))
	} else if op == "-" {
		return convertFloat64ToString(subtract(a, b))
	} else if op == "*" {
		return convertFloat64ToString(multiply(a, b))
	} else if op == "/" {
		return convertFloat64ToString(divide(a, b))
	} else if op == "%" {
		return convertInt64ToString(modulo(int64(a), int64(b)))
	}
	return "Couldn't calculate correctly."
}

func validOp(op string) bool {
	// slice declared without element count
	validOps := []string{"+", "-", "*", "/", "%"}
	for _, validOp := range validOps {
		if validOp == op {
			return true
		}
	}
	return false
}

func convertFloat64ToString(float float64) string {
	return strconv.FormatFloat(float, 'f', -1, 64)
}

func convertInt64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func trim(value, cutset string) string {
	return strings.Trim(value, cutset)
}

// Go has two floating point types: float32 and float64 (also often referred to as single precision and double precision respectively)
func addition(first, second float64) float64 {
	return first + second
}

func subtract(first, second float64) float64 {
	return first - second
}

func multiply(first, second float64) float64 {
	return first * second
}

func divide(first, second float64) float64 {
	return first / second
}

func modulo(first, second int64) int64 {
	return first % second
}
