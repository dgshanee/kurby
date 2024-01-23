package main

import (
	"cs/kurby/dom"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid format")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "display":
		handleDisplay()
	default:
		fmt.Println("Invalid format")
		os.Exit(1)
	}
}

func handleDisplay() {
	dom.Display()
}
