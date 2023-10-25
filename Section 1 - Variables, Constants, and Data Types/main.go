package main

// global variables

const url = "https://frontendmasters.com"

func main() {
	print("Hello World!\n")

	// variables are initialized to their zero value, which is: 0 for numeric types, false for the boolean type, and "" (the empty string) for strings.; while pointers are initialized to nil.

	// function-scoped variables

	var x int 
	var name string

	print(x, name)

	x = 3
	name = "Basith"

	print(x, name)

	const y = 2 // constants can only be bool, string or number

	var z int = 3
	var text string = "Going Gaga over Go"

	print(y, z, text)
}