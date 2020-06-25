package array

import (
	"reflect"
	"testing"
)

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}

func TestArray_Add(t *testing.T) {
	type args struct {
		index int
		e     interface{}
	}
	type want struct {
		data     []interface{}
		size     int
		capacity int
	}
	type student struct {
		name  string
		score int
	}
	tests := []struct {
		name string
		data *Array
		args args
		want want
	}{
		{
			name: "int type add normal",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{index: 1, e: 5},
			want: want{data: []interface{}{0, 5, 1, 2, 3}, size: 5, capacity: 10},
		},
		{
			name: "int type add with expand",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 4),
			args: args{index: 1, e: 5},
			want: want{data: []interface{}{0, 5, 1, 2, 3}, size: 5, capacity: 8},
		},
		{
			name: "string type add normal",
			data: NewWithInitElementsAndCapacity([]interface{}{"Apple", "Orange", "Banana", "Strawberry"}, 10),
			args: args{index: 1, e: "Mango"},
			want: want{data: []interface{}{"Apple", "Mango", "Orange", "Banana", "Strawberry"}, size: 5, capacity: 10},
		},
		{
			name: "struct type add normal",
			data: NewWithInitElementsAndCapacity([]interface{}{student{name: "NonePro", score: 80}, student{name: "Bob", score: 88}, student{name: "Jack", score: 77}, student{name: "Tom", score: 98}}, 10),
			args: args{index: 1, e: student{name: "LiLei", score: 99}},
			want: want{data: []interface{}{student{name: "NonePro", score: 80}, student{name: "LiLei", score: 99}, student{name: "Bob", score: 88}, student{name: "Jack", score: 77}, student{name: "Tom", score: 98}}, size: 5, capacity: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			a.Add(tt.args.index, tt.args.e)
			for i, v := range tt.want.data {
				if nv := a.Get(i); v != nv {
					t.Errorf("Except value of %v at index %d is %v, but get %v", a, i, v, nv)
				}
			}
			if e := a.Get(tt.args.index); e != tt.args.e {
				t.Errorf("Expect %v add %d at index %d then get at the same index should got %v, but got %v", a, tt.args.e, tt.args.index, tt.args.e, e)
			}
			if a.size != tt.want.size {
				t.Errorf("Expect size of %v is %d , but got %d", a, tt.want.size, a.size)
			}
			if a.GetCapacity() != tt.want.capacity {
				t.Errorf("Expect capacity of %v is %d , but got %d", a, tt.want.capacity, a.GetCapacity())
			}
		})
	}
}

func TestArray_AddWithPanic(t *testing.T) {
	type args struct {
		index int
		e     interface{}
	}
	type want struct {
		data []interface{}
		size int
	}
	tests := []struct {
		name string
		data *Array
		args args
		want want
	}{
		{
			name: "int type index < 0",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{index: -1, e: 5},
			want: want{data: []interface{}{5, 0, 1, 2, 3}, size: 5},
		},
		{
			name: "int type index > size-1",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{index: 5, e: 5},
			want: want{data: []interface{}{5, 0, 1, 2, 3}, size: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			shouldPanic(t, func() {
				a.Add(tt.args.index, tt.args.e)
			})
		})
	}
}

func TestArray_AddFirst(t *testing.T) {
	type args struct {
		e interface{}
	}
	type want struct {
		data []interface{}
		size int
	}
	tests := []struct {
		name string
		data *Array
		args args
		want want
	}{
		{
			name: "int type add normal",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{e: 5},
			want: want{data: []interface{}{5, 0, 1, 2, 3}, size: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			a.AddFirst(tt.args.e)
			for i, v := range tt.want.data {
				if nv := a.Get(i); v != nv {
					t.Errorf("Except value of %d is %v, but get %v", i, v, nv)
				}
			}
			if e := a.Get(0); e != tt.args.e {
				t.Errorf("Expect %v add %d at index %d then get %v at index %d, but get %d", a, tt.args.e, 0, tt.args.e, 0, e)
			}
			if a.size != tt.want.size {
				t.Errorf("Expect the size of %v is %d after add first , but got %d", a, tt.want.size, a.size)
			}
		})
	}
}

func TestArray_AddLast(t *testing.T) {
	type args struct {
		e interface{}
	}
	type responses struct {
		data []interface{}
		size int
	}
	tests := []struct {
		name      string
		data      *Array
		args      args
		responses responses
	}{
		{
			name:      "int type add normal",
			data:      NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args:      args{e: 5},
			responses: responses{data: []interface{}{0, 1, 2, 3, 5}, size: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			a.AddLast(tt.args.e)
			for i, v := range tt.responses.data {
				if nv := a.Get(i); v != nv {
					t.Errorf("Except value of %d is %v, but get %v", i, v, nv)
				}
			}
			if e := a.Get(a.size - 1); e != tt.args.e {
				t.Errorf("Expect %v add %d at last then at the same index got %v, but get %v", a, tt.args.e, tt.args.e, e)
			}
			if a.size != tt.responses.size {
				t.Errorf("Expect size of %v is %d , but get %d", a, tt.responses.size, a.size)
			}
		})
	}
}

func TestArray_Contains(t *testing.T) {
	type args struct {
		e interface{}
	}
	tests := []struct {
		name string
		data *Array
		args args
		want bool
	}{
		{
			name: "int type contains",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{e: 3},
			want: true,
		},
		{
			name:
			"int type not contains",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{e: 8},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.Contains(tt.args.e); got != tt.want {
				t.Errorf("%v Contains(%v) = %v, want %v", a, tt.args.e, got, tt.want)
			}
		})
	}
}

func TestArray_Find(t *testing.T) {
	type args struct {
		e interface{}
	}
	tests := []struct {
		name string
		data *Array
		args args
		want int
	}{
		{
			name: "int type find first",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{e: 0},
			want: 0,
		},
		{
			name: "int type find last",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{e: 3},
			want: 3,
		},
		{
			name: "int type can not find",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{e: 5},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.Find(tt.args.e); got != tt.want {
				t.Errorf("%v Find(%v) = %v, want %v", a, tt.args.e, got, tt.want)
			}
		})
	}
}

func TestArray_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		data *Array
		args args
		want interface{}
	}{
		{
			name: "int type get first",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{0},
			want: 0,
		},
		{
			name: "int type get last",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{3},
			want: 3,
		},
		{
			name: "int type get normal",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_GetWithPanic(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		data *Array
		args args
		want interface{}
	}{
		{
			name: "int type get first",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{-1},
			want: 0,
		},
		{
			name: "int type get last",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{4},
			want: 3,
		},
		{
			name: "int type get normal",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			args: args{5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			shouldPanic(t, func() {
				a.Get(tt.args.index)
			})
		})
	}
}

func TestArray_GetCapacity(t *testing.T) {
	tests := []struct {
		name string
		data *Array
		want int
	}{
		{
			name: "int type get first",
			data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10),
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.GetCapacity(); got != tt.want {
				t.Errorf("GetCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_GetSize(t *testing.T) {
	tests := []struct {
		name string
		data *Array
		want int
	}{
		{name: "int type array get size", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.GetSize(); got != tt.want {
				t.Errorf("%v GetSize() = %v, want %v", a, got, tt.want)
			}
		})
	}
}

func TestArray_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		data *Array
		want bool
	}{
		{name: "int type is empty", data: NewWithInitElementsAndCapacity([]interface{}{}, 10), want: true},
		{name: "int type is not empty", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.IsEmpty(); got != tt.want {
				t.Errorf(" %v isEmpty() = %v, want %v", a, got, tt.want)
			}
		})
	}
}

func TestArray_Remove(t *testing.T) {
	type args struct {
		index int
	}
	type want struct {
		e        interface{}
		capacity int
	}
	tests := []struct {
		name string
		data *Array
		args args
		want want
	}{
		{name: "int type remove normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{index: 0}, want: want{0, 10}},
		{name: "int type remove normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{index: 3}, want: want{3, 10}},
		{name: "int type remove with shrink", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3, 4, 5}, 10), args: args{index: 3}, want: want{3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.Remove(tt.args.index); !reflect.DeepEqual(got, tt.want.e) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestArray_RemoveWithPanic(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		data *Array
		args args
	}{
		{name: "int type remove normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{index: -1}},
		{name: "int type remove normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{index: 4}},
		{name: "int type remove normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{index: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			shouldPanic(t, func() {
				a.Remove(tt.args.index)
			})
		})
	}
}

func TestArray_RemoveElement(t *testing.T) {
	type args struct {
		e interface{}
	}
	tests := []struct {
		name string
		data *Array
		args args
	}{
		{name: "int type exist", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{2}},
		{name: "int type not exist", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			i := a.Find(tt.args.e)
			size := a.GetSize()
			a.RemoveElement(tt.args.e)
			if i != -1 {
				if a.GetSize() != size-1 {
					t.Errorf("Remove exist element %v excpet size become %d, but got %d", tt.args.e, size-1, a.GetSize())
				}
			} else {
				if a.GetSize() != size {
					t.Errorf("Remove not exist element excpet size not change with %d, but got %d", size-1, a.GetSize())
				}
			}

		})
	}
}

func
TestArray_RemoveFirst(t *testing.T) {
	tests := []struct {
		name string
		data *Array
		want interface{}
	}{
		{name: "int type normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), want: 0},
		{name: "int type normal", data: NewWithInitElementsAndCapacity([]interface{}{9, 1, 2, 3}, 10), want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			size := a.GetSize()
			if got := a.RemoveFirst(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirst() = %v, want %v", got, tt.want)
			}
			if a.GetSize() != size-1 {
				t.Errorf("After RemoveFirst element of %v require size %d, but got %d", a, size-1, a.GetSize())
			}
		})
	}
}

func
TestArray_RemoveLast(t *testing.T) {
	tests := []struct {
		name string
		data *Array
		want interface{}
	}{
		{name: "int type normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.RemoveLast(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArray_Set(t *testing.T) {
	type args struct {
		index int
		e     interface{}
	}
	type want struct {
		data []interface{}
	}
	tests := []struct {
		name string
		data *Array
		args args
		want want
	}{
		{name: "int type normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{1, 3}, want: want{data: []interface{}{0, 3, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			a.Set(tt.args.index, tt.args.e)
			for i, v := range tt.want.data {
				if v != a.Get(i) {
					t.Errorf("After set value %v with %v at index %d expect value %v at index %d, but got value %v", a, tt.args.e, tt.args.index, v, i, a.Get(i))
				}
			}
		})
	}
}

func TestArray_SetWithPanic(t *testing.T) {
	type args struct {
		index int
		e     interface{}
	}
	tests := []struct {
		name string
		data *Array
		args args
	}{
		{name: "int type normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{-1, 3}},
		{name: "int type normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{4, 3}},
		{name: "int type normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			shouldPanic(t, func() {
				a.Set(tt.args.index, tt.args.e)
			})
		})
	}
}

func
TestArray_String(t *testing.T) {
	tests := []struct {
		name string
		data *Array
		want string
	}{
		{name: "int type normal", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), want: "Array:{ size = 4, capacity = 10, data = [0, 1, 2, 3]}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			if got := a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func
TestArray_resize(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		data *Array
		args args
	}{
		{name: "normal not change expand", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{10}},
		{name: "normal expand", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{20}},
		{name: "normal shrink", data: NewWithInitElementsAndCapacity([]interface{}{0, 1, 2, 3}, 10), args: args{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.data
			a.resize(tt.args.capacity)
			if a.GetCapacity() != tt.args.capacity {
				t.Errorf("resize %v with capacity %d expect %d, but got %d", a, tt.args.capacity, tt.args.capacity, a.GetCapacity())
			}
		})
	}
}

func
TestNew(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *Array
	}{
		{name: "normal", args: args{10}, want: NewWithInitElementsAndCapacity([]interface{}{}, 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
