package astar

import "github.com/golang/glog"

// PVector 图中结点的向量类型
type PVector []int

// Point 结点需要保存的信息
type Point struct {
	Index       int     `json:"index"`       // 结点标号
	Value       byte    `json:"value"`       // 结点字符
	Pvt         PVector `json:"pvector"`     // 向量
	Depth       int     `json:"depth"`       // 最大深度
	MaxLeftStep int     `json:"maxLeftStep"` // 最大剩余步数
	PredcArray  []int   `json:"predcArray"`  // 前驱集合
	IfLongest   bool    `json:"ifLongest"`   // 是否是最长路的
}

// Print for sth.
func (ths *Point) Print(msg string) {
	if !*debugDetail {
		return
	}
	var vt []int
	for i := 0; i < len(ths.Pvt); i++ {
		vt = append(vt, ths.Pvt[i]+1)
	}
	glog.V(100).Infof("> %v < 编号: %v, 字符: %c, 层数: %v, 剩余步数: %v, 坐标: %v\n",
		msg,
		ths.Index+1,
		ths.Value,
		ths.Depth,
		ths.MaxLeftStep,
		vt,
	)
}
