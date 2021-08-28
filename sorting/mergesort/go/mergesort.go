package main

import (
	"fmt"
	"math/rand"
	"mergesort/helper"
	"time"

	"github.com/bxcodec/faker/v3"
)

type Person struct {
	FirstName string
	LastName  string
}
type StudentGrade struct {
	Student Person
	Grade   uint
}
type StudentGradeList []StudentGrade

func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func GenerateStudentGradeList(quantity int) (list StudentGradeList) {
	var fn, ln string
	var sgl = make(StudentGradeList, 0)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < quantity; i++ {
		fn = faker.FirstName()
		ln = faker.LastName()

		sgl = append(sgl, StudentGrade{
			Student: Person{fn, ln},
			Grade:   uint(rand.Intn(1000)),
		})
	}
	return sgl
}

func mergeSort(items StudentGradeList) StudentGradeList {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(left StudentGradeList, right StudentGradeList) StudentGradeList {
	var temp = make(StudentGradeList, 0)

	i := 0
	j := 0

	for (i < len(left)) && (j < len(right)) {
		if left[i].Grade < right[j].Grade {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}

	if i >= len(left) {
		temp = append(temp, right[j:]...)
	} else if j >= len(right) {
		temp = append(temp, left[i:]...)
	}

	return temp
}

func main() {
	list := GenerateStudentGradeList(15)
	temp := make([]uint, 0)
	for _, l := range list {
		temp = append(temp, l.Grade)
	}

	fmt.Println(helper.ArrayToString(temp, ","))

	list = mergeSort(list)
	temp = []uint{}
	for _, l := range list {
		temp = append(temp, l.Grade)
	}
	fmt.Println(helper.ArrayToString(temp, ","))
}
