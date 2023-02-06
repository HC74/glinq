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
