package utils

type HeapNode[Value interface{}] struct {
	Value  Value
	Weight int
	Idx    int
}

type Heap[Value interface{}] []*HeapNode[Value]

func (v Heap[Value]) Len() int {
	return len(v)
}

func (v Heap[Value]) Less(i, j int) bool {
	return v[i].Weight < v[j].Weight
}

func (v Heap[Value]) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
	v[i].Idx = i
	v[j].Idx = j
}

func (v *Heap[Value]) Push(x interface{}) {
	*v = append(*v, x.(*HeapNode[Value]))
	x.(*HeapNode[Value]).Idx = v.Len() - 1
}

func (v *Heap[Value]) Pop() interface{} {
	old := *v
	n := len(old)
	x := old[n-1]
	*v = old[0 : n-1]
	x.Idx = -1
	return x
}
