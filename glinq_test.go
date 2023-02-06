package glinq

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

type (
	MyStruct struct {
		Name  string
		Order int
	}
)

func TestSelect(t *testing.T) {
	var ms []*MyStruct
	for i := 0; i < 5; i++ {
		ms = append(ms, &MyStruct{Name: fmt.Sprintf("zs%v", i), Order: rand.Intn(50)})
	}
	var s []string
	Instance(ms).Select(func(m *MyStruct) any {
		return m.Name
	}).ToList(&s)
	fmt.Println(s)
}

func TestInstance(t *testing.T) {
	var ms []*MyStruct
	for i := 0; i < 5; i++ {
		ms = append(ms, &MyStruct{Name: fmt.Sprintf("zs%v", i), Order: rand.Intn(50)})
	}
	fmt.Println("===>><<")
	for _, m := range ms {
		fmt.Println(m)
	}
	fmt.Println("===>><<")
	fmt.Println("===>><<")
	// < 正序 > 倒叙
	list := Instance(ms).Sort(func(current, next *MyStruct) bool {
		return current.Order < next.Order
	}).ToList()
	for _, itm := range list {
		fmt.Println(itm)
	}
}

// 140.82.114.3
func TestIndex(t *testing.T) {
	var ms []*MyStruct
	ms = append(ms, &MyStruct{Name: fmt.Sprintf("zs%v", 1), Order: 5})
	ms = append(ms, &MyStruct{Name: fmt.Sprintf("zs%v", 1), Order: 5})
	ms = append(ms, &MyStruct{Name: fmt.Sprintf("zs%v", 1), Order: 6})
	ms = append(ms, &MyStruct{Name: fmt.Sprintf("zs%v", 1), Order: 1})
	ms = append(ms, &MyStruct{Name: fmt.Sprintf("zs%v", 1), Order: 3})
	ms = append(ms, &MyStruct{Name: fmt.Sprintf("zs%v", 1), Order: 3})
	Instance(ms).Distinct(func(item *MyStruct) any {
		return item.Order
	}).ForEach(func(index int, item *MyStruct) {
		fmt.Println(item)
	})
}
func TestF(t *testing.T) {
	fn := func(item int) string {
		return "1"
	}
	of := reflect.TypeOf(fn)
	fmt.Println(of)
}
