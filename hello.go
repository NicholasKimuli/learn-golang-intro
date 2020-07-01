package main

import (
	"fmt"
	"errors"
	"math"
)

// Based off of the following intro video: https://www.youtube.com/watch?v=C8LgvuEBraI
func main() {
	// PrintHelloWorld()
	// getSum := sum( 6, 10 )
	// fmt.Println( "Sum:", getSum )
	// declareAndInitializeVariables()
	// ifStatements( 4 )
	// arrays()
	// maps()
	// loopingOverArraysAndMaps()
	
	// result, err := sqrt( -16 )

	// if err != nil {
	// 	fmt.Println( err )
	// } else {
	// 	fmt.Println( result )
	// }

	// workingWithStructs()
	workingWithPointers()
}

func PrintHelloWorld() {
	fmt.Println( "Helloworld!" )
}

func declareAndInitializeVariables() {
	// Method 1. Use [(var) keyword] [name of the variable] [type of variable] = [Value]
	var rk int = 10

	// Method 2. [variable name] := [value]
	// := is used because Golang can infer the type under the hood
	// Neater to read
	nk := "Captain Sunshine"

	fmt.Println( "NK:", nk, ", RK:", rk)
}

func ifStatements( x int ) {
	if x > 6 {
		fmt.Println( "More than 6" )
	} else if x <= 2 {
		fmt.Println( "Less than or equal to 2" )
	} else {
		fmt.Println( "Not More than 6 & Not Less than or equal to 2" )
	}
}

func arrays() {
	// Contains a fixed number of elements
	// If a variable is not initializes, it is give the types zero value. For ints, that's 0
	var a [5] int
	// Set value for index 2
	a[ 2 ] = 7
	// Using short-hand syntax to initialize array
	b := [ 5 ] int { 5, 4, 3, 2, 1 }
	// Use slices to overcome the fixed number of elements limitation
	// Don't specify the array length to create a slice
	c := [] int { 10, 9, 8, 7, 6 }
	// Use the append function to add more elements to a slice
	c = append( c, 14 )
	fmt.Println( "a:", a )
	fmt.Println( "b:", b )
	fmt.Println( "c:", c )
}

func maps() {
	// Are similar to associative arrays in PHP
	// They hold key:value pairs
	vertices := make( map [ string ] int )

	vertices[ "triangle" ] = 2
	vertices[ "square" ] = 3
	vertices[ "dodecagon" ] = 12

	fmt.Println( "Before:", vertices )

	// Removing an entry
	delete( vertices, "square" )

	fmt.Println( "After:", vertices )
}

func loops() {
	// The only type of loop available is the for loop
	fmt.Println( "For loop" )

	for i := 0; i < 5; i++ {
		fmt.Println( "i:", i )
	}

	// Doubles as a while loop
	fmt.Println( "While loop" )

	j := 0

	for j < 5 {
		fmt.Println( "j:", j )
		j++
	}
}

func loopingOverArraysAndMaps() {
	arr := [] string{ "a", "b", "c" }

	// Use the range keyword to loop over arrays and maps
	fmt.Println( "Looping over array" )
	for index, value := range arr {
		fmt.Println( "index:", index, "value:", value )
	}

	fmt.Println( "Looping over map" )
	
	m := make( map[ string ] string )
	m[ "a"] = "alpha"
	m[ "b" ] = "beta"

	for key, value := range m {
		fmt.Println( "key:", key, "value:", value )
	}
}

func sum( x int, y int ) int {
	return x + y
}

func sqrt( x float64) ( float64, error ) {
	// Function can return multiple types
	if x < 0 {
		return 0, errors.New( "Undefined for negative numbers" )
	}

	// Golang doesn't have exceptions so you have to handle errors yourself. The error package comes in handy for those.
	return math.Sqrt( x ), nil
}

// Structs - collection of fields that allow you to collect things(could be variables or other structs) together
type person struct {
	name string
	age int
}

func workingWithStructs() {
	p := person{ name: "Nick", age: 26 }
	fmt.Println( "Print created struct:" )
	fmt.Println( p )
	fmt.Println( "Print specific struct value:" )
	// Use dot notation for this
	fmt.Println( "Name:", p.name )
}

func workingWithPointers() {
	// A pointer is the memory address of where a variable is stored
	// Wikipedia: is a programming language obj that stores a memory address
	i := 7
	fmt.Println( "Get pointer:", &i ) 
	inc( &i )
	fmt.Println( "Increment value:", i)
}

func inc( x *int ) {
	// This function is able to look at the value of the variable in that pointer and increment the original value
	*x++
}