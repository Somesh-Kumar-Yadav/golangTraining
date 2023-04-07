package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(input)
		newInput, e := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(newInput + 1)
		}
	}

}
