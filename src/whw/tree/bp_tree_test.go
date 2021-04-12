package tree

import (
	"fmt"
	"testing"
)

func TestBTreePlus_Insert(tt *testing.T) {
	t := NewBpt(5)
	arr := []int{1, 5, 4, 8, 60, 40, 25, 18, 47, 86, 24, 11, 15, 36, 2, 44, 69, 11}
	arrMap := make(map[int]int)

	for i, v := range arr {
		t.Insert(v, v)
		fmt.Printf("insert: %2d == ", v)
		t.Print()

		arrMap[i] = 0
	}

	fmt.Println(t.Find(4))
	fmt.Println(t.Find(22))
	fmt.Println(t.Find(11))
	fmt.Println(t.Find(66))

	for i, _ := range arrMap {
		fmt.Printf("delete %2d == ", arr[i])
		key, ok := t.Delete(arr[i])
		fmt.Printf("%5v %2d ~~~  ", ok, key)
		t.Print()
	}
}
