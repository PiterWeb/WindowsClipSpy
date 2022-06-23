package main

import (
	"fmt"

	"github.com/PiterWeb/WindowsClipSpy/builder"
)

func main() {

	var Url string
	var Method string
	var Filename string

	
	fmt.Println("Url: ")
	fmt.Scanln(&Url)
	fmt.Println("Method: ")
	fmt.Scanln(&Method)
	fmt.Println("Filename: ")
	fmt.Scanln(&Filename)

	var Config = builder.Config{
		Url:      Url,
		Method:   Method,
		Filename: Filename,
	}

	builder.BuildClipboard(&Config)

}
