package glinq

import "sort"

type sortCore[T any] struct {
	items []T
	cmp   func(T, T) bool
}

func (s sortCore[T]) Len() int { return len([]T(s.items)) }

// Swap 交换
func (s sortCore[T]) Swap(i, j int) { s.items[i], s.items[j] = s.items[j], s.items[i] }

// Less 比较
func (s sortCore[T]) Less(i, j int) bool { return s.cmp(s.items[i], s.items[j]) }

// Sort 排序
func (g *GLinq[T]) Sort(callback func(current, next T) bool) *GLinq[T] {
	core := &sortCore[T]{items: g.ToList(), cmp: callback}
	sort.Sort(core)
	return g
}
