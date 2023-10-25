package main

import "fmt"

var functionsText string = "This text can be used in main.go file since both functions.go and main.go use the same package 'main'"

// note that below we have named the global function as PrintData.. uppercase 'P' this denotes that PrintData's visibility is 'public' whereas if it was printData then its visibility would've been 'private' to the package. Here, PrintData is public to other packages.

func PrintData(){
	fmt.Println("Printing from functions.go file")
}
