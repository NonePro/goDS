package array

type Array struct {
	data []int // 未指定长度是不是slice
	size int
}

// 构造函数，传入数组的容量capacity构造Array
func New(capacity int) *Array {
	return &Array{
		// 这种方式指定的长度，算是数组还是slice ?
		data: make([]int, capacity),
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


