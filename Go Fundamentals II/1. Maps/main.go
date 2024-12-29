package main

import (
	"fmt"
	"strings"
)

func main() {
	dictionary := map[string]string{
		"Go":     "A programming language created by Google\n",
		"Gopher": "A software engineer who builds with Go\n",
		"Golang": "Another name for Go\n",
	}

	fmt.Println(dictionary)
	fmt.Println(dictionary["Gopher"])

	// Exercises
	courses := map[uint16]string{

		1: "Calculus",
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		6: "English",
		7: "Cantonese",
	}
	for id, course := range courses {
		if strings.HasPrefix(course, "C") {
			fmt.Println(id, course)
		}
	}

	courses[4] = "Algorithms"
	courses[8] = "Spanish"
	delete(courses, 1)

	fmt.Println("Updated courses: ", courses)

}
