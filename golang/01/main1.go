package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Not so Regular things to use \n")

	//
	fmt.Println("-----------DATE-------------")
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	//
	fmt.Println("-----------STRUCT-------------")
	Anni := User{"Anideep", "example.com", true, 100}
	fmt.Printf("%+v \n %v \n %v", Anni, Anni, Anni.Name)

	//
	fmt.Println("-----------SWITCH-CASE-------------")
	num := 5
	switch num {
	case 1:
		fmt.Println("Hello there is 1 here")
	case 5:
		fmt.Println("Here is 5")
	default:
		fmt.Println("This is default")
	}

	//
	fmt.Println("\n", "---------DEFER------------")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	//
	fmt.Println("\n", "---------FILE-MANAGEMENT------------")
	file, _ := os.Create("./mylcogofile.txt")
	length, _ := io.WriteString(file, "Writen something which is not meant to be")
	fmt.Println("length is: ", length)
	defer file.Close()
	res, _ := os.ReadFile("./mylcogofile.txt")
	fmt.Println(string(res))
	// fmt.Println(res)

}
