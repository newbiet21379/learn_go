// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// EXERCISE: Get the `max` from the user
//           And create the table according to that.


func main() {
	// print the header
	var (
		max int
		err error
	)
	if len(os.Args) != 2{
		fmt.Println("Give me a number")
		return
	}
	s := os.Args[1]
	max, err = strconv.Atoi(s)
	if err != nil {
		fmt.Printf("You input %q is not a number", s)
		return
	}

	fmt.Println("Multiply Table: ")
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}

	fmt.Println("Addition Table: ")
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i+j)
		}
		fmt.Println()
	}

	fmt.Println("Subtraction Table: ")
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i-j)
		}
		fmt.Println()
	}

	fmt.Println("Division Table: ")
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			if j == 0{
				fmt.Printf("%5s", "X")
				continue
			}
			fmt.Printf("%5.2f", float32(i)/float32(j))
		}
		fmt.Println()
	}

	s = "Lazy cat jumps again and again and again"
	var (
		i int
		v string
	)
	words := strings.Fields(s)
	for i = 0 ;i< len(words);i++{
		fmt.Printf("%-2d : %q\n", i+1 , words[i])
	}
	fmt.Println()
	for i, v = range words {
		fmt.Printf("%-2d : %q\n", i+1 , v)
	}
}
