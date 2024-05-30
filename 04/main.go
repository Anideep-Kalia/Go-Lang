package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	// MAPS
	fmt.Println("Maps section")
	languages := make(map[int]string)

	languages[1] = "python"
	languages[2] = "english"
	languages[3] = "hindi"
	languages[4] = "javascript"

	fmt.Printf("key and values are %v => %v", 1, languages[1])
	fmt.Printf("\n")
	fmt.Println("all the data in map : ", languages)

	delete(languages, 1)
	fmt.Println("all the data in map after deleting : ", languages)
	fmt.Printf("\n")

	// LOOPS
	fmt.Println("Loops start here")
	for key, value := range languages {
		fmt.Printf("for key %v we have value %v \n", key, value)
	}
	fmt.Printf("\n")

	// STRUCT IMPLEMENTATION
	fmt.Println("struct starts here")
	Anideep := User{"Anideep", "anideepkalia@gmail.com", true, 22}
	fmt.Printf("The details of the user are :  %+v\n", Anideep)
	fmt.Printf("\n")

	//IF ELSE
	fmt.Println("if-else starts here")
	loginCount := 23

	if loginCount < 10 {
		fmt.Println("number is less than 10")
	} else if loginCount < 20 {
		fmt.Println("number is less than 20")
	} else {
		fmt.Println("number is way too high")
	}
	fmt.Printf("\n")

	// switch case and random number
	fmt.Printf("\nswitch case starts here \n")
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println(dicenumber)
	switch dicenumber {
	case 1:
		fmt.Printf("dice number is 1 \n")
		fallthrough // so if 1 comes all cases will be triggered
	default:
		fmt.Printf("dice number is not 1 \n")
	}

	// while loop
	fmt.Println("WHILE LOOPS START HERE")
	idx := 1
	for idx < 10 {
		if idx == 5 {
			idx++
			continue
		} else {
			fmt.Println(idx)
			idx++
		}
	}
}
