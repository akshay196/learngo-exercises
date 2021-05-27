// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------
import "fmt"

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	phone := map[string]string{
		"Gaikwad": "333999000",
		"Ghodake": "222003333",
	}

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	avail := map[int]bool{
		1: true,
		2: false,
		3: true,
		4: false,
		5: true,
	}

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	contacts := map[string][]string{
		"Joshi": {"123", "345"},
		"Gandhi": {"677", "234"},
	}

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	basket := map[int]map[int]int{
		51: {1: 1, 2: 0, 3: 2},
		52: {1: 5, 2: 0, 3: 1},
	}
	fmt.Printf("%#v\n", phone)
	fmt.Printf("%#v\n", avail)
	fmt.Printf("%#v\n", contacts)
	fmt.Printf("%#v\n", basket)
}
