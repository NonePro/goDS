package main

import "fmt"

func main() {
	// 声明指定长度数组,注意这里只是声明 arr0:=[5]int 是错误的，因为没弄清楚变量声明和赋值的区别
	var arr0 [5]int
	for i := 0; i < len(arr0); i++ {
		fmt.Println(arr0[i])
	}

	// 声明并赋值，注意这里类型在等号右边，实际的值通过打括号赋予
	scores := [...]int{100, 96, 80}
	for _, v := range scores {
		fmt.Println(v)
	}

	// 通过索引访问并修改值
	scores[0] = 98
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}
}
