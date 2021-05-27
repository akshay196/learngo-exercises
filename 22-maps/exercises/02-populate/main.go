// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Populate and Lookup
//
//  Add elements to the maps that you've declared in the
//  first exercise, and try them by looking up for the keys.
//
//  Either use the `make()` or `map literals`.
//
//  After completing the exercise, remove the data and check
//  that your program still works.
//
//
//  1. Phone numbers by last name
//     --------------------------
//     bowen  202-555-0179
//     dulin  03.37.77.63.06
//     greco  03489940240
//
//     Print the dulin's phone number.
//
//
//  2. Product availability by Product ID
//     ----------------
//     617841573 true
//     879401371 false
//     576872813 true
//
//     Is Product ID 879401371 available?
//
//
//  3. Multiple phone numbers by last name
//     ------------------------------------------------------
//     bowen  [202-555-0179]
//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
//     greco  [03489940240 03489900120]
//
//     What is Greco's second phone number?
//
//
//  4. Shopping basket by Customer ID
//     -------------------------------
//     100 [617841573:4 576872813:2]
//     101 [576872813:5 657473833:20]
//     102 []
//
//     How many of 576872813 the customer 101 is going to buy?
//                (Product ID)  (Customer ID)
//
//
// EXPECTED OUTPUT
//
//   1. Run the solution to see the output
//   2. Here is the output with empty maps:
//
//      dulin's phone number: N/A
//      Product ID #879401371 is not available
//      greco's 2nd phone number: N/A
//      Customer #101 is going to buy 0 from Product ID #576872813.
//
// ---------------------------------------------------------
import "fmt"

func main() {
	phone := map[string]string{
		"Gaikwad": "333999000",
		"Ghodake": "222003333",
	}

	avail := map[int]bool{
		1: true,
		2: false,
		3: true,
		4: false,
		5: true,
	}

	contacts := map[string][]string{
		"Joshi":  {"123", "345"},
		"Gandhi": {"677", "234"},
		"Naik":   {"11"},
	}

	baskets := map[int]map[int]int{
		51: {1: 1, 2: 0, 3: 2},
		52: {1: 5, 2: 0, 3: 1},
	}

	fmt.Print("Gaikwad's phone number: ")
	if value, ok := phone["Gaikwad"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("N/A")
	}

	fmt.Print("Product ID #5 is ")
	if value, ok := avail[5]; ok && value {
		fmt.Println("available")
	} else {
		fmt.Println("not available")
	}

	fmt.Print("Gandhi's 2nd phone number: ")
	gandhiPhones, ok := contacts["Gandhi"]
	if ok && len(gandhiPhones) >= 2 {
		fmt.Println(gandhiPhones[1])
	} else {
		fmt.Println("N/A")
	}

	cID := 52
	pID := 1
	basket, ok := baskets[cID]
	if ok {
		if count, found := basket[pID]; found {
			fmt.Printf("Customer #%d is going to buy %d from Product ID #%d\n", cID, count, pID)
		}
	}
}
