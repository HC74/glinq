package glinq

// FlatMapI 转map 键为int
func (g *GLinq[T]) FlatMapI(callback func(int, T) int) map[int]T {
	results := make(map[int]T)
	var elem T
	for i := range g.elems {
		elem = g.elems[i]
		key := callback(i, elem)
		results[key] = elem
	}
	return results
}

// FlatMapS 转map 键为string
func (g *GLinq[T]) FlatMapS(callback func(int, T) string) map[string]T {
	results := make(map[string]T)
	var elem T
	for i := range g.elems {
		elem = g.elems[i]
		key := callback(i, elem)
		results[key] = elem
	}
	return results
}
