package ch5

import "testing"

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6}

	arr3_sec := arr3[:3] //取前 3 个元素
	arr3_thr := arr3[:len(arr3)]

	t.Log(arr3_thr, arr3_sec)
}

func TestSliceInit(t *testing.T) {
	var s0 []int

	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)

	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}

	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) // 只初始化 前 3 个

	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[3])  // error
	t.Log(s2[0], s2[1])
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[3])

}

func TestShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Unknow" // 共享内存， 会影响 Q2
	t.Log(Q2)
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}

	t.Log(m1[1])

	m1[2] = 0
	t.Log(m1[2])

	if v, ok := m1[3]; ok {
		t.Logf("key 3`s value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 4, 3: 9}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
