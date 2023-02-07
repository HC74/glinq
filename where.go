package glinq

func (q *Query) Where(callback func(any) bool) *Query {
	var results []any
	var item any
	for i := range q.items {
		item = q.Get(i)
		ok := callback(item)
		if ok {
			results = append(results, item)
		}
	}
	q.items = results
	q.len = len(q.items)
	return q
}

// Where 过滤条件,入参不包含索引
func (g *GLinq[T]) Where(callback func(T) bool) *GLinq[T] {
	var results []T
	var elem T
	for i := range g.elems {
		elem = g.elems[i]
		if callback(elem) {
			results = append(results, elem)
		}
	}
	g.elems = results
	g.len = len(results)
	return g
}

// WhereI 过滤条件,入参包含索引
func (g *GLinq[T]) WhereI(callback func(int, T) bool) *GLinq[T] {
	var results []T
	var elem T
	for i := range g.elems {
		elem = g.elems[i]
		if callback(i, elem) {
			results = append(results, elem)
		}
	}
	g.elems = results
	g.len = len(results)
	return g
}

// Distinct 去重
func (g *GLinq[T]) Distinct(callback func(item T) any) *GLinq[T] {
	// 使用map保持唯一
	var m = make(map[any]struct{})
	var elem T
	var results []T
	for i := range g.elems {
		elem = g.elems[i]
		k := callback(elem)
		if _, isOk := m[k]; !isOk {
			m[k] = struct{}{}
			results = append(results, elem)
		}
	}
	g.elems = results
	g.len = len(results)
	return g
}

// FirstWhere 查到第一个元素 根据条件,入参不包含索引
func (g *GLinq[T]) FirstWhere(callback func(item T) bool) (T, error) {
	result, err := g.Where(callback).First()
	return result, err
}

// FirstWhereI 查到第一个元素 根据条件,入参包含索引
func (g *GLinq[T]) FirstWhereI(callback func(int, T) bool) (T, error) {
	result, err := g.WhereI(callback).First()
	return result, err
}
