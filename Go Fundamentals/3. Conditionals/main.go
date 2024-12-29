package main

import "fmt"

func main() {
	//	Conditionals

	number := 28000

	if number < 0 {
		fmt.Printf("%d is negative", number)
	} else if number > 0 && number < 100 {
		fmt.Printf("%d is positive", number)
	} else {
		fmt.Printf("%d is positve and is a large number!", number)
	}
}
