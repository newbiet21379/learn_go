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
	"strings"
	"unicode/utf8"
)

func main() {
	// names are case-sensitive:
	// MyAge, myAge, and MYAGE are different variables

	// USE-CASE:
	// When to use a parallel declaration?
	//
	// NOT GOOD:
	// var myAge int
	// var yourAge int
	//
	// SO-SO:
	// var (
	// 	myAge int
	// 	yourAge int
	// )
	//
	// BETTER:
	name := "Lô lê la"
	var input []string
	input = os.Args
	l := len(input[3])
	repeat := strings.Repeat("!",l)
	var myAge, yourAge int
	fmt.Println(myAge, yourAge,repeat, strings.ToUpper(input[3]),repeat ,utf8.RuneCountInString(name))
	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"
	json = `{
        "Items": [{
                "Item": {
                        "name": "Teddy Bear"
                }
        }]
}
`
	msg := `
	
	The weather looks good.
I should go and play.
	`
	fmt.Println(json)

	const (
		Winter = 12 - 3 * iota
		Spring
		Summer
		Fall
	)
	fmt.Println("message", len(strings.TrimSpace(msg)), Winter, Spring , Summer , Fall)
	name = "inanç           "
	fmt.Println(len(strings.TrimRight(name," ")))
}
