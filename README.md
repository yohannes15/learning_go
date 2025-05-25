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
