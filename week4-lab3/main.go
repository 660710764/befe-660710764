package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json":email"`
	Year  int     `json:year`
	GPA   float64 `json:gpa`
}

func (s *Student) IsHonor() bool {
	return s.GPA >= 3.75
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("name is rquired")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 1-4")
	}

	return nil
}

func main() {
	// var st Student = Student{ID: "1", Name: "Banlu", Email: "chimsing_b@silpakorn.edu", Year: 3, GPA: 4.00}

	students := []Student{
		{ID: "1", Name: "Banlu", Email: "chimsing_b@silpakorn.edu", Year: 3, GPA: 4.00},
		{ID: "2", Name: "Chimsing", Email: "banlu_c@silpakorn.edu", Year: 4, GPA: 4.00},
	}

	newStudent := Student{ID: "3", Name: "Mon", Email: "doraemon@silpakorn.edu", Year: 1, GPA: 4.00}

	students = append(students, newStudent)

	for i, students := range students {
		fmt.Printf("%d Honor %v\n", i, students.IsHonor())
		fmt.Printf("%d Validation %v\n", i, students.Validate())
	}

}
