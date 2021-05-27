// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------
import (
	"fmt"
	"strconv"
	"os"
)

const usage = `Usage:
       go run main.go <min> <max>
`

func main() {
	var min, max int
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	min, err := strconv.Atoi(args[1])
	max, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Enter numberic min and max values")
		return
	}
	if min > max {
		fmt.Println("Min is greater than max")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}
		
	}
	fmt.Printf(" = %d\n", sum)
}
