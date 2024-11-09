package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"sort"
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
	fmt.Print("No next line will be made after this so adding \n")
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

	//
	fmt.Println("\n", "---------POINTERS------------")
	myNumber := 23
	var ptr = &myNumber
	println(ptr)
	println(*ptr)
	println(*ptr + 2)

	//
	fmt.Println("\n", "---------ARRAYS & OPERATIONS------------")
	// INITIALISATIONS
	// var fruitList [4]string
	fruitList := [4]string{"aple"}
	fruits2 := []string{"apple", "banana", "kiwi", "chiku", "aam"}
	fruits3 := []string{}
	fruits5 := make([]int, 5)

	// FUNCTIONS
	fruitList[0] = "Apple"
	fruits5 = []int{1, 2, 56, 7, 8}
	fruits3 = append(fruits3, "added some fruits")
	fruits4 := append(fruits2[:3])                // gives first 3 elements
	fruits5 = append(fruits5[:1], fruits5[2:]...) // spliting array
	sort.Ints(fruits5)

	// PRINTING
	fmt.Println("Fruit list is: ", len(fruitList))
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is:", fruits2)
	fmt.Println("Fruit list is:", fruits3)
	fmt.Println("Fruit list is:", fruits4)
	fmt.Println("Fruit list is:", fruits5)

	//
	fmt.Println("\n", "---------MAPS------------")
	maping := make(map[int]string)
	maping[0] = "first element"

	delete(maping, 2)

	// Traversing
	fmt.Println(maping[2], maping[0])
	for it, value := range maping {
		fmt.Printf("%v -> %v", it, value)
	}
}
