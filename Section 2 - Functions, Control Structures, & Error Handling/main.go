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

func main(){
	sum := add(5,3)
	additionResult, subtractionResult := addAndSubtract(5,3) // 2 variables are declared on the fly to store both the returned values from the 'addAndSubtract' function

	stateTax, cityTax := tax(100)
	stateTax2, _ := tax(200) // if you don't wish to use a value returned by a function, you can use an  underscore '_', this way you don't end up creating an identifier to store the other value returned from the function.

	fmt.Println(sum)
	fmt.Println(additionResult, subtractionResult)

	fmt.Println("State tax is: ", stateTax)
	fmt.Println("City tax is: ", cityTax)
	fmt.Println(stateTax2)

}





