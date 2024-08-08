package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("welcome to go files")

	content := "this need to go to the files-learncode"
	file, err := os.Create("./myGoFile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Print(length)
	defer file.Close()
	readFile("./myGoFile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Print("text data inside file: \n", string(dataByte))

}
