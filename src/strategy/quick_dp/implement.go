package quickdp

import (
	"MLCS_GO/src/pretreat"
	PreTreat "MLCS_GO/src/pretreat"
	Scenario "MLCS_GO/src/scenario"

	"github.com/golang/glog"

	"fmt"
)

func (ths *QuickDP) doInitialize() {
	pretreat := PreTreat.GetInstance()
	scenario := Scenario.GetInstance()
	startIndex := pretreat.GetStartIndex()

	var vt []int
	seqNum := scenario.GetSeqNum()
	for i := 0; i < seqNum; i++ {
		vt = append(vt, startIndex)
	}

	startPoint = &Point{
		Index:    -3,
		Pvt:      vt,
		Depth:    0,
		SucArray: []*Point{},
	}

	endPoint = &Point{
		Index: -4,
	}
	ths.TolIndex = 0
}

func (ths *QuickDP) mySort(sons []*PreTreat.Result) []*PreTreat.Result {
	judgeLess := func(v1 []int, v2 []int) bool {
		sz := len(v1)
		for i := 0; i < sz; i++ {
			if v1[i] >= v2[i] {
				return false
			}
		}
		return true
	}
	var ans []*pretreat.Result
	for i, it1 := range sons {
		sign := true
		for j, it2 := range sons {
			if i == j {
				continue
			}
			if judgeLess(it2.Vector, it1.Vector) {
				sign = false
				break
			}
		}
		if sign {
			ans = append(ans, it1)
		}
	}
	return ans
}

func (ths *QuickDP) createNewNode(sons []*PreTreat.Result, fathers map[string][]*Point) {
	ths.Queue = ths.Queue[0:0]
	for _, son := range sons {
		point := &Point{
			Index:    ths.TolIndex,
			Pvt:      son.Vector,
			Depth:    fathers[son.String][0].Depth + 1,
			SucArray: []*Point{},
			Value:    son.Alp,
		}
		ths.TolIndex++
		ths.Queue = append(ths.Queue, point)
		for _, fa := range fathers[son.String] {
			ths.TolEdgeNum++
			fa.SucArray = append(fa.SucArray, point)
		}
	}
}

func (ths *QuickDP) doConstructRelationGraph() {
	scenario := Scenario.GetInstance()
	pretreat := PreTreat.GetInstance()
	alpSet := scenario.GetAlpSet()
	ths.Queue = append(ths.Queue, startPoint)
	for len(ths.Queue) > 0 {
		glog.V(100).Infof("* 层数: %v, 结点数目: %v", ths.LongestPath, len(ths.Queue))

		var totalSons []*PreTreat.Result
		fathers := make(map[string][]*Point)
		for _, head := range ths.Queue {
			ths.LongestPath = Max(ths.LongestPath, head.Depth)
			sign := false
			var sons []*PreTreat.Result
			for _, alp := range alpSet {
				success, result := pretreat.GetSucVector(head.Pvt, alp)
				if !success {
					continue
				}
				sign = true
				_, ok := fathers[result.String]
				if ok {
					fathers[result.String] = append(fathers[result.String], head)
				} else {
					sons = append(sons, result)
					fathers[result.String] = []*Point{head}
				}
			}
			if !sign {
				head.SucArray = append(head.SucArray, endPoint)
			}
			totalSons = append(totalSons, ths.mySort(sons)...)
		}
		totalSons = ths.mySort(totalSons)
		ths.createNewNode(totalSons, fathers)
	}
}

func (ths *QuickDP) doSearchLongestPath() {
	var dfs func(father *Point, path string)
	dfs = func(father *Point, path string) {
		for _, son := range father.SucArray {
			if son.Index == endPoint.Index {
				if len(path) == ths.LongestPath {
					ths.MLCS = append(ths.MLCS, path)
				}
				continue
			}
			dfs(son, path+string(fmt.Sprintf("%c", son.Value)))
		}
	}
	dfs(startPoint, "")
}
