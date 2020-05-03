package main

// it is important to mention during this section that what is the scope of a function

import "fmt"

// This is the main function, everytime we execute a go program, the go compiler will fill the main function in order to execute

func main() {
	fmt.Println(add(4, 3))
	fmt.Println(multi(4, 3))
	firstWord, secondWork := multiReturn("primerPalabra", "segundaPalabra")
	fmt.Println(firstWord, secondWork)
	fmt.Println(namedReturnValues(1))
}

// this is a simple function
func add(x int, y int) int {
	return x + y
}

// here we can see how to declare different paramenters for a function
func multi(x, y int) int {
	return x * y
}

// we can also return multiple results from a function
func multiReturn(x, y string) (string, string) {
	return y, x
}

// named return values

func namedReturnValues(firstValue int) (ret1, ret2 int) {
	ret1 = firstValue + 1
	ret2 = firstValue + 2
	return
}
