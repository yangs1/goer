package ch4

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 3, 4, 6} // 数组，可以比较
	c := []int{1, 2, 3, 4}    // 切片， 不能比较
	// 数组比较需要相等
	//t.Log(a == b)
	t.Log(a, b, c)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

// &^ 按位清0
func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestStringChange(t *testing.T) {
	//var a string = "abcd"

	//a[2] = "e"

	//t.Log(a)
}
