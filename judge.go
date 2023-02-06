package glinq

// Empty 集合是否为空
func (g *GLinq[T]) Empty() bool {
	return g.len == 0
}

// Any 是否包含任意元素
func (g *GLinq[T]) Any() bool {
	return g.len > 0
}

// AnyN 是否包含任意元素根据条件
func (g *GLinq[T]) AnyN(callback func(i int, item T) bool) bool {
	var elem T
	for i := range g.elems {
		elem = g.elems[i]
		if callback(i, elem) {
			return true
		}
	}
	return false
}
