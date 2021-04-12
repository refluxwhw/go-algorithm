package sort

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func check(arr []int) bool {
	last := 0
	for i:=0; i< len(arr); i++ {
		fmt.Printf("%d ", arr[i])
		if i > 0 && arr[i]<last {
			return false
		}
		last = arr[i]
	}

	return true
}

func Test_Sort(t *testing.T) {
	max := 10000
	arr := make([]int, max)
	for i:=0; i< len(arr); i++ {
		arr[i] = rand.Int() % max
	}

	var s Sorter

	convey.Convey("shell", t, func(c convey.C) {
		s = NewSorter("shell")
		s.Sort(arr)
		c.So(check(arr), convey.ShouldBeTrue)
	})
}


