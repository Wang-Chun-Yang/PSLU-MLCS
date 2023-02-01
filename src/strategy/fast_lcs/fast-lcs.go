package fastlcs

import (
	PreTreat "MLCS_GO/src/pretreat"
	"time"

	"github.com/golang/glog"
)

// FastLCS for sth.
type FastLCS struct {
	Queue       []*Point
	TolIndex    int
	LongestPath int
	TolEdgeNum  int
	MLCS        []string
}

var (
	startPoint = &Point{}
	endPoint   = &Point{}
)

// Initialize for sth.
func (ths *FastLCS) Initialize() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* FAST_LCS 初始化")
	start := time.Now()

	pretreat := PreTreat.GetInstance()
	pretreat.ConstructSucTable()
	ths.doInitialize()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

// Execute for sth.
func (ths *FastLCS) Execute() {
	ths.constructRelationGraph()
	ths.searchPath()
	ths.Print()
}

func (ths *FastLCS) constructRelationGraph() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 创建DAG图")
	start := time.Now()

	ths.doConstructRelationGraph()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

func (ths *FastLCS) searchPath() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 寻找路径")
	start := time.Now()

	ths.doSearchLongestPath()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

// Print for sth.
func (ths *FastLCS) Print() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infof("* 点数: %v", ths.TolIndex)
	glog.V(100).Infof("* 边数: %v", ths.TolEdgeNum)
	glog.V(100).Infof("* MLCS个数: %v", len(ths.MLCS))
	glog.V(100).Infof("* MLCS长度: %v", ths.LongestPath)
	// for index, v := range ths.MLCS {
	// 	glog.V(100).Infof("* %v: %v", index+1, v)
	// }
	glog.V(100).Infoln("------------------------------------------------------")
}
