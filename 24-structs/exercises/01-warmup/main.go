// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------
import (
	"fmt"
)

type game struct {
	id     int
	name   string
	price  int
}
type store struct {
	game
	genre  string
}

func main() {
	inanc_store := []store{
		{
			game: game{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{
			game: game{id: 2, name: "x-com 2", price: 40},
			genre: "strategy",
		},
		{
			game: game{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}

	fmt.Printf("Inanc's game store has %d games.\n\n", len(inanc_store))

	for _, game := range inanc_store {
		fmt.Printf("#%d: %-15q %-20s $%d\n", game.id, game.name, "("+game.genre+")", game.price)
	}
}
