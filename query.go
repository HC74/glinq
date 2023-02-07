package glinq

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	items 	[]any
	len 	int
}

func (s *Query) Get(index int) any {
	return s.items[index]
}

// GetI 返回int类型
func (s *Query) GetI(index int) int {
	return s.Get(index).(int)
}

// GetS 返回字符串类型
func (s *Query) GetS(index int) string {
	return s.Get(index).(string)
}

// GetF32 返回Float32类型
func (s *Query) GetF32(index int) float32 {
	return s.Get(index).(float32)
}

// GetF64 返回Float64类型
func (s *Query) GetF64(index int) float64 {
	return s.Get(index).(float64)
}

// GetI64 返回Int64类型
func (s *Query) GetI64(index int) int64 {
	return s.Get(index).(int64)
}

// GetB 返回布尔值
func (s *Query) GetB(index int) bool {
	return s.Get(index).(bool)
}

// ExcludeEmpty 排除空元素
func (s *Query) ExcludeEmpty() *Query {
	var results []any
	var item any
	for i := range s.items {
		item = s.Get(i)
		if item == nil {
			continue
		}
		sitem, ok := item.(string)
		if !ok {
			results = append(results, item)
			continue
		}
		if strings.TrimSpace(sitem) == "" {
			continue
		}
		results = append(results, item)
	}
	return s
}

// First 查找第一个元素
func (g *GLinq[T]) First() (T, error) {
	var item T
	slice := g.ToList()
	if len(slice) > 0 {
		return slice[0], nil
	}
	return item, ErrorCannotFound
}

// Skip 跳过指定数量的元素，如果num比元素数量大，则跳过全部元素
func (g *GLinq[T]) Skip(num int) *GLinq[T] {
	if g.len == 0 || num == 0 {
		return g
	}
	if num > g.len {
		num = g.len
	}
	g.elems = g.elems[num:]
	g.len = len(g.elems)
	return g
}

// Take 拿取指定数量的元素，如果num比元素数量大，则取出所有元素
func (g *GLinq[T]) Take(num int) *GLinq[T] {
	if g.len == 0 || num == 0 {
		return g
	}
	if num > g.len {
		num = g.len
	}
	g.elems = g.elems[:num]
	g.len = len(g.elems)
	return g
}

// Select 获取切片元素中指定属性
func (g *GLinq[T]) Select(callback func(T) any) *Query {
	if g.len == 0 {
		return nil
	}
	var elem T
	var results []any
	for i := range g.elems {
		elem = g.Get(i)
		v := callback(elem)
		results = append(results, v)
	}
	return &Query{items: results,len: len(results)}
}

func (q *Query) ToList(data interface{}) {
	marshal, _ := json.Marshal(q.items)
	err := json.Unmarshal(marshal, &data)
	if err != nil {
		panic(fmt.Sprintf("序列化异常 err:%v", err.Error()))
	}
}

// SelectI 获取切片元素中指定属性
func (g *GLinq[T]) SelectI(callback func(int, T) any) *Query {
	if g.len == 0 {
		return nil
	}
	var elem T
	var results []any
	for i := range g.elems {
		elem = g.Get(i)
		v := callback(i, elem)
		results = append(results, v)
	}
	return &Query{items: results}
}

// Get 根据索引获取元素
func (g *GLinq[T]) Get(index int) T {
	return g.elems[index]
}
