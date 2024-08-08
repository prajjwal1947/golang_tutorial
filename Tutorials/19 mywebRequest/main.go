package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummy.restapiexample.com/api/v1/employees"

func main() {
	fmt.Println("welcome to web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // defer closing the connection

	if response.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("failed to fetch data: %s", response.Status))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type: %T\n", response)
	fmt.Println(string(body))
}
