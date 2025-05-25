## Go Language

### Package

A package is a way to group functions, and it's made up of all the files in the same directory.

### Main Function

A main function executes by default when you run the main package.

### Exported Name

a Function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name
a Function whose name starts with a lowercase letter can be called only by a function in the same package. Not exported.

### := Operator

the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). 
Taking the long way, you might have written this as:

var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)

### Slice

A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types. 
When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed. 
Link: https://go.dev/blog/slices-intro

### Map

A map is a set of key/value pairs. The key is a unique identifier for the pair. The value is the data associated with that key.
In Go, you initialize a map with the following syntax: make(map[key-type]value-type). E.g. messages := make(map[string]string)
Link: https://go.dev/blog/maps

### Range

The range keyword is used to iterate over items in a slice or map.

### Blank Identifier

The blank identifier is an underscore (_). It's a special identifier that doesn't refer to a value.

### Fatal

The Fatal function is a built-in function that logs a message to the console and then calls os.Exit(1). log.Fatal("fatal message")

### Printf

The Printf function is a built-in function that prints a formatted string to the console. fmt.Printf("formatted string %v", variable)

### Unit Testing

Ending a file's name with _test.go tells the go test command that this file contains test functions. 
Test function names have the form TestName, where Name says something about the specific test. 
Also, test functions take a pointer to the testing package's testing.T type as a parameter. 
You use this parameter's methods for reporting and logging from your test. 
The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go). 
You can add the -v flag to get verbose output that lists all of the tests and their results. 
