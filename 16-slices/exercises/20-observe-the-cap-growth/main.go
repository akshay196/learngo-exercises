// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Observe the capacity growth
//
//  Write a program that appends elements to a slice
//  10 million times in a loop. Observe how the capacity of
//  the slice changes.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 10e6 times
//
//  3. On each iteration: Append an element to the slice
//
//  4. Print the length and capacity of the slice "only"
//     when its capacity changes.
//
//  BONUS: Print also the growth rate of the capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------
import (
	"fmt"
	"time"
)

func main() {
	var slice []int
	var preCap int
	for i:=0; i<=10e6; i++ {
		var growth int
		if preCap == 0 {
			growth = 0
		} else {
			growth = cap(slice)/preCap
		}
		fmt.Printf("len:%d\tcap:%d\tgrowth:%d\n", len(slice), cap(slice), growth)
		preCap = cap(slice)
		slice = append(slice, i)
		time.Sleep(time.Second)
	}
}
