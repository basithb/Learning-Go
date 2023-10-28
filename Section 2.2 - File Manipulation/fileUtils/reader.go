package fileUtils

import "os"

func ReadTextFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)

	if (err != nil){
		// could not read the file
		return "", err
	} else {
		return string(content),nil // here content is a [] bytes (a slice of type bytes); so we cannot return it like that, as a result we explicitly cast the type to string and return it.
	}
}