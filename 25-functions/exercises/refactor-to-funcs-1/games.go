package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func newGame(id, price int, name, genre string) game {
	return game{item: item{id: id, name: name, price: price}, genre: genre}
}

func addGame(games []game, g game) []game{
	return append(games, g)
}

func load() []game{
	var games []game
	g1 := newGame(1, 50, "god of war", "action adventure")
	g2 := newGame(2, 40, "x-com 2", "strategy")
	g3 := newGame(3, 20, "minecraft", "sandbox")
	games = addGame(games, g1)
	games = addGame(games, g2)
	games = addGame(games, g3)
	return games
}

func indexByID(games []game) map[int]game{
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}
	return byID
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
