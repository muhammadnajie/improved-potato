package main

import (
	"fmt"
	"insertion/helper"
	"math/rand"
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

func insertionSort(items StudentGradeList) StudentGradeList {
	count := len(items)
	for i := 1; i < count; i++ {
		sorted := i - 1
		pointed := i
		for j := sorted; j >= 0; j-- {
			if items[j].Grade > items[pointed].Grade {
				items[j], items[pointed] = items[pointed], items[j]
				pointed--
			} else {
				break
			}
		}
	}
	return items
}

func main() {
	list := GenerateStudentGradeList(15)
	temp := make([]uint, 0)
	for _, l := range list {
		temp = append(temp, l.Grade)
	}

	fmt.Println(helper.ArrayToString(temp, ","))

	list = insertionSort(list)
	temp = []uint{}
	for _, l := range list {
		temp = append(temp, l.Grade)
	}
	fmt.Println(helper.ArrayToString(temp, ","))
}
