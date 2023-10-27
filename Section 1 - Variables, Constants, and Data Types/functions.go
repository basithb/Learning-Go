package main

import "fmt" // imports are on a per-file basis and not on per-package basis like how we transfer the variable below on to the other as it is contained within the same package.

import "frontendmasters.com/go/basics/section-1/data"

var functionsText string = "This text can be used in main.go file since both functions.go and main.go use the same package 'main'"

// note that below we have named the global function as PrintData.. uppercase 'P' this denotes that PrintData's visibility is 'public' whereas if it was printData then its visibility would've been 'private' to the package. Here, PrintData is public to other packages.

// if you want to export a variable, constant, or functions it must be must be TitleCase

func PrintData(){
	fmt.Println("Printing from functions.go file")
	fmt.Println("Printing MaxSpeed value from constants.go:", data.MaxSpeed)
}
