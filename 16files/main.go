package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")
	content := "This needs to go in a file.\nHello"

	file, err := os.Create("./myfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println(length)
	readFile("./myfile.txt")
	defer file.Close()

}

func readFile(filePath string) {
	dataByte, err := ioutil.ReadFile(filePath)

	checkNilErr(err)

	fmt.Println(dataByte)
	fmt.Println(string(dataByte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
