package main

import (
	"fmt"
	"goDS/02-Arrays/02-Create-Our-Own-Array/array"
)

func main() {
	arr := array.New(10)
	fmt.Println(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())
}
