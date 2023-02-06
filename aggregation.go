package glinq

// Max 求最大值
func (g *GLinq[T]) Max(callback func(T) int) T {
	var max = -1
	var maxV T
	for i := range g.elems {
		currentNum := callback(g.elems[i])
		if currentNum > max {
			max = currentNum
			maxV = g.elems[i]
		}
	}
	return maxV
}

// Min 求最小值
func (g *GLinq[T]) Min(callback func(T) int) T {
	if g.Empty() {
		panic("集合中无元素")
	}
	var min = callback(g.elems[0])
	var minV T
	for i := range g.elems {
		currentNum := callback(g.elems[i])
		if currentNum < min {
			min = currentNum
			minV = g.elems[i]
		}
	}
	return minV
}
