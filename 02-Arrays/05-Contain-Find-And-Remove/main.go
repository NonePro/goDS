package main

import (
	"fmt"
	"goDS/02-Arrays/05-Contain-Find-And-Remove/array"
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
}
