package main //  Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory). 

import (
	"fmt" // Import the fmt package (fmt is a package that contains functions for formatting text, including printing to the console).
	"log"
	"rsc.io/quote" // Import the quote package (quote is a package that contains functions for generating quotes).
	"math/rand" // Import the math/rand package (rand is a package that contains functions for generating random numbers).
	"example.com/greetings" // Import the greetings package (greetings is a package that contains functions for generating greetings).
)

func main() { // Declare a main function (a function is a block of code that performs a specific task). 
	// set properties of the predefined Logger. 
	log.SetPrefix("greetings: ") // log prefix for each log entry
	log.SetFlags(0) // without a time stamp or source file information. 

	message, err := greetings.Hello("") // Request a greeting message
	if err != nil {
		log.Fatal(err) // If an error was returned print it to the console and exit the program
	}

	fmt.Println(quote.Go())
	fmt.Println(rand.Intn(100))
	fmt.Println(message)
}

// A main function executes by default when you run the main package.
