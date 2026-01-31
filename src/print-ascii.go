package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "github.com/romance-dev/ascii-art"
	_ "github.com/romance-dev/ascii-art/fonts" // load all fonts
)

func main() {
	print()
}

func print() {
	allTheArgs := os.Args[1:]

	if len(allTheArgs) == 0 {
		fmt.Println("No input provided on what to print!")
		return
	}

	NormalString := strings.Join(allTheArgs, "")

	myFigure := asciiart.NewFigure(NormalString, "", false)
	myFigure.Print()
}
