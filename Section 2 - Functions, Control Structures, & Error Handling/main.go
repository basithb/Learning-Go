package main

import "fmt"

func save() {
	fmt.Println("Hello")
}

func saveText (text string){
	fmt.Println(text)
}

// notice how we explicitly mention the type of the parameter and also mention the type of the value returned

func add(a int, b int) int { 
	return a + b	
}

// in Go, a function can return more than one value and we have also mentioned the type for both these values

func addAndSubtract(a int, b int) (int, int) {
	return a + b , a - b
} 

func tax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
} 

func taxWithName(price float32) (stateTax float32, cityTax float32){ // we are naming the returned values
	stateTax = price * 0.09
	cityTax = price * 0.02
	return // we don't have to explicitly mention the return values since it's already mentioned up top.
}

func birthday (pointerAge *int){

	if(*pointerAge > 60){
		panic("Too old to be true") 
		
		// panic means an unexpected condition arises in your Go program due to which the execution of your program is terminated.

		// any deferred functions are executed normally during the event of a panic.
	}

	fmt.Printf("The pointer is %v and the pointer value is %v\n", pointerAge, *pointerAge) // fmt.Printf is used for template strings; %d for integers, %s for string and %v for values
	*pointerAge += 1
	
}

func main(){

	defer fmt.Println("first statement inside main() however due to 'defer' keyword it gets called towards the end")
	defer fmt.Println("another defer")

	// defer is commonly used to simplify functions that perform various clean-up actions. you can have multiple statements with defer.

	fmt.Println("Main function output begins here")
	sum := add(5,3)
	additionResult, subtractionResult := addAndSubtract(5,3) // 2 variables are declared on the fly to store both the returned values from the 'addAndSubtract' function

	stateTax, cityTax := tax(100)
	stateTax2, _ := tax(200) // if you don't wish to use a value returned by a function, you can use an  underscore '_', this way you don't end up creating an identifier to store the other value returned from the function.

	fmt.Println(sum)
	fmt.Println(additionResult, subtractionResult)

	fmt.Println("State tax is: ", stateTax)
	fmt.Println("City tax is: ", cityTax)
	fmt.Println(stateTax2)

	age := 22
	birthday(&age) 

	// why are we using pointers above? in Go, arguments are passed by value..  we would only be changing the copy of 'age' variable and not the variable itself. by doing this we are pointing to the original 'age' variable which we desire to mutate.
	
	fmt.Println("New age:", age)
}





