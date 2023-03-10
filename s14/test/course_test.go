package test

import (
	"testing"

	"github.com/vinicius4006/detonando-projeto-api-golang/entity"
)

func TestCourse(t *testing.T) {
	engineering := entity.Course{Name: "Computer Science", Students: make([]entity.Student, 0)}
	verificarRegister := func(courseM entity.CourseInterface, students ...entity.Student) {
		for _, s := range students {
			courseM.Register(s)
		}
		result := len(engineering.Students)
		lengthStudents := courseM.TotalStudents(students...)
		if result != lengthStudents {
			t.Errorf("result %v expected %v", result, lengthStudents)
		}

	}
	t.Run("Register", func(t *testing.T) {

		verificarRegister(&engineering, entity.Student{Name: "Vinicius", Age: 24}, entity.Student{Name: "Elvis", Age: 20})
	})

	t.Run("Name of course and all students", func(t *testing.T) {
		result := engineering.PrintCourse()
		expectedPrint := "Course: Computer Science\nStudents:\nVinicius 24\nElvis 20\n"
		if result != expectedPrint {
			t.Errorf("result:\n%v\nexpected:\n%v", result, expectedPrint)
		}
	})

}
