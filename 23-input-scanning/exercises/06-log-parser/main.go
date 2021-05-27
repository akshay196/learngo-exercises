// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------
import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

func main() {
	var (
		domains  map[string]int    // Visit per domain
		total    int               // Total number of visits
		lnum     int               // Number of lines parsed
	)
	domains = make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		lnum++
		line := in.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)", fields, lnum)
			return
		}
		visit, err := strconv.Atoi(fields[1])
		if err != nil || visit < 0{
			fmt.Printf("wrong input: %q (line #%d)", fields[1], lnum)
			return
		}
		domains[fields[0]] += visit
		total += visit
	}

	fmt.Printf("%-50s %20s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 72))
	for domain, visit := range domains {
		fmt.Printf("%-50s %20d\n", domain, visit)
	}
	fmt.Printf("\n%-50s %20d\n", "TOTAL", total)
}
