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

// 向所有元素后添加一个新元素
func (a *Array) AddLast(e int) {
	a.Add(a.size, e)
}

func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}

// 向数组的指定位置添加元素
func (a *Array) Add(index int, e int) {
	if a.size == a.GetCapacity() {
		panic("Add failed. Array is full")
	}
	// 这里有疑问，为啥是小于等于size呢，元素存在空位会有什么问题吗
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
