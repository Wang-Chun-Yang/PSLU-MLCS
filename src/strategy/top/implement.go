package top

import (
	PreTreat "MLCS_GO/src/pretreat"
	Scenario "MLCS_GO/src/scenario"
	"fmt"
)

func (ths *Top) doInitialize() {
	pretreat := PreTreat.GetInstance()
	scenario := Scenario.GetInstance()
	startIndex := pretreat.GetStartIndex()

	var vt []int
	seqNum := scenario.GetSeqNum()
	for i := 0; i < seqNum; i++ {
		vt = append(vt, startIndex)
	}

	startPoint = &Point{
		Index:      -3,
		Value:      '#',
		Pvt:        vt,
		Depth:      0,
		SucArray:   []int{},
		PredcArray: []int{},
	}

	endPoint = &Point{
		Index:      -4,
		Value:      '#',
		PredcArray: []int{},
	}

	ths.PointsDom = ths.PointsDom[0:0]
	ths.CreatedMap = make(map[string]int)
}

func (ths *Top) doConstructRelationGraph() {
	pretreat := PreTreat.GetInstance()
	scenario := Scenario.GetInstance()
	alpSet := scenario.GetAlpSet()

	ths.Queue = append(ths.Queue, startPoint)
	for len(ths.Queue) > 0 {
		head := ths.Queue[0]
		ths.Queue = ths.Queue[1:]

		sign := false
		for _, alp := range alpSet {
			success, result := pretreat.GetSucVector(head.Pvt, alp)
			if !success {
				continue
			}
			sign = true
			if _, ok := ths.CreatedMap[result.String]; ok {
				v := ths.CreatedMap[result.String]
				point := ths.PointsDom[v]
				head.SucArray = append(head.SucArray, v)
				point.PredcArray = append(point.PredcArray, head.Index)
				point.PredcDegree++
			} else {
				point := &Point{
					Index:       ths.answer.TolIndex,
					Value:       alp,
					Pvt:         result.Vector,
					Depth:       head.Depth + 1,
					SucArray:    []int{},
					PredcArray:  []int{head.Index},
					PredcDegree: 1,
				}
				ths.Queue = append(ths.Queue, point)
				ths.CreatedMap[result.String] = point.Index
				head.SucArray = append(head.SucArray, point.Index)
				ths.PointsDom = append(ths.PointsDom, point)
				ths.answer.TolIndex++
				ths.answer.TolEdgeNum += 2
			}
		}
		if !sign {
			head.SucArray = append(head.SucArray, endPoint.Index)
			endPoint.PredcArray = append(endPoint.PredcArray, head.Index)
			endPoint.PredcDegree++
			ths.answer.TolEdgeNum += 2
		}
	}
}

func (ths *Top) doForwardTopologySort() {
	var queue []*Point
	queue = append(queue, startPoint)
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		ths.answer.MLCSLength = Max(ths.answer.MLCSLength, head.Depth)
		for _, v := range head.SucArray {
			if v == endPoint.Index {
				continue
			}

			point := ths.PointsDom[v]
			point.PredcDegree--
			if point.PredcDegree == 0 {
				point.Depth = head.Depth + 1
				queue = append(queue, point)
			}
		}
	}
}

func getString(a []byte) string {
	ans := ""
	len := len(a)
	for i := len - 1; i >= 0; i-- {
		ans += string(fmt.Sprintf("%c", a[i]))
	}
	return ans
}

func (ths *Top) doBackwardTopologySort() {
	var dfs func(fa *Point, step int)
	var path []byte
	dfs = func(fa *Point, step int) {
		for _, u := range fa.PredcArray {
			if u == startPoint.Index {
				ths.answer.MLCS = append(ths.answer.MLCS, getString(path))
				continue
			}
			point := ths.PointsDom[u]
			if point.Depth+step != ths.answer.MLCSLength {
				continue
			}
			path = append(path, point.Value)
			dfs(point, step+1)
			path = path[0 : len(path)-1]
		}
	}
	dfs(endPoint, 0)
}
