package scenario

import (
	"github.com/golang/glog"
)

// Imp 调用接口
type Imp interface {
	Print()                     // 打印日志
	ReadSequences()             // 读入数据
	GetSequences() []string     // 获取读入的所有序列
	GetIDSequence(p int) string // 获取第p个序列
	GetAlpSet() []byte          // 获取字符集合
	GetMaxSeqLength() int       // 序列最大长度
	GetMinSeqLength() int       // 序列最小长度
	GetSeqNum() int             // 序列个数
}

type scenario struct {
	maxSeqLength int      // 序列最大长度
	minSeqLength int      // 序列最小长度
	alpSet       []byte   // 字符集合
	sequences    []string // 序列集合
	seqNum       int      // 序列个数
	doubleAlpSet []string // 双步长序列集合
}

var instance *scenario

// GetInstance for sth.
func GetInstance() *scenario {
	if instance == nil {
		instance = &scenario{}
	}
	return instance
}

// LogSeq 打印读入序列信息
func (ths *scenario) Print() {
	glog.V(100).Infoln("------------------------------------------------------")
	glog.V(100).Infoln("* Scenario")
	glog.V(100).Infof("* 序列个数: %v", ths.seqNum)
	glog.V(100).Infof("* 最小长度: %v", ths.minSeqLength)
	glog.V(100).Infof("* 最大长度: %v", ths.maxSeqLength)
	glog.V(100).Infof("* 字符集合: %s", ths.alpSet)
	glog.V(100).Infoln("------------------------------------------------------")
}

// ReadSequences 读入序列
func (ths *scenario) ReadSequences() {
	ths.doReadSequences()
	ths.Print()
}

// GetSequences for sth.
func (ths *scenario) GetSequences() []string {
	return ths.sequences
}

// GetIDSequence for sth.
func (ths *scenario) GetIDSequence(p int) string {
	// assert.True(p >= 0 && p < instance.seqNum, "序列下标越界")
	return ths.sequences[p]
}

// GetAlpSet for sth.
func (ths *scenario) GetAlpSet() []byte {
	return ths.alpSet
}

// GetDoubleAlpSet for sth.
func (ths *scenario) GetDoubleAlpSet() []string {
	return ths.doubleAlpSet
}

// GetMaxSeqLength for sth.
func (ths *scenario) GetMaxSeqLength() int {
	return ths.maxSeqLength
}

// GetMinSeqLength for sth.
func (ths *scenario) GetMinSeqLength() int {
	return ths.minSeqLength
}

// GetSeqNum for sth.
func (ths *scenario) GetSeqNum() int {
	return ths.seqNum
}
