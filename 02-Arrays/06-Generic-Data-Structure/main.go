package main

import (
	"fmt"
	"goDS/02-Arrays/06-Generic-Data-Structure/array"
	"goDS/02-Arrays/06-Generic-Data-Structure/student"
)

func main() {
	arr := array.New(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)
	fmt.Println(arr.Get(5))
	fmt.Println(arr.Remove(2), arr)
	arr.RemoveElement(6)
	fmt.Println(arr)

	students := array.New(10)
	student1 := student.New("NonePro", 60)
	students.AddLast(student1)
	student2 := student.New("Alice", 70)
	students.AddLast(student2)
	students.AddLast(student.New("Jack", 80))
	fmt.Println(students)

	i := students.Find(student2)
	fmt.Println(i)
}
