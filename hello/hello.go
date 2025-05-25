package main //  Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory). 

import (
	"fmt" // Import the fmt package (fmt is a package that contains functions for formatting text, including printing to the console).
	"rsc.io/quote" // Import the quote package (quote is a package that contains functions for generating quotes).
	"math/rand" // Import the math/rand package (rand is a package that contains functions for generating random numbers).
	"example.com/greetings" // Import the greetings package (greetings is a package that contains functions for generating greetings).
)

func main() { // Declare a main function (a function is a block of code that performs a specific task). 
	fmt.Println(quote.Go())
	fmt.Println(rand.Intn(100))
	fmt.Println(greetings.Hello("Gladys"))
}

// A main function executes by default when you run the main package.
