package fastlcs

import (
	PreTreat "MLCS_GO/src/pretreat"
	Scenario "MLCS_GO/src/scenario"

	"fmt"

	"github.com/golang/glog"
)

func (ths *FastLCS) doInitialize() {
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

func (ths *FastLCS) createNewNode(sons Sons, fathers map[string][]*Point) {
	ths.Queue = ths.Queue[0:0]
	judgeLess := func(v1 []int, v2 []int) bool {
		sz := len(v1)
		for i := 0; i < sz; i++ {
			if v1[i] >= v2[i] {
				return false
			}
		}
		return true
	}
	for k1, v1 := range sons {
		sign := true
		for k2, v2 := range sons {
			if k1 == k2 {
				continue
			}
			if judgeLess(v2.Vector, v1.Vector) {
				sign = false
				break
			}
		}
		if sign {
			father := fathers[v1.String]
			point := &Point{
				Index:    ths.TolIndex,
				Pvt:      v1.Vector,
				Depth:    father[0].Depth + 1,
				SucArray: []*Point{},
				Value:    v1.Alp,
			}
			ths.TolIndex++
			ths.Queue = append(ths.Queue, point)
			for _, fa := range father {
				ths.TolEdgeNum++
				fa.SucArray = append(fa.SucArray, point)
			}
		}
	}
}

func (ths *FastLCS) doConstructRelationGraph() {
	scenario := Scenario.GetInstance()
	pretreat := PreTreat.GetInstance()
	alpSet := scenario.GetAlpSet()
	ths.Queue = append(ths.Queue, startPoint)
	for len(ths.Queue) > 0 {
		glog.V(100).Infof("* 层数: %v, 结点数目: %v", ths.LongestPath, len(ths.Queue))

		var sons Sons
		fathers := make(map[string][]*Point)

		for _, head := range ths.Queue {
			ths.LongestPath = Max(ths.LongestPath, head.Depth)
			sign := false
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
					fathers[result.String] = []*Point{head}
					sons = append(sons, result)
				}
			}
			if !sign {
				head.SucArray = append(head.SucArray, endPoint)
			}
		}
		ths.createNewNode(sons, fathers)
	}
}

func (ths *FastLCS) doSearchLongestPath() {
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
