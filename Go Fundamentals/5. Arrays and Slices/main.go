package main

import "fmt"

func main() {
	// Arrays
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums[3])

	// Slices
	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums2[0])

	// Adding to a slice
	nums2 = append(nums2, 6)

	fmt.Println(nums2)

	// Providing a low and high-bound
	fmt.Println(nums2[0:4])

	// Exercises
	languages := []string{"Go", "JavaScript", "Ruby", "Python"}
	fmt.Println(languages)

	fmt.Println(len(languages))

	fmt.Println(languages[1:3])

	languages = append(languages, "PHP")
	fmt.Println(languages)
}
