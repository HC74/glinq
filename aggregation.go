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


// Avg 求平均值
func (g *GLinq[T]) Avg(callback func(T) float64) float64 {
	var item T
	var sum float64 = 0
	for i := range g.elems {
		item = g.Get(i)
		sum += callback(item)
	}
	return sum / float64(g.len)
}

// AvgI 求平均值 入参中含索引
func (g *GLinq[T]) AvgI(callback func(int, T) float64) float64 {
	var item T
	var sum float64 = 0
	for i := range g.elems {
		item = g.Get(i)
		sum += callback(i, item)
	}
	return sum / float64(g.len)
}
