//I want to create a simple elfster dupe for my families christmas swap

// each person can add books to the overall list
//brainstorm
//need to work on
//randomly assigns a person someone else for the gift exchange

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 1
	members := [7]string{"Jonah", "Maile", "Theo", "MB", "Mark", "Miah", "Abby"}
	for i < len(members) {
		fmt.Println(rand.Intn(len(members)))
		fmt.Println(members[i])
		i++
	}
}
