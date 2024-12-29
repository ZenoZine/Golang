package main

import "fmt"

func fizzbuzz(n int) []string {
	words := []string{}
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			words = append(words, "FizzBuzz")
		} else if i%3 == 0 {
			words = append(words, "Fizz")
		} else if i%5 == 0 {
			words = append(words, "Buzz")
		} else {
			words = append(words, fmt.Sprintf("%d", i))

			// Solution uses this line:
			// words = append(words, strconv.Itoa(i))
			// I need to look into what the strconv import does
			// I assume that it converts values into strings

			// Almost ^ It turns strings into integers
		}
	}
	return words
}

func main() {
	// sum := 0

	// // Regular for-loop
	// for i := 0; i <= 20; i += 2 {
	// 	fmt.Println(i)
	// }

	// // While loop
	// i := 0

	// for i < 5 {
	// 	sum = sum + i
	// 	i++
	// }

	// fmt.Println(sum)

	fmt.Println(fizzbuzz(15))
}
