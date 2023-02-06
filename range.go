package glinq

// ForEach 循环
func (g *GLinq[T]) ForEach(callback func(index int, item T)) {
	var elem T
	for i := range g.elems {
		elem = g.elems[i]
		callback(i, elem)
	}
}
