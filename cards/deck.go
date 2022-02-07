// Details of deck
package main

import "fmt"

//create a new type of deck which is a slice of strings
type deck []string

func (d deck) print() {
	for i, eachcard := range d {
		fmt.Println(i, eachcard)
	}
}
