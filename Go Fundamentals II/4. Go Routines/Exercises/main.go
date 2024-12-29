package main

import (
	"fmt"
	"time"
)

func sliceString(strings []string) {
	for _, words := range strings {
		fmt.Println(words)
	}
}

func main() {
	start := time.Now()

	names := []string{"Joshua", "Tyler", "Sean", "Easton"}
	classes := []string{"Numerical Analysis", "Database Design", "Senior Seminar"}

	go sliceString(names)
	go sliceString(classes)

	duration := time.Since(start)

	fmt.Println(duration.String())

	time.Sleep(time.Second)
}
