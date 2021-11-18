// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// package main is a special package
// it allows Go to create an executable file
package main

/*
This is a multi-line comment.

import keyword makes another package available
  for this .go "file".

import "fmt" lets you access fmt package's functionality
  here in this file.
*/
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		n int
		newN int
		err error
		newErr error
	)

	if len(os.Args) < 3 {
		fmt.Println("Input something")
		return
	}
	s := os.Args[1]
	newS := os.Args[2]
	newN, newErr = strconv.Atoi(os.Args[2])
	n, err = strconv.Atoi(s)
	if err != nil {
		fmt.Printf("%q is not a number", s)
		return
	}
	if newErr != nil{
		fmt.Printf("%q is not a number", newS)
	}

	if n > 0 {
		fmt.Printf("%d feet is %.2f meters", n , float32(n) / 3.28)
	}else {
		fmt.Printf("%d is lower than 0\n", n)
	}

	switch n+=10;{
	case newN > 100:
		fmt.Print("Big ")
		fallthrough
	case newN > 0:
		fmt.Print("Positive ")
		fallthrough
	default:
		fmt.Print("Number\n")
	}

	switch now := time.Now().Hour();{
	case now >= 18:
		fmt.Println("Good Evening")
	case now >= 12:
		fmt.Println("Good Afternoon")
	case now >= 6:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Night")
	}
}
