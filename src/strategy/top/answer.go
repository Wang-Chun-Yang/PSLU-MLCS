package top

import "github.com/golang/glog"

// Answer for sth.
type Answer struct {
	TolIndex   int
	TolEdgeNum int
	MLCS       []string
	MLCSLength int
}

// Print for sth.
func (ths *Answer) Print() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infof("* 点数: %v", ths.TolIndex)
	glog.V(100).Infof("* 边数: %v", ths.TolEdgeNum)
	glog.V(100).Infof("* MLCS个数: %v", len(ths.MLCS))
	glog.V(100).Infof("* MLCS长度: %v", ths.MLCSLength)
	// for index, v := range ths.MLCS {
	// 	glog.V(100).Infof("* %v: %v", index+1, v)
	// }
	glog.V(100).Infoln("------------------------------------------------------")
}
