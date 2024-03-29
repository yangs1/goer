package type_test

import "testing"

type MyInt int64

func TestImpliccit(t *testing.T) {
	var a int = 1
	var b int64

	b = int64(a)

	var c MyInt
	c = MyInt(b)

	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a

	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)
}

func TestString(t *testing.T) {
	var a string
	t.Log(len(a))
}