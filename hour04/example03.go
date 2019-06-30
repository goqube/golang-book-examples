package main

import "fmt"

func sumNumbers(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index)
		total += number
	}
	return total
}

func main() {
	result := sumNumbers(1, 2, 3, 4)
	fmt.Printf("The result is %v\n", result)
}
