package astar

import (
	Scenario "MLCS_GO/src/scenario"

	"flag"

	"github.com/golang/glog"
)

// Answer for sth.
type Answer struct {
	TolIndex        int         // 结点总数
	TolEdgeNum      int         // 边总数
	MaxPathLength   int         // 最大深度
	TolSearchCount  int         // 总搜索次数
	RelaxedCount    int         // 松弛次数
	RemoveCount     int         // 删除点数目
	RelaxedTolCount map[int]int // 松弛总计
	MLCS            []string    // 答案
	RemoveInfo      []int       // 删除点的记录信息
	LongestPointNum int         // 最长路上的结点数目
	FirstToMaxDepth int         // 第一次找到一条最长路创建了多少结点
}

// Initialize for sth.
func (ths *Answer) Initialize() {
	sce := Scenario.GetInstance()
	ths.TolIndex = 0
	ths.TolEdgeNum = 0
	ths.MaxPathLength = 0
	ths.TolSearchCount = 0
	ths.RelaxedCount = 0
	ths.RemoveCount = 0
	ths.RelaxedTolCount = make(map[int]int)
	ths.MLCS = ths.MLCS[0:0]
	ths.RemoveInfo = make([]int, sce.GetMaxSeqLength())
	ths.LongestPointNum = 0
	ths.FirstToMaxDepth = 0
}

// Print for sth.
func (ths *Answer) Print() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infof("* 点数: %v", ths.TolIndex)
	glog.V(100).Infof("* 边数: %v", ths.TolEdgeNum)
	glog.V(100).Infof("* MLCS个数: %v", len(ths.MLCS))
	glog.V(100).Infof("* MLCS长度: %v", ths.MaxPathLength)
	// for index, v := range ths.MLCS {
	// 	glog.V(100).Infof("* %v: %v", index+1, v)
	// }
	glog.V(100).Infoln("------------------------------------------------------")
}

var (
	lastLongestPath = 0
	lastPointNum    = 0
	debugDetail     = flag.Bool("debug_detail", false, "qaq")
)

// LogStep for sth.
func (ths *Answer) LogStep() {
	if *debugDetail {
		return
	}
	if ths.MaxPathLength > lastLongestPath || (ths.TolIndex > lastPointNum && ths.TolIndex%100000 == 0) {
		glog.V(200).Infof("* 层数: %v, 点数: %v, 边数: %v, 松弛: %v",
			ths.MaxPathLength,
			ths.TolIndex,
			ths.TolEdgeNum,
			ths.RelaxedCount,
		)
		lastLongestPath = ths.MaxPathLength
		lastPointNum = ths.TolIndex
	}
}
