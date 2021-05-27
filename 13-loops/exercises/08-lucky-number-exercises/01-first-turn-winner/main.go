// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

import (
	"os"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Pick a number.")
		return
	}
	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Wrong number")
		return
	}

	rand.Seed(time.Now().Unix())
	target := rand.Intn(9)
	fmt.Println(target)

	if num == target {
		fmt.Println("Correct guess.")
		return
	}
	fmt.Println("Try again next time.")
}
