package main

import (
	"fmt"
	"os"

	"frontendmasters.com/go/basics/section-2.2/fileUtils"
)

func main() {

	rootPath, _ := os.Getwd() // gets the root directory and stores it in the 'rootPath' variable, we don't really require the returned error value from the Getwd() method so let's store it in '_'

	fmt.Println(rootPath)

	content, err := fileUtils.ReadTextFile(rootPath + "/data/go.txt")

	if err != nil {
		fmt.Printf("Error %v", err)
	} else {
		fmt.Println(content)
	}
}
