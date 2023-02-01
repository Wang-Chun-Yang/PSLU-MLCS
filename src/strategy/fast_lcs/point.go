package fastlcs

import "github.com/golang/glog"

// Point for sth.
type Point struct {
	Index    int
	Value    byte
	Pvt      []int
	Depth    int
	SucArray []*Point
}

// Print for sth.
func (ths *Point) Print() {
	var vt []int
	for i := 0; i < len(ths.Pvt); i++ {
		vt = append(vt, ths.Pvt[i]+1)
	}
	var vt1 []int
	for _, v := range ths.SucArray {
		vt1 = append(vt1, v.Index+1)
	}
	glog.V(100).Infof("编号: %v, 字符: %c, 层数: %v, 坐标: %v, 后继: %v\n",
		ths.Index+1,
		ths.Value,
		ths.Depth,
		vt,
		vt1,
	)
}
