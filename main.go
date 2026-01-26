package main

import (
	"fmt"
	"os"
	"strconv"

	"ascii_art/Lib/check"
	"ascii_art/Lib/print"
)

func main() {
	input := os.Args

	data, err := check.Args(input)

	if !err {
		fmt.Println(data)
	}
	someStr := []string{}
	for x := 0; x < 95; x++ {
		someStr = append(someStr, strconv.Itoa(x))
	}
	listToPrint := print.AsciiArt(data, someStr)

	for _, str := range listToPrint {
		fmt.Print(str)
	}
	fmt.Print("\n")
}
