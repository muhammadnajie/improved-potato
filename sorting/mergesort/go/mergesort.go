package main

import (
	"fmt"
	"math/rand"
	"strings"
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
	// var temp = make(StudentGradeList, len(left)+len(right))
	temp := StudentGradeList{}

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

	for ; i < len(left); i++ {
		temp = append(temp, left[i])
	}
	for ; j < len(right); j++ {
		temp = append(temp, right[j])
	}

	// if i > len(left) {
	// 	temp = append(temp, right[j:]...)
	// } else {
	// 	temp = append(temp, left[i:]...)
	// }

	return temp
}

func arrayToString(a []uint, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func main() {
	list := GenerateStudentGradeList(5)
	temp := make([]uint, 0)
	for _, l := range list {
		temp = append(temp, l.Grade)
	}

	fmt.Println(arrayToString(temp, ","))

	list = mergeSort(list)
	temp = []uint{}
	for _, l := range list {
		temp = append(temp, l.Grade)
	}
	fmt.Println(arrayToString(temp, ","))
}
