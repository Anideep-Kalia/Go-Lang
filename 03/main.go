package main

import (
	"fmt"
	"time"
)

func main() {
	// Now we are going to talk about time
	presenttime := time.Now()
	createdtime := time.Date(2020, time.October, 12, 06, 56, 0, 0, time.Local)

	fmt.Println(presenttime)
	fmt.Println("formated present time => ", presenttime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(createdtime)
	fmt.Println("formated time created", createdtime.Format("01-02-2006 Monday"))
	fmt.Printf("\n")

	// POINTERS
	mynumber := 23
	var ptr = &mynumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", mynumber)
	fmt.Printf("\n")

	// Arrays
	var fruitlist [4]string
	// var fruitlist [4] {"apple", "beans", "orange"}  => another way to declare
	// fruitlist := make([]int,4) => another one

	fruitlist[0] = "apple"
	fruitlist[1] = "banana"
	fruitlist[3] = "orange"

	fmt.Println("fruitlist:", fruitlist)
	fmt.Println("the length of the fruit list Declared", len(fruitlist))
	fmt.Printf("\n")

	// Slices
	var listing = []string{"apple", "banana", "orange"}
	fmt.Printf("types of listing %T\n", listing)
	listing = append(listing, "mango", "peach")

	fmt.Println(listing)
	var aadhi = append(listing[1:])   // except first
	var aadhi2 = append(listing[1:3]) //excep first and before 4
	var aadhi3 = append(listing[:3])  // before 4
	var removing = append(listing[:2], listing[3:]...)
	fmt.Println(aadhi)
	fmt.Println(aadhi2)
	fmt.Println(aadhi3)
	fmt.Println(removing)
	fmt.Println("so the major advantage of using append in the array is we can add more number of elements than declared when the array is made")

	// Sorting
	// sort.Ints(listing) => listing contains integers only
	// sort.IntsAreSorted(listing) => true if sorted else vica versa
}
