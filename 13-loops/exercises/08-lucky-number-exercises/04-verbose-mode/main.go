// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your numbers are, harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	var verbose bool
	if args[0] == "-v" {
		verbose = true
		args = args[1:]
	}

	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess1, err := strconv.Atoi(args[0])
	guess2, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		var n int
		if guess1 > guess2 {
			n = rand.Intn(guess1 + 1)
		} else {
			n = rand.Intn(guess2 + 1)
		}
		if verbose {
			fmt.Println(n)
		}

		if n != guess1 && n != guess2 {
			continue
		}

		fmt.Println("🎉  YOU WON!")
		return
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
