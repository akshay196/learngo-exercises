// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Assign the Arrays
//
//  1. Create an array named books
//
//  2. Add book titles to the array
//
//  3. Create two more copies of the array named: upper and lower
//
//  4. Change the book titles to uppercase in the upper array only
//
//  5. Change the book titles to lowercase in the lower array only
//
//  6. Print all the arrays
//
//  7. Observe that the arrays are not connected when they're copied.
//
// NOTE
//  Check out the strings package, it has functions to convert letters to
//  upper and lower cases.
//
// BONUS
//  Invent your own arrays with different types other than string,
//  and do some manipulations on them.
//
// EXPECTED OUTPUT
//   Note: Don't worry about the book titles here, you can use any title.
//
//   books: ["Kafka's Revenge" "Stay Golden" "Everythingship"]
//   upper: ["KAFKA'S REVENGE" "STAY GOLDEN" "EVERYTHINGSHIP"]
//   lower: ["kafka's revenge" "stay golden" "everythingship"]
// ---------------------------------------------------------
import (
	"fmt"
	"strings"
)

func main() {
	books := [...]string{"Kafka", "Stay Golden", "Everythingship"}
	var ubooks [3]string
	var lbooks [3]string
	var prices [3]float64 = [3]float64{120, 340, 960.45}

	for i:=0; i<len(books); i++ {
		ubooks[i] = strings.ToUpper(books[i])
		lbooks[i] = strings.ToLower(books[i])
	}

	fmt.Printf("Before price = %#v\n", prices)

	// Increase price by 40 each
	for i:=0; i<len(prices); i++ {
		prices[i] += 40
	}

	fmt.Printf("books = %#v\n", books)
	fmt.Printf("ubooks = %#v\n", ubooks)
	fmt.Printf("lbooks = %#v\n", lbooks)
	fmt.Printf("After price = %#v\n", prices)
}
