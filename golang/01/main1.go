package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Not so Regular things to use \n")

	fmt.Println("-----------DATE-------------")
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

	fmt.Println("")
}
