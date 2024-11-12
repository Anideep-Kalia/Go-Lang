package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl) // returns pointer

	// fmt.Println(result.Scheme)	=> https
	// fmt.Println(result.Host)		=> lco.dev:3000
	// fmt.Println(result.Path)		=> /learn
	// fmt.Println(result.Port())	=> 3000
	// fmt.Println(result.RawQuery)	=> coursename=reactjs&paymentid=ghbj456ghb

	qparams := result.Query() // makes map of strings i.e. making map of params of url
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"]) // gives reactjs as Query() fucntion was used

	for _, val := range qparams { // iterates over every parameter of the URL
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{ // as the we need pointer in case of url what we have done by parse function
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
