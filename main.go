package main

import (
	"os"
	"ascii_art/Lib/check"
	"fmt"
)

func main() {
	input := os.Args
	
	data, err := check.Args(input)

	if !err {
		fmt.Println(data)
	}
	// add a function
	// ascii-art.printAsciiArt()
	// and use it 
}
