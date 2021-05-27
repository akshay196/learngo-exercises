// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	for _, n := range args {
		// k := 5
		// w := 2

		i, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		var isPrime bool
		if i == 2 || i == 3 {
			isPrime = true
			goto done
		}
		if i == 1 || i % 2 == 0 || i % 3 == 0 {
			goto done
		}

		for k, w := 5, 2; k * k <= i; {
			if i % k == 0 {
				goto done
			}
			k += w
			w = 6 - w
		}
		isPrime = true
	done:
		if isPrime {
			fmt.Print(n, " ")
		}
	}

}
