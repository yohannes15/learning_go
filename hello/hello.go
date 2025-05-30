package main //  Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory). 

import (
	"fmt" // Import the fmt package (fmt is a package that contains functions for formatting text, including printing to the console).
	"log"
	"example.com/greetings" // Import the greetings package (greetings is a package that contains functions for generating greetings).
)

func main() { // Declare a main function (a function is a block of code that performs a specific task). 
	// set properties of the predefined Logger. 
	log.SetPrefix("greetings: ") // log prefix for each log entry
	log.SetFlags(0) // without a time stamp or source file information. 

	message, err := greetings.Hello("Gladys") // Request a greeting message
	if err != nil {
		log.Fatal(err) // If an error was returned print it to the console and exit the program
	}

	fmt.Printf("single hello %v", message)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("multiple hello %v", messages)
}

// A main function executes by default when you run the main package.
