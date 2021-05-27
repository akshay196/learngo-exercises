// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}
	for _, s := range strings {
		fmt.Printf("%q\n", s)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\thas %d bytes and %d runes\n", len(s), utf8.RuneCountInString(s))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes     : % x\n", s)
		
		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\trunes     : ")
		for _, x := range s {
			fmt.Printf("%x ", x)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Printf("\trunes     : ")
		for _, x := range s {
			fmt.Printf("%q ", x)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst     : %c (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tlast      : %c (%d bytes)\n", r, size)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:])
		fmt.Printf("\tfirst2    : %q\n", string(s[:first+second]))

		// Slice and print the last two runes of the strings
		_, last := utf8.DecodeLastRuneInString(s)
		_, secondLast := utf8.DecodeLastRuneInString(s[:len(s)-last])
		fmt.Printf("\tlast2     : %q\n", string(s[len(s)-(last+secondLast):]))

		// Convert the string to []rune
		// Print the first and last two runes
		runes := []rune(s)
		fmt.Printf("\tfirst2    : %q\n", string(runes[:2]))
		fmt.Printf("\tlast2     : %q\n", string(runes[len(runes)-2:]))
	}
}
