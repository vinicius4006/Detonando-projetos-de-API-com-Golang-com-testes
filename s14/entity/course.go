package entity

import "fmt"

type CourseInterface interface {
	Register(student Student)
	PrintCourse() string
	TotalStudents(students ...Student) int
}

type Course struct {
	Name     string
	Students []Student
}

type Student struct {
	Name string
	Age  int
}

func (c *Course) Register(student Student) {
	c.Students = append(c.Students, student)
}

func (c Course) PrintCourse() string {
	studentsInfo := "Students:\n"
	for s := range c.Students {
		studentsInfo += fmt.Sprintf("%v %v\n", c.Students[s].Name, c.Students[s].Age)
	}
	return fmt.Sprintf("Course: %v\n%v", c.Name, studentsInfo)
}

func (c Course) TotalStudents(students ...Student) int {
	return len(students)
}
