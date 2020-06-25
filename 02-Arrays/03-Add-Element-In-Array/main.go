package main

import (
	"fmt"
	"goDS/02-Arrays/03-Add-Element-In-Array/array"
)

func main() {
	arr := array.New(10)
	fmt.Println(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())

	arr.AddLast(6)
	fmt.Print(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())

	arr.Add(0, 9)
	fmt.Print(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())

	arr.Add(1, 10)
	fmt.Print(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())

	arr.Add(1, 6)
	fmt.Print(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())
}
