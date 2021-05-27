// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------
import (
	"fmt"
	"os"
	"bufio"
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

var commands string = `
  > list   : lists all the games
  > quit   : quits
`

func main() {
	inanc_store := []store{
		{game: game{id: 1, name: "god of war", price: 50},genre: "action adventure"},
		{game: game{id: 2, name: "x-com 2", price: 40},genre: "strategy"},
		{game: game{id: 3, name: "minecraft", price: 20},genre: "sandbox"},
	}

	fmt.Printf("Inanc's game store has %d games.\n", len(inanc_store))
	fmt.Println(commands)
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {

		switch in.Text() {
		case "list":
			for _, g := range inanc_store {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		case "quit":
			fmt.Println("bye!")
			return
		}
		fmt.Println(commands)
	}
	
}
