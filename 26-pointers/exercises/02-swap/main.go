// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	var num1, num2 float64
	num1 = 10
	num2 = 20
	fmt.Printf("Before func swapVal: num1 %f and num2 %f\n", num1, num2)
	swapVal(&num1, &num2)
	fmt.Printf("After func swapVal: num1 %f and num2 %f\n", num1, num2)

	var pNum1, pNum2 *float64
	pNum1 = &num1
	pNum2 = &num2
	fmt.Printf("Before func swapAddr: num1 %f and num2 %f\n", *pNum1, *pNum2)
	swapAddr(pNum1, pNum2)
	fmt.Printf("After func swapAddr: num1 %f and num2 %f\n", *pNum1, *pNum2)
}

func swapVal(n1, n2 *float64) {
	*n2, *n1 = *n1, *n2
	// tmp := *n1
	// *n1 = *n2
	// *n2 = tmp
}

func swapAddr(p1, p2 *float64) {
	tmp := *p1
	*p1 = *p2
	*p2 = tmp
}
