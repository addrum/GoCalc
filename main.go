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
	var trim = strings.Trim(first, "\r\n")
	num1, _ := strconv.ParseFloat(trim, 64)

	fmt.Printf("Enter your second number: ")
	second, _ := reader.ReadString('\n')

	// remove  new line chars
	trim = strings.Trim(second, "\r\n")
	num2, _ := strconv.ParseFloat(trim, 64)

	var result = addition(num1, num2)

	fmt.Printf("Result: " + strconv.FormatFloat(result, 'f', -1, 64))

	/*fmt.Printf("Add: " + strconv.FormatFloat(addition(2, 5), 'f', -1, 64))
	fmt.Printf("\nSub: " + strconv.FormatFloat(subtract(5, 2), 'f', -1, 64))
	fmt.Printf("\nMult: " + strconv.FormatFloat(multiply(5, 2), 'f', -1, 64))
	fmt.Printf("\nDiv: " + strconv.FormatFloat(divide(10, 2), 'f', -1, 64))
	fmt.Printf("\nMod: " + strconv.FormatInt(modulo(10, 2), 10))*/
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
