package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&payment"

func main() {
	fmt.Println("Welcome to URL parsing")

	fmt.Println(myUrl)

	// Parsing
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("These are the query params of type %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for key, val := range qparams {
		fmt.Println("Param key:", key, "Param value:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Constructed URL is:", anotherUrl)
}
