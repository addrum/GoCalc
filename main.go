package main

import "fmt"
import "strconv"

var a = 0.0
var b = 0.0

func main() {
	fmt.Printf("This is me calculator!\nEnter your first number:\n")
	fmt.Printf("Add: " + strconv.FormatFloat(addition(2, 5), 'f', -1, 64))
	fmt.Printf("\nSub: " + strconv.FormatFloat(subtract(5, 2), 'f', -1, 64))
	fmt.Printf("\nMult: " + strconv.FormatFloat(multiply(5, 2), 'f', -1, 64))
	fmt.Printf("\nDiv: " + strconv.FormatFloat(divide(10, 2), 'f', -1, 64))
	fmt.Printf("\nMod: " + strconv.FormatInt(modulo(10, 2), 10))
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
