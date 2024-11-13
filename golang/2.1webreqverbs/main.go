package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myurl = "http://localhost:8000/get"

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
	// PerformPostJsonRequest()
	// PerformPostFormRequest()
}

func PerformGetRequest() {
	response, _ := http.Get(myurl)
	defer response.Body.Close()
	// response. ContentLength, body, StatusCode
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	responseString.Write(content)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	// so this basically typecasting of our json format so that it can be comaptible with https protocol
	requestBody := strings.NewReader(`     
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)
	response, _ := http.Post(myurl, "application/json", requestBody)
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

// used to send headers of the request too
func PerformPostFormRequest() {
	data := url.Values{} // make map for the values which you want to send over URL
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")
	response, _ := http.PostForm(myurl, data)
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
