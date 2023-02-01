package simulator

import (
	"MLCS_GO/src/strategy"
	"time"

	Scenario "MLCS_GO/src/scenario"

	"github.com/golang/glog"
)

// Simulator for sth.
type Simulator struct {
	Factory strategy.Factory
}

// RunFrameWork for sth.
func (ths *Simulator) RunFrameWork() {
	start := time.Now()

	sce := Scenario.GetInstance()
	sce.ReadSequences()

	strategy := ths.Factory.CreateStrategy()
	strategy.Initialize()
	strategy.Execute()

	// 打印运行时间
	elapsed := time.Since(start).Seconds()
	glog.V(100).Infof("* [总时间: %.3fs]", elapsed)
}
