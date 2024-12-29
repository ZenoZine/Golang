package main

import "fmt"

// Part of the exercise
func reduce(nums []int) int {
	sum := 0
	for _, numbers := range nums {
		sum = sum + numbers
	}
	return sum
}

func main() {
	athletes := []string{"Stephen", "Klay", "Draymond", "Harrison", "Andrew"}

	for i, name := range athletes {
		fmt.Printf("i = %d, name = %s\n", i, name)
	}

	// Make note of the blank spot where i is in the first example
	for _, name := range athletes {
		fmt.Printf("Name = %s\n", name)
	}

	nums := []int{30, 11, 40, 23, 12}

	for _, numbers := range nums {
		if numbers%2 == 0 {
			fmt.Println(numbers)
		}
	}

	nums2 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	fmt.Print(reduce(nums2))
}
