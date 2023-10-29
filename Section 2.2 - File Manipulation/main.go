package main

import (
	"fmt"
	"os"

	"frontendmasters.com/go/basics/section-2.2/fileUtils"
)

func main() {

	price := 50.23

	stringPrice := fmt.Sprintf("%.2f", price) // Sprintf returns the resulting string; it does not print it to the console.

	fmt.Printf("Printing price in string format using Sprintf %v\n", stringPrice)

	rootPath, _ := os.Getwd() // gets the root directory and stores it in the 'rootPath' variable, we don't really require the returned error value from the Getwd() method so let's store it in '_'

	fmt.Println(rootPath)

	filePath := rootPath + "/data"
	content, err := fileUtils.ReadTextFile(filePath + "/go.txt")

	if err != nil {
		fmt.Printf("Error %v", err)
	} else {
		fmt.Println(content)
	}

	insertContent := "This is some sample text that will be inside of new.txt file"
	writeErr := fileUtils.WriteToFile(filePath+"/new.txt", insertContent)

	if writeErr != nil {
		fmt.Printf("Error %v", writeErr)

	} else {
		newTxtFileContent, newTxtFileErr := fileUtils.ReadTextFile(filePath + "/new.txt")

		if newTxtFileErr != nil {
			fmt.Println("Unable to read from file")
		} else {
			fmt.Println(newTxtFileContent)
		}

	}

}
