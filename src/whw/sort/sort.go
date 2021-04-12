package sort

import (
	"reflect"
)

type Sorter interface {
	Sort([]int) []int
}

var gSorterMap = make(map[string]func() Sorter)
var gSorterNames = make([]string, 0)

func init() {
	registerSorter(&bubble{})
	registerSorter(&insertion{})
	registerSorter(&insertion1{})
	registerSorter(&merge{})
	registerSorter(&quick{})
	registerSorter(&selection{})
	registerSorter(&shell{})
	registerSorter(&shell1{})
	registerSorter(&heap{})
	registerSorter(&bucket{})
	registerSorter(&bucket1{})
	registerSorter(&radix{})
}

func registerSorter(s Sorter) {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	name := t.Name()
	if _, ok := gSorterMap[name]; !ok {
		gSorterNames = append(gSorterNames, name)
	}
	gSorterMap[name] = func() Sorter {
		v := reflect.New(t)
		return v.Interface().(Sorter)
	}
}

func NewSorter(name string) Sorter {
	f, ok := gSorterMap[name]
	if !ok {
		return nil
	}
	return f()
}

func GetSorterNames() []string {
	return gSorterNames
}

