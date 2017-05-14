package hello

import (
	"fmt"
	"os"
)

//Hello world
func Hello() {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf("Hello world. I coudln't get my current workign directory. Instead i got an error %v", err)
	} else {
		fmt.Printf("Hello world. My current working directory = " + currentWorkingDirectory)
	}
}
