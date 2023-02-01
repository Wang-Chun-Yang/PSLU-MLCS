package astar

import (
	DPLength "MLCS_GO/src/common/dplength"
	DumpJson "MLCS_GO/src/common/dumpjson"
	PreTreat "MLCS_GO/src/pretreat"
	Scenario "MLCS_GO/src/scenario"
	"container/heap"
	"fmt"
	"math"

	mapset "github.com/deckarep/golang-set"
	"github.com/golang/glog"
)

// Initialize 初始化
func (ths *Astar) doInitialize() {
	pretreat := PreTreat.GetInstance()
	sce := Scenario.GetInstance()
	startIndex := (int)(pretreat.GetStartIndex())
	var startVt PVector
	seqNum := sce.GetSeqNum()
	for i := int(0); i < seqNum; i++ {
		startVt = append(startVt, int(startIndex))
	}

	startPoint = &Point{
		Index:       -3,
		Value:       '#',
		Pvt:         startVt,
		Depth:       0,
		MaxLeftStep: sce.GetMaxSeqLength(),
		IfLongest:   false,
	}

	endPoint = &Point{
		Index:      -4,
		PredcArray: []int{},
		Value:      '$',
		IfLongest:  false,
	}

	ths.OpenTable = make(SearchPriorityQueue, 0)
	heap.Init(&ths.OpenTable)
	ths.CloseTable = mapset.NewSet()
	ths.PointsDom = ths.PointsDom[0:0]
	ths.CreatedMap = make(map[string]int)
}

func (ths *Astar) updateState(head *Point) {
	ths.answer.TolSearchCount++
	ths.answer.MaxPathLength = Max(ths.answer.MaxPathLength, head.Depth)
	ths.answer.LogStep()
}

func (ths *Astar) matchExplore(head *Point) bool {
	if ths.CloseTable.Contains(head.Index) {
		return false
	}
	ths.CloseTable.Add(head.Index)
	return true
}

func (ths *Astar) calMaxLeftStep(result *PreTreat.Result) int {
	sce := Scenario.GetInstance()
	pretreat := PreTreat.GetInstance()

	if len(result.Vector) > pretreat.GetNeedHeuristicLength() {
		if result.Vector[result.Maxid1] == len(sce.GetIDSequence(result.Maxid1))-1 {
			return 0
		}
		if result.Vector[result.Maxid1] == len(sce.GetIDSequence(result.Maxid2))-1 {
			return 0
		}
		s1 := sce.GetIDSequence(result.Maxid1)
		s2 := sce.GetIDSequence(result.Maxid2)
		sz1, sz2 := len(s1), len(s2)
		return DPLength.DP(
			s1[result.Vector[result.Maxid1]+1:sz1],
			s2[result.Vector[result.Maxid2]+1:sz2],
		)
	}

	sz, maxLeftStep := len(result.Vector), math.MaxInt32
	for i := 0; i < sz; i++ {
		for j := i + 1; j < sz; j++ {
			if result.Vector[i] == len(sce.GetIDSequence(i))-1 {
				return 0
			}
			if result.Vector[j] == len(sce.GetIDSequence(j))-1 {
				return 0
			}
			x := pretreat.GetHeuristicValue(i, result.Vector[i]+1, j, result.Vector[j]+1)
			maxLeftStep = Min(maxLeftStep, x)
		}
	}
	return maxLeftStep
}

func (ths *Astar) matchOneStep(head *Point) bool {
	sce := Scenario.GetInstance()
	pretreat := PreTreat.GetInstance()
	alpSet := sce.GetAlpSet()
	sign := false

	for _, alp := range alpSet {
		success, result := pretreat.GetSucVector(head.Pvt, alp)
		if !success {
			continue
		}
		sign = true

		value1, ok1 := ths.CreatedMap[result.String]

		if ok1 {
			oldPoint := ths.PointsDom[value1]
			ths.answer.MaxPathLength = Max(ths.answer.MaxPathLength, head.Depth+1)

			if head.Depth+1 > oldPoint.Depth {
				sz := len(oldPoint.PredcArray)

				ths.answer.TolEdgeNum = ths.answer.TolEdgeNum - sz + 1
				ths.answer.RelaxedCount++
				oldPoint.PredcArray = []int{head.Index}
				oldPoint.Depth = head.Depth + 1

				heap.Push(&ths.OpenTable, oldPoint)
				ths.CloseTable.Remove(value1)

				oldPoint.Print("更优")
			} else if head.Depth+1 == oldPoint.Depth {
				ths.answer.TolEdgeNum++
				oldPoint.PredcArray = append(oldPoint.PredcArray, head.Index)
				oldPoint.Print("相同")
			}
		} else {
			maxLeftStep := ths.calMaxLeftStep(result)
			if head.Depth+1+maxLeftStep < ths.answer.MaxPathLength {
				for i := 0; i < len(result.Vector); i++ {
					result.Vector[i]++
				}
				glog.V(200).Infof("* 删除: %v, 剩余步数: %v", result.Vector, maxLeftStep)
				continue
			}
			point := &Point{
				Index:       ths.answer.TolIndex,
				Value:       alp,
				Pvt:         result.Vector,
				Depth:       head.Depth + 1,
				MaxLeftStep: maxLeftStep,
				PredcArray:  []int{head.Index},
			}
			ths.CreatedMap[result.String] = point.Index
			ths.answer.TolEdgeNum++
			ths.PointsDom = append(ths.PointsDom, point)
			heap.Push(&ths.OpenTable, point)
			ths.answer.TolIndex++
			point.Print("新建")
		}
	}
	return sign
}

var (
	have = mapset.NewSet()
)

func (ths *Astar) connectEndNode(head *Point) {
	if head.Depth >= ths.answer.MaxPathLength && !have.Contains(head.Index) {
		ths.answer.TolEdgeNum++
		endPoint.PredcArray = append(endPoint.PredcArray, head.Index)
		have.Add(head.Index)
		head.Print("终点")
	}
}

func (ths *Astar) doConstructRelationGraph() {
	heap.Push(&ths.OpenTable, startPoint)

	for ths.OpenTable.Len() > 0 {
		head := heap.Pop(&ths.OpenTable).(*Point)
		ths.updateState(head)
		if !ths.matchExplore(head) {
			continue
		}
		head.Print("P*")
		if ths.matchOneStep(head) {
			continue
		}
		ths.connectEndNode(head)
	}
}

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

func (ths *Astar) doFindLongestPath() {
	var dfs func(u, fa int, step int)
	var path []int
	st := mapset.NewSet()

	dfs = func(u, fa int, step int) {
		var predc []int

		if u == endPoint.Index {
			predc = endPoint.PredcArray
		} else {
			predc = ths.PointsDom[u].PredcArray
		}

		for _, v := range predc {
			if fa == v {
				continue
			}
			if v == startPoint.Index {
				s := ""
				for _, v := range path {
					ths.PointsDom[v].IfLongest = true
					s += (string)(ths.PointsDom[v].Value)

					if !st.Contains(v) {
						ths.answer.LongestPointNum++
						st.Add(v)
					}
				}
				ths.answer.MLCS = append(ths.answer.MLCS, reverse(s))
				return
			}
			if ths.PointsDom[v].Depth+step != ths.answer.MaxPathLength {
				continue
			}
			path = append(path, v)
			dfs(v, u, step+1)
			sz := len(path)
			path = path[:sz-1]
		}
	}

	pretreat := PreTreat.GetInstance()
	dfs(endPoint.Index, pretreat.GetFailedIndex(), 0)
}

func (ths *Astar) dump() {
	content := ths.PointsDom
	content = append(content, endPoint)
	DumpJson.Dump(content)
}
