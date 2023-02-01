package top

import "github.com/golang/glog"

// Point for sth.
type Point struct {
	Index       int
	Value       byte
	Pvt         []int
	Depth       int
	SucArray    []int
	PredcArray  []int
	PredcDegree int
}

// Print for sth.
func (ths *Point) Print() {
	var vt []int
	for i := 0; i < len(ths.Pvt); i++ {
		vt = append(vt, ths.Pvt[i]+1)
	}
	glog.V(100).Infof("编号: %v, 字符: %c, 层数: %v, 坐标: %v, 前驱: %v\n",
		ths.Index+1,
		ths.Value,
		ths.Depth,
		vt,
		ths.PredcArray,
		// ths.PredcDegree,
		// ths.SucArray,
	)
}
