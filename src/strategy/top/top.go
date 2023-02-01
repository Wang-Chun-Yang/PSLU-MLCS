package top

import (
	PreTreat "MLCS_GO/src/pretreat"
	"log"
	"runtime"
	"time"

	"github.com/golang/glog"
)

// Top for sth.
type Top struct {
	Queue      []*Point
	PointsDom  []*Point
	CreatedMap map[string]int
	answer     Answer
}

var (
	startPoint = &Point{}
	endPoint   = &Point{}
)

// Initialize for sth.
func (ths *Top) Initialize() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* Top 初始化")
	start := time.Now()

	pretreat := PreTreat.GetInstance()
	pretreat.ConstructSucTable()
	ths.doInitialize()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

// Execute for sth.
func (ths *Top) Execute() {
	ths.constructRelationGraph()
	ths.forwardTopologySort()
	ths.backwardTopologySort()
	ths.answer.Print()
}

func (ths *Top) constructRelationGraph() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 创建DAG图")
	start := time.Now()

	ths.doConstructRelationGraph()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

func (ths *Top) forwardTopologySort() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 前向拓扑排序")
	start := time.Now()
	ths.doForwardTopologySort()
	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
	printMemStats()
}

func (ths *Top) backwardTopologySort() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 后向拓扑排序")
	start := time.Now()
	ths.doBackwardTopologySort()
	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)


}