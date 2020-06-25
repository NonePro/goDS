package array

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{} // 未指定长度是不是slice
	size int
}

// 构造函数，传入数组的容量capacity构造Array
func New(capacity int) *Array {
	return &Array{
		// 这种方式指定的长度，算是数组还是slice ?
		data: make([]interface{}, capacity),
	}
}

// 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获取数组的size
func (a *Array) GetSize() int {
	return a.size
}

// 判断数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 向所有元素后添加一个新元素
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

// 向数组的指定位置添加元素
func (a *Array) Add(index int, e interface{}) {
	if a.size == a.GetCapacity() {
		panic("Add failed. Array is full")
	}
	// 这里有疑问，为啥是小于等于size呢，元素存在空位会有什么问题吗
	// 现在到理解是这种实现考虑到是数组索引不具备语意性，需要大量到遍历操作。而具备语意性的场景直接可以访问操作。不在研究范围内
	if index < 0 || index > a.size {
		panic("Add failed. Required index >=0 and <= size")
	}
	// 移动数据
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = e
	a.size++
}

// 获取指定位置的元素值。 保护未被赋值的索引位置不会被访问到
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed. Out of index range.")
	}
	return a.data[index]
}

// 修改index索引位置的元素
func (a *Array) Set(index int, e int) {
	if index < 0 || index >= a.size {
		panic("Set failed. Out of index range.")
	}
	a.data[index] = e
}

// 查找数组中是否有元素 e
func (a *Array) Contains(e interface{}) bool {
	/* 这种实现会循环到 size 到 capacity 区间到元素，没必要
	for _, v := range a.data {
		if v == e {
			return true
		}
	}
	*/
	for i := 0; i < a.size; i++ {
		if a.Get(i) == e {
			return true
		}
	}
	return false
}

// 查看数组中是否包含指定元素
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Remove failed. out of the range")
	}
	v := a.data[index]
	for i := index; i < a.size; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	a.data[a.size] = nil //loitering object != memory leak

	return v
}

// 删除第一个元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// 删除最后一个元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 从数组中删除元素e
func (a *Array) RemoveElement(e interface{}) {
	i := a.Find(e)
	if i != -1 {
		a.Remove(i)
	}
}

// 重写Array 的 string 方法
func (a *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d \n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != a.size-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
