package data

import "fmt"

// collections in Go

var Hobbies [5]string // arrays
var NewHobbies = [2]string{"Programming", "Football"} // arrays

var Slices []int // slices

var Codes map[string]int // map; key-value pair

// init functions are functions that initialize something when the program begins that start execution. You cna have more than one init() functions and init() functions are different form main() functions.

func init() {
	Hobbies[0] = "India"

	lengthHobbiesArray := len(Hobbies)
	capacityOfSlices := cap(Slices)

	fmt.Println(lengthHobbiesArray, capacityOfSlices)
}

func init(){
	fmt.Println("Another init function")
}