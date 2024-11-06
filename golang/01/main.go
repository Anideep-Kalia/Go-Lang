package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	// DATATYPES
	fmt.Println("\n", "----------DATATYPES---------")
	var username string = "hitesh"
	var isLoggedIn bool = false
	var smallVal uint8 = 255
	var smallFloat float64 = 255.45544511254451885
	var anotherVariable int
	var website = "learncodeonline.in"
	numberOfUser := 300000.0

	// PRINT OPTIONS
	fmt.Println("\n", "----------Print---------")
	fmt.Printf(username, "\n", isLoggedIn, "\n", smallVal, "\n", smallFloat, "\n", anotherVariable, " \n", website, "\n", "%T", numberOfUser, "\n")
	fmt.Println(username, "\n", isLoggedIn, "\n", smallVal, "\n", smallFloat, "\n", anotherVariable, " \n", website, "\n", "%T", numberOfUser, "\n")
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%T\n", username, isLoggedIn, smallVal, smallFloat, anotherVariable, website, numberOfUser)

	//
	fmt.Println("\n", "----------INPUT AND IF-ELSE---------")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // the data will be stored in input or it will show error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", input)
	}

	//
	fmt.Println("\n", "----------TYPECASTING AND STRINGS---------")
	var mynumberTwo float64 = 6.5
	var mynumberOne int = 2
	fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // converting input data into float64
	fmt.Println(numRating)

	//
	fmt.Println("\n", "----------MATH RANDOM AND CRYPTO---------")
	myRandomNum, error := rand.Int(rand.Reader, big.NewInt(5)) //will random numbers upto 5
	fmt.Println(myRandomNum)
	fmt.Println(error)
}
