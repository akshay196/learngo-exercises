// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

import (
	"os"
	"fmt"
	"time"
	"strings"
)
func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	var leap bool
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		leap = true
	}

	if strings.ToLower(args[1]) == "january" {
		fmt.Printf("%q has 31 days.", args[1])
	} else if strings.ToLower(args[1]) == "february" {
		if leap {
			fmt.Printf("%q has 29 days.", args[1])
		} else {
			fmt.Printf("%q has 28 days.", args[1])
		}
	} else if strings.ToLower(args[1]) == "march" {
		fmt.Printf("%q has 31 days.", args[1])
	} else if strings.ToLower(args[1]) == "april" {
		fmt.Printf("%q has 30 days.", args[1])
	} else if strings.ToLower(args[1]) == "may" {
		fmt.Printf("%q has 31 days.", args[1])
	} else if strings.ToLower(args[1]) == "june" {
		fmt.Printf("%q has 30 days.", args[1])
	} else if strings.ToLower(args[1]) == "july" {
		fmt.Printf("%q has 31 days.", args[1])
	} else if strings.ToLower(args[1]) == "august" {
		fmt.Printf("%q has 31 days.", args[1])
	} else if strings.ToLower(args[1]) == "september" {
		fmt.Printf("%q has 30 days.", args[1])
	} else if strings.ToLower(args[1]) == "october" {
		fmt.Printf("%q has 31 days.", args[1])
	} else if strings.ToLower(args[1]) == "november" {
		fmt.Printf("%q has 30 days.", args[1])
	} else if strings.ToLower(args[1]) == "december" {
		fmt.Printf("%q has 31 days.", args[1])
	} else {
		fmt.Printf("%q is not a month.", args[1])
	}
}