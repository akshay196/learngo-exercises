// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. i don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
outer:
	for {
		fmt.Printf(`
  > list   : lists all the games
  > quit   : quits
  > id N   : queries a game by id

`)

		if !in.Scan() {
			break
		}

		fmt.Println()
		cmd := in.Text()
		cmds := strings.Fields(cmd)
		num := ""
		if len(cmds) == 2 {
			cmd = cmds[0]
			num = cmds[1]
		}

		switch cmd {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}
		case "id":
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("wrong id")
				continue
			}
			for _, g := range games {
				if g.id == n {
					fmt.Printf("#%d: %-15q %-20s $%d\n",
						g.id, g.name, "("+g.genre+")", g.price)
					goto outer
				}
			}
			fmt.Println("sorry. i don't have the game")
		}
	}
}
