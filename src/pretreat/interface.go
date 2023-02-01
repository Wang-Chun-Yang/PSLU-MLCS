package pretreat

import (
	Scenario "MLCS_GO/src/scenario"
	"flag"

	"github.com/golang/glog"
)

// Imp 调用接口
type Imp interface {
	ConstructSucTable()                                     // 创建后继表
	ConstructPredcTable()                                   // 创建前驱表
	ConstructPreDPTable()                                   // 创建股价函数表
	GetSucIndex(i, j int, alp byte) int                     // 第i个序列第j个位置往后(>j)字符alp的后继位置
	GetPredcIndex(i, j int, alp byte) int                   // 第i个序列第j个位置往前(<j)字符alp的前驱位置
	GetStartIndex() int                                     // 获取ST表开始坐标
	GetFailedIndex() int                                    // 获取ST表失败坐标(找不到后继或者前驱)
	GetSucVector(pvt []int, alp byte) ([]int, string, bool) // 获取向量pvt在字符alp下的后继坐标
	GetHeuristicValue(i, p1 int, j, p2 int) int             // 获取启发式函数值i序列p1位置，j序列p2位置
}

type dp [][]int

// pretreat 数据结构
type pretreat struct {
	startIndex     int                // 后继表前驱表开始坐标 -2
	failedIndex    int                // 后继表前驱表失败坐标 -1
	sucTable       [][]map[byte]int   // 后继表[i][j][alp]
	predcTable     [][]map[byte]int   // 前驱表[i][j][alp]
	heuristicTable [][]dp             // 预处理任意两个字符串DP表
	doubleSucTable [][]map[string]int // 双步长后继表[i][j]["c1c2"]
}

// Result for sth.
type Result struct {
	Vector []int
	String string
	Hash1  int64
	Hash2  int64
	Maxid1 int
	Maxid2 int
	Alp    byte
}

var instance *pretreat

// GetInstance 单利模式调用
func GetInstance() *pretreat {
	if instance == nil {
		instance = &pretreat{startIndex: -2, failedIndex: -1}
	}
	return instance
}

var (
	debugSucTable       = flag.Bool("debug_suctable", false, "qaq")
	needHeuristicLength = flag.Int("need_heuristic_length", 0, "")
)

// ConstructSucTable for sth.
func (ths *pretreat) ConstructSucTable() {
	glog.V(100).Infoln("* 构造后继表")
	ths.doConstructSucTable()
	ths.printSucTable()
}

// ConstructPredcTable for sth.
func (ths *pretreat) ConstructPredcTable() {
	glog.V(100).Infoln("* 构造前驱表")
	ths.doConstructPredcTable()
}

// ConstructPreLCSTable for sth.
func (ths *pretreat) ConstructHeuristicTable() {
	if Scenario.GetInstance().GetSeqNum() > *needHeuristicLength {
		return
	}
	glog.V(100).Infoln("* 构造估价函数表")
	ths.doConstructHeuristicTable()
}

// ConstructDoubleSucTable for sth.
func (ths *pretreat) ConstructDoubleSucTable() {
	glog.V(100).Infoln("* 构造双步长后继表")
	ths.doConstructDoubleSucTable()
}

// GetSucIndex for sth.
func (ths *pretreat) GetSucIndex(i, j int, alp byte) int {
	return ths.doGetSucIndex(i, j, alp)
}

// GetDoubleSucIndex for sth.
func (ths *pretreat) GetDoubleSucIndex(i, j int, ds string) int {
	return ths.doGetDoubleSucIndex(i, j, ds)
}

// GetPredcIndex for sth.
func (ths *pretreat) GetPredcIndex(i, j int, alp byte) int {
	return ths.doGetPredcIndex(i, j, alp)
}

// GetStartIndex for sth.
func (ths *pretreat) GetStartIndex() int {
	return ths.startIndex
}

// GetFailedIndex for sth.
func (ths *pretreat) GetFailedIndex() int {
	return ths.failedIndex
}

// GetSucVector for sth.
func (ths *pretreat) GetSucVector(newVt []int, alp byte) (bool, *Result) {
	return ths.doGetSucVector(newVt, alp)
}

// GetDoubleSucVector for sth.
func (ths *pretreat) GetDoubleSucVector(newVt []int, ds string) ([]int, string, bool) {
	return ths.doGetDoubleSucVector(newVt, ds)
}

func (ths *pretreat) GetHeuristicValue(i, p1 int, j, p2 int) int {
	return ths.doGetHeuristicValue(i, p1, j, p2)
}

func (ths *pretreat) GetNeedHeuristicLength() int {
	return *needHeuristicLength
}
