package quickdp

import (
	PreTreat "MLCS_GO/src/pretreat"
	"time"

	"github.com/golang/glog"
)

// QuickDP for sth.
type QuickDP struct {
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
func (ths *QuickDP) Initialize() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* QuickDP 初始化")
	start := time.Now()

	pretreat := PreTreat.GetInstance()
	pretreat.ConstructSucTable()
	ths.doInitialize()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

// Execute for sth.
func (ths *QuickDP) Execute() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* QuickDP 开始运行")
	start := time.Now()

	ths.constructRelationGraph()
	ths.searchPath()
	ths.Print()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

func (ths *QuickDP) constructRelationGraph() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 创建DAG图")
	start := time.Now()

	ths.doConstructRelationGraph()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

func (ths *QuickDP) searchPath() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 寻找路径")
	start := time.Now()

	ths.doSearchLongestPath()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

// Print for sth.
func (ths *QuickDP) Print() {
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
