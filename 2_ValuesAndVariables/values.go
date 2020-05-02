package main

/**
In order to see the medium go variable we can go to this link:

https://gist.github.com/thatisuday/c17e05de591c2e2021ab402e4c2d4bdc#file-medium-go-variables-data-types-csv

*/

/*

	The first thing we need to know is that Go doesn't have
	null or undefined, insted it has ZERO-VALUES, that means that even is a variable is not initializaed it will hold a zero value.

*/

import "fmt"

func main() {

	// To declare a variable we need to use this syntax
	// var variable name [dataType] [= initialValue]

	/*
		var integer1 int = 15
		var integer2 int8 = -25
		var integer3 int32 // default 0
		var float1 float32 = 63.2
		var string1 string = "Hello World!"
		var boolean1 bool // default false
		var boolean2 bool = true

		// The current variable name conventin is the camelCase
		// Event thougth Go is statically typed lenguage go provides the inference of the type declaretion

		var integer1A = 52           // int
		var string1A = "Hello World" // string
		var boolean1A = false        // bool

		// go also have a short-hand notation where you don't need to use var nor dataType, but this short-hand notation only works inside functions

		integer1B := 52           //int
		string1B := "Hello World" //string
		boolean1B := false        //bool

		// We can also have multiple var declaretion
		var var1, var2, var3 int
		var var1B, var2B, var3B int = 1, 2, 3
		var var1, var2, var3 = 1, 2.2, false
		var1, var2, var3 := 1, 2.2, false
	*/

	var1 := 10
	var2 := 8.6

	varString := "Hello"

	fmt.Println(var1 + int(var2))
	fmt.Println([]int32(varString))

	fmt.Println("This is a new line")

	/*
		There is some other topic related to variables:
			- Defining a type:
				type newType fromType

			- Type Alias:
				type aliasName = oldType
	*/

	// in go we also have constatns and this should be declaru during compilation time

	const constante = 3

	const (
		a = 1 // a == 1
		b = 2 // b == 2
		c     // c == 2
		d     // d == 2
	)

	fmt.Println(a, b, c, d, constante)

	const (
		ab = iota // a == 0
		bb = iota // b == 1
		cb = iota // c == 2
		db        // d == 3 (implicitely d = iota)
		_
		eb
	)

	fmt.Println(ab, bb, cb, db, eb)

	// one thing that it is importante to remember is that numeric expressions  will only have as a result the type of the operans

	numericExpression := 11 / 2
	fmt.Println(numericExpression)

}
