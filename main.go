//comment

package main

import "fmt"

func main() {

	// var g int = 55
	y := 42
	z := 42.0

	//Declare a Variable to hold a value of a certain type
	//then assign a value of that type to the variable initializing a variable
	//initializing a variable
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	z = float64(m)
	//conversion
	fmt.Printf("%v of type %T \n", z, z)

	//zero value is usually used when var is declared
	//use shortcut to assign variables typically - short declaration operator

	//numeral systems -hexadecimal is actually super cool
	//jonah we got this dude
	//sjcndscjndic
}
