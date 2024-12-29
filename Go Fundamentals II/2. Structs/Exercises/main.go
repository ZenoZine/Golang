package main

import "fmt"

type Student struct {
	id   uint16
	name string
}

type Classroom struct {
	id          uint16
	capacity    uint16
	subject     string
	studentList []Student
}

func main() {
	c1 := Classroom{
		id:       12,
		capacity: 26,
		subject:  "History",
		studentList: []Student{
			{
				id:   5,
				name: "Joshua",
			},
			{
				id:   30,
				name: "Sloan",
			},
		},
	}
	c2 := new(Classroom)
	c2.id = 2
	c2.capacity = 100
	c2.subject = "Theater"
	c2.studentList = []Student{
		{
			id:   40,
			name: "Vince",
		},
		{
			id:   50,
			name: "Joshua",
		},
	}

	fmt.Println(c1)
	fmt.Println(c2)
}
