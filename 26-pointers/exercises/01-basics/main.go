// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var cp *computer

	// compare the pointer variable to a nil value, and say it's nil
	if cp == nil {
		fmt.Println("cp is nil pointer")
	}

	// create an apple computer by putting its address to a pointer variable
	c := computer{brand: "apple"}
	cp = &c

	// put the apple into a new pointer variable
	cp1 := &c

	// compare the apples: if they are equal say so and print their addresses
	fmt.Printf("cp poninter address :  %p\n", cp)
	fmt.Printf("cp1 poninter address:  %p\n", cp1)
	if cp == cp1 {
		fmt.Println("cp and cp1 are equal.")
	}

	// create a sony computer by putting its address to a new pointer variable
	scp := &computer{brand: "sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if cp != scp {
		fmt.Printf("cp pointer address : %p\n", cp)
		fmt.Printf("scp pointer address: %p\n", scp)
		fmt.Println("cp and scp are unequal.")
	}

	// put apple's value into a new ordinary variable
	newVar := *cp

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("cp pointer address : %p\n", &cp)
	fmt.Printf("cp points to addr  : %p\n", cp)
	fmt.Printf("var apple address  : %p\n", &newVar)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *cp == newVar {
		fmt.Println("*cp is equals to newVar")

		// print the values:
		// the value that is pointed by the apple and the new variable
		fmt.Printf("cp : %+v\nnewVar : %+v\n", *cp, newVar)
	}


	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func
	change(scp, "hp")
	// print sony's brand
	fmt.Printf("scp is : %+v\n", *scp)

	// print the returned value
	fmt.Printf("returned value : %+v\n", getValue(cp1))

	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell pointer : %p\n", newComputer("dell"))
}


func change(c *computer, brand string) {
	c.brand = brand
}

// write a func that returns the value that is pointed by the given *computer
func getValue(c *computer) computer {
	return *c
}

// write a new constructor func that returns a pointer to a computer
func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
