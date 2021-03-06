// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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
	var i = min
	for {
		if i > max {
			break
		}
		if i%2 != 0 {
			i++
			continue
		}
		sum += i
		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}
		i++
	}
	fmt.Printf(" = %d\n", sum)
}














