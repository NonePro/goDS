package main

import (
	"fmt"
	"goDS/02-Arrays/04-Query-And-Update-Element/array"
)

func main() {
	arr := array.New(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)
	fmt.Println(arr.Get(5))
}
