package astar

import (
	PreTreat "MLCS_GO/src/pretreat"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/golang/glog"
)

var (
	startPoint = &Point{}
	endPoint   = &Point{}
)

// Astar 搜索类
type Astar struct {
	OpenTable  SearchPriorityQueue // OPEN表
	CloseTable mapset.Set          // CLOSE表
	PointsDom  []*Point            // 所有结点集合
	CreatedMap map[string]int      // Hash表
	answer     Answer              // 记录答案
}

// Initialize 初始化[创建后继表，创建股价函数表，变量初始化]
func (ths *Astar) Initialize() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* Astar 初始化")
	start := time.Now()

	pretreat := PreTreat.GetInstance()
	pretreat.ConstructSucTable()
	pretreat.ConstructHeuristicTable()
	ths.doInitialize()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

// Execute 执行[创建DAG关系图, 寻找最长路]
func (ths *Astar) Execute() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* Astar 开始运行")
	start := time.Now()

	ths.constructRelationGraph()
	ths.findLongestPath()
	ths.answer.Print()
	ths.dump()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

func (ths *Astar) constructRelationGraph() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 创建 DAG 关系图")
	start := time.Now()

	ths.doConstructRelationGraph()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}

func (ths *Astar) findLongestPath() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* 寻找最长路")
	start := time.Now()

	ths.doFindLongestPath()

	time := time.Since(start).Seconds()
	glog.V(100).Infof("* 运行时间: %.3fs", time)
	glog.V(100).Infoln("------------------------------------------------------")
}
