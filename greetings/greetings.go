package greetings

import (
    "errors"
	"fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message
    if name == "" {
        return "", errors.New("empty name")
    }
    
    // In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). 
    // Taking the long way, you might have written this as:
    // var message string
    // message = fmt.Sprintf("Hi, %v. Welcome!", name)

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error){
    // A map to associate names with messages
    messages := make(map[string]string) // The make built-in function allocates and initializes an object of type slice, map, or chan
    // Loop through the recieved slice of names, calling
    // the Hello function to get a message for each name

    // In this for loop, range returns two values: 
    // the index of the current item in the loop and a copy of the item's value. 
    // You don't need the index, so you use the Go blank identifier (an underscore) to ignore it
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }

    return messages, nil
}

// returns one of a set of greeting messages. The returned message is selected at random
// randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string {
    // A slice of message formats
    formats := []string{
        "Hi %v, Welcome!",
        "Great to see you %v!",
        "Hail, %v! Well met!",
    }
    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}





