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
	members := [7]string{"Jonah - 0 ", "Maile - 1 ", "Theo - 2 ", "MB - 3 ", "Mark - 4 ", "Miah - 5 ", "Abby - 6 "}
	i := 0
	for i < len(members) {
		fmt.Println(rand.Intn(len(members)))
		fmt.Print(members[i])
		i++
	}
}

//need to think of business ideas
//need to trust in the Lord
