package astar

// SearchPriorityQueue A*搜索
type SearchPriorityQueue []*Point

func (pq SearchPriorityQueue) Len() int { return len(pq) }

func (pq SearchPriorityQueue) Less(i, j int) bool {
	rate1, rate2 := 1.0, 1.0
	sum1 := float64(pq[i].Depth)*rate1 + float64(pq[i].MaxLeftStep)*rate2
	sum2 := float64(pq[j].Depth)*rate1 + float64(pq[j].MaxLeftStep)*rate2

	if sum1 == sum2 {
		return pq[i].Depth > pq[j].Depth
	}
	return sum1 > sum2
}
func (pq SearchPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push for sth.
func (pq *SearchPriorityQueue) Push(x interface{}) {
	item := x.(*Point)
	*pq = append(*pq, item)
}

// Pop for sth.
func (pq *SearchPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// Front for sth.
func (pq *SearchPriorityQueue) Front() interface{} {
	x := (*pq)[0]
	return x
}
