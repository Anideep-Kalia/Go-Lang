package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ways of declaration of the variables
const Globaly string = "so the capital letter starting and outisde the main function"

func main() {
	var website = "this is without defining type"
	numberofuser := "no use of var but not allowed outside main()"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the input you need to: ")
	// comma ok || err syntax
	// input , err :=reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	// now trying to add 1 to the input given by the user the trimspace is actually triming the extra line that we are inputing accidentaly
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("rating with extra +1", numrating+1)
		fmt.Println("original input", input)
	}
	fmt.Println(numberofuser)
	fmt.Println(Globaly)
	fmt.Println(website)
}
