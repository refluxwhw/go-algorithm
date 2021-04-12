package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-algorithm/whw/sort"
)

func main() {
	max := 10000 * 10
	arr := make([]int, max)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(max * 10)
	}

	// arr := []int{4, 42, 94, 46, 88, 39, 33, 6, 60, 15}

	//test("shell", arr, 50)
	testAll(arr, 1)
}

func testAll(arr []int, n int) {
	names := sort.GetSorterNames()
	for _, name := range names {
		test(name, arr, n)
	}
}

func test(name string, arr []int, n int) {
	var s sort.Sorter
	s = sort.NewSorter(name)
	var tmpArr []int
	cost := int64(0)
	for i := 0; i < n; i++ {
		tmpArr = make([]int, len(arr))
		copy(tmpArr, arr)
		start := time.Now()
		tmpArr = s.Sort(tmpArr)
		cost += time.Now().Sub(start).Nanoseconds()
	}

	toStr := func(n int64) string {
		sec := n / 1e9
		ms := float64(n%1e9) / 1e6
		return fmt.Sprintf("%4ds %7.3fms", sec, ms)
	}

	ns := cost / int64(n)
	fmt.Printf("%10s: test %d times, arvage cost: %s, check result: %v\n",
		name, n, toStr(ns), check(tmpArr))
}

func check(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Printf("index: arr[%d]=%d, arr[%d+1]=%d", i, arr[i], i, arr[i+1])
			return false
		}
	}

	return true
}
