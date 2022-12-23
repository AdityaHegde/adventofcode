package proboscidea_volcanium

type valve struct {
  nodeName string
  weight   int
  idx      int
}

type valveHeap []*valve

func (v valveHeap) Len() int {
  return len(v)
}

func (v valveHeap) Less(i, j int) bool {
  return v[i].weight > v[j].weight
}

func (v valveHeap) Swap(i, j int) {
  v[i], v[j] = v[j], v[i]
  v[i].idx = i
  v[j].idx = j
}

func (v *valveHeap) Push(x interface{}) {
  *v = append(*v, x.(*valve))
  x.(*valve).idx = v.Len() - 1
}

func (v *valveHeap) Pop() interface{} {
  old := *v
  n := len(old)
  x := old[n-1]
  *v = old[0 : n-1]
  x.idx = -1
  return x
}
