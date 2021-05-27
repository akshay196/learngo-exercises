// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"
// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier
	output := "Your name is %s and your lastname is %s.\n"
	var (
		first string
		last  string
	)
	fmt.Println("Enter your first name:")
	fmt.Scanln(&first)
	fmt.Println("Enter your last name:")
	fmt.Scanln(&last)
	
	fmt.Printf(output, first, last)
}
