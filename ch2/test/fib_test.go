package test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		fmt.Println("   ", b)
		tmp := a
		a = b
		b = tmp + a
	}

	fmt.Println()
}

func TestExchange(t *testing.T) {
	a, b := 1, 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		a, b = b, b+a

	}

	t.Log(b)
}

func TestMakeCap(t *testing.T) {
	scores := make([]int, 0, 10)
	scoresA := scores[0:8]
	scoresA[7] = 9033
	t.Log(scores)
	t.Log(scoresA)
}

func TestCopyArr(t *testing.T) {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:3])
	fmt.Println(worst, len(worst), cap(worst))
}
