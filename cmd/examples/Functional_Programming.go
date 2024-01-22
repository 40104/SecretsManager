package main

// Import packages
import (
	"fmt"
)

// Main function
func main() {
	// Only final data structures
	//------------------------------------------------------------------------------------------------------------------
	// Create a constant that stores an immutable value of type string
	const str = "Hello World!!!!"

	// Trying to change the value of a constant
	str = "I am Go" // Compilation error: cannot assign to name

	// Create a variable that stores an immutable value of type slice
	var slice = []int{1, 2, 3}
	slice = append(slice, 4) // Runtime error: panic: runtime error: slice bounds out of range
	//------------------------------------------------------------------------------------------------------------------

	// Side-effect-free functions (Mostly) 
	//------------------------------------------------------------------------------------------------------------------
	// Function that checks a number for even parity
	Is_Even := func (number int) (string) {
		if number % 2 == 0 {
			return "Even" // No side effects, just return value
		} else {
			return "Not even" // No side effects, just return value
		}
	}
	fmt.Println(Is_Even(9)) // Call this function
	fmt.Println(Is_Even(10)) // Call this function
	//------------------------------------------------------------------------------------------------------------------

	// The use of higher-order functions
	//------------------------------------------------------------------------------------------------------------------
	// A function that takes a function as an argument, applies it, and returns its value.
	apply_function := func (a int, b int, f func(int, int) (int)) (int){
		return f(a,b)
	}
	// A function that returns a function as a result
	return_function := func () (func(int,int) (int)) {
		return func(a int, b int) int { return a + b } // Return an anonymous function
	}
	//Initialization of variables and functions
	var a int = 5
	var b int = 10
	fmt.Println(apply_function(a, b, return_function())) // Get an anonymous function and pass it as an argument
	//------------------------------------------------------------------------------------------------------------------

	// Functions as parameters and return values
	//------------------------------------------------------------------------------------------------------------------
	// Function that multiplies 2 numbers
	multiplication := func (x int, y int) (int) {
		return x * y
	}
	// A function that receives numbers and a function as input, executes it and, depending on the result, returns one of the functions
	check := func (x int, y int, f func (int, int) int) (func(int) (int)) {
		result := f(x, y) //Calculate the value
		if result >= 10 {
			return func (number int) (int) { return number - 5} //Return the function
		} else {
			return func (number int) (int) { return number * 2} //Return the function
		}
	}
	
	//Initialization of variables and functions
	var x int = 3
	var y int = 5
	var z int = 7
	fn := check(x, y, multiplication) // Get an anonymous function
	fmt.Println(fn(z)) // Call it as an argument
	//------------------------------------------------------------------------------------------------------------------

	// Use closures / anonymous functions
	//------------------------------------------------------------------------------------------------------------------
	// A function that returns a function that returns the square of its argument
	square := func () (func(int) int) {
		return func (x int) (int) { return x * x } // Return an anonymous function
	}
	
	// Get the square function as the return value of the square function
	sqr := square()
	fmt.Println(sqr(5)) // Call the sqr function
	//------------------------------------------------------------------------------------------------------------------
}

