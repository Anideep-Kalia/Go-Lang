package main

import "fmt"

func adder(values ...int) (int, string) { // when you know that all inputs are integer but don't know number of them
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "this is second output"
}

func main() {
	fmt.Printf("Functions started here \n")
	intres, _ := adder(2, 3, 4, 5, 6)
	_, stringres := adder(2, 3, 4, 5, 6)
	fmt.Println(intres, stringres)
}
