// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------
import (
	"os"
	"strings"
	"fmt"
)
const usage = "\t[command] [string]\n\n\tAvailable commands: lower, upper and title"
func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	cmd, str := args[1], args[2]
	var res string
	switch cmd {
	case "lower":
		res = strings.ToLower(str)
	case "upper":
		res = strings.ToUpper(str)
	case "title":
		res = strings.Title(str)
	default:
		res = fmt.Sprintf("Unknown command: %q", cmd)
	}
	fmt.Println(res)
}
