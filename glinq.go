package glinq

type GLinq[T any] struct {
	elems []T
	len   int
}

func Instance[T any](source []T) *GLinq[T] {
	return &GLinq[T]{elems: source, len: len(source)}
}

// ToList 转换为集合数据返回
func (g *GLinq[T]) ToList() []T {
	return g.elems
}

// Count 获取集合中的元素数量
func (g *GLinq[T]) Count() int {
	return g.len
}

// AsSlice 快速转换为数组
func AsSlice[T any](datas ...T) []T {
	var results []T
	var item T
	for i := range datas {
		item = datas[i]
		results = append(results, item)
	}
	return results
}
