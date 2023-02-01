package pretreat

import (
	Scenario "MLCS_GO/src/scenario"
	"bytes"
	"fmt"
	"strconv"

	"github.com/golang/glog"
)

func (ths *pretreat) doConstructSucTable() {
	sce := Scenario.GetInstance()
	sequences := sce.GetSequences()
	seqNum := sce.GetSeqNum()
	alpSet := sce.GetAlpSet()

	ths.sucTable = make([][]map[byte]int, seqNum)
	for i := 0; i < seqNum; i++ {
		seq := sequences[i]
		sz := len(seq)
		ST := make([]map[byte]int, sz)
		for j := 0; j < sz; j++ {
			ST[j] = make(map[byte]int)
		}
		for _, alp := range alpSet {
			ST[sz-1][alp] = ths.failedIndex
		}
		for j := int(sz - 1); j >= 1; j-- {
			for _, alp := range alpSet {
				ST[j-1][alp] = ST[j][alp]
			}
			ST[j-1][seq[j]] = (int)(j)
		}
		ths.sucTable[i] = ST
	}
}

func (ths *pretreat) doGetSucIndex(i, j int, alp byte) int {
	sce := Scenario.GetInstance()
	// assert.True(i >= 0 && i < (int)(sce.GetSeqNum()), "获取后继Index越界")
	// sz := (int)(len(sce.GetIDSequence((int)(i))))
	// assert.True(j == ths.startIndex || j == ths.failedIndex || (j >= 0 && j < sz))

	if j == ths.failedIndex {
		return j
	}
	if j == ths.startIndex {
		s := sce.GetIDSequence(i)
		if s[0] == alp {
			return 0
		}
		return ths.sucTable[i][0][alp]
	}
	return ths.sucTable[i][j][alp]
}

func (ths *pretreat) doConstructPredcTable() {
	sce := Scenario.GetInstance()
	sequences := sce.GetSequences()
	seqNum := sce.GetSeqNum()
	alpSet := sce.GetAlpSet()

	ths.predcTable = make([][]map[byte]int, seqNum)
	for i := 0; i < seqNum; i++ {
		seq := sequences[i]
		sz := len(seq)
		ST := make([]map[byte]int, sz)
		for j := 0; j < sz; j++ {
			ST[j] = make(map[byte]int)
		}
		for _, alp := range alpSet {
			ST[0][alp] = ths.failedIndex
		}
		for j := 0; j < (int)(sz-1); j++ {
			for _, alp := range alpSet {
				ST[j+1][alp] = ST[j][alp]
			}
			ST[j+1][seq[j]] = (int)(j)
		}
		ths.predcTable[i] = ST
	}
}
func (ths *pretreat) doGetPredcIndex(i, j int, alp byte) int {
	sce := Scenario.GetInstance()
	// assert.True(i >= 0 && i < (int)(sce.GetSeqNum()), "获取反向后继Index越界")
	// sz := (int)(len(sce.GetIDSequence((int)(i))))
	// assert.True(j == ths.startIndex || j == ths.failedIndex || (j >= 0 && j < sz))

	if j == ths.failedIndex {
		return j
	}
	if j == ths.startIndex {
		s := sce.GetIDSequence((int)(i))
		sz := len(s)
		if s[sz-1] == alp {
			return (int)(sz - 1)
		}
		return ths.predcTable[i][sz-1][alp]
	}
	return ths.predcTable[i][j][alp]
}

func (ths *pretreat) doConstructDoubleSucTable() {
	sce := Scenario.GetInstance()
	sequences := sce.GetSequences()
	seqNum := sce.GetSeqNum()
	alpSet := sce.GetAlpSet()
	doubleAlpSet := sce.GetDoubleAlpSet()

	ths.doubleSucTable = make([][]map[string]int, seqNum)
	for i := 0; i < seqNum; i++ {
		seq := sequences[i]
		sz := len(seq)
		ST := make([]map[string]int, sz)
		for j := 0; j < sz; j++ {
			ST[j] = make(map[string]int)
		}
		for _, alp := range doubleAlpSet {
			ST[sz-1][alp] = ths.failedIndex
		}
		for _, alp := range alpSet {
			ST[sz-1][string(alp)] = ths.failedIndex
		}

		for j := int(sz - 1); j >= 1; j-- {
			for _, alp := range doubleAlpSet {
				ST[j-1][alp] = ST[j][alp]
			}
			for _, alp := range alpSet {
				ST[j-1][string(alp)] = ST[j][string(alp)]
				s := fmt.Sprintf("%c%c", seq[j], alp)
				ST[j-1][s] = ST[j][string(alp)]
			}
			ST[j-1][string(seq[j])] = (int)(j)
		}
		ths.doubleSucTable[i] = ST
	}
}

func (ths *pretreat) doGetDoubleSucIndex(i, j int, ds string) int {
	sce := Scenario.GetInstance()
	// assert.True(i >= 0 && i < (int)(sce.GetSeqNum()), "获取后继Index越界")
	// assert.True(len(ds) == 2, "获取Double后继，字符串长度不为2")
	// sz := (int)(len(sce.GetIDSequence((int)(i))))
	// assert.True(j == ths.startIndex || j == ths.failedIndex || (j >= 0 && j < sz))

	if j == ths.failedIndex {
		return j
	}
	if j == ths.startIndex {
		s := sce.GetIDSequence((int)(i))
		if s[0] == ds[0] {
			return ths.doubleSucTable[i][0][string(ds[1])]
		}
		return ths.doubleSucTable[i][0][ds]
	}
	return ths.doubleSucTable[i][j][ds]
}

// Max for sth.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getDp(s1, s2 string) dp {
	n, m := len(s1), len(s2)
	var ans dp
	ans = make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
	}
	for i := n - 1; i >= 0; i-- {
		if s1[i] == s2[m-1] {
			ans[i][m-1] = 1
		} else {
			if i == n-1 {
				ans[i][m-1] = 0
			} else {
				ans[i][m-1] = ans[i+1][m-1]
			}
		}
	}
	for j := m - 1; j >= 0; j-- {
		if s1[n-1] == s2[j] {
			ans[n-1][j] = 1
		} else {
			if j == m-1 {
				ans[n-1][j] = 0
			} else {
				ans[n-1][j] = ans[n-1][j+1]
			}
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			if s1[i] == s2[j] {
				ans[i][j] = ans[i+1][j+1] + 1
			} else {
				ans[i][j] = Max(ans[i+1][j], ans[i][j+1])
			}
		}
	}
	return ans
}

func (ths *pretreat) doConstructHeuristicTable() {
	sce := Scenario.GetInstance()
	seqNum := sce.GetSeqNum()

	ths.heuristicTable = make([][]dp, seqNum)

	for i := 0; i < seqNum; i++ {
		ths.heuristicTable[i] = make([]dp, seqNum)
		for j := 0; j < seqNum; j++ {
			s1 := sce.GetIDSequence(i)
			s2 := sce.GetIDSequence(j)
			ths.heuristicTable[i][j] = getDp(s1, s2)
		}
	}
}

func (ths *pretreat) doGetHeuristicValue(i, p1 int, j, p2 int) int {
	// sce := Scenario.GetInstance()
	// assert.True(i >= 0 && i < sce.GetSeqNum(), "i不在序列集合内")
	// assert.True(j >= 0 && j < sce.GetSeqNum(), "j不在序列集合内")
	// assert.True(p1 >= 0 && p1 < len(sce.GetIDSequence(i)))
	// assert.True(p2 >= 0 && p2 < len(sce.GetIDSequence(j)))
	return ths.heuristicTable[i][j][p1][p2]
}

const (
	mod1 = 1610612741
	mod2 = 805306457
)

func (ths *pretreat) doGetSucVector(newVt []int, alp byte) (bool, *Result) {
	sce := Scenario.GetInstance()
	sz := len(newVt)
	// assert.True(sz == sce.GetSeqNum(), "结点向量个数错误")

	var ansVt []int
	var ansStr bytes.Buffer
	max1, maxid1, max2, maxid2 := 0, 0, 0, 0
	hash1, hash2 := int64(newVt[0]+1), int64(newVt[0]+1)
	base := int64(sce.GetMaxSeqLength() + 1)

	for i := 0; i < sz; i++ {
		x := ths.doGetSucIndex(i, int(newVt[i]), alp)
		if x == ths.failedIndex {
			return false, nil
		}
		if i != 0 {
			ansStr.WriteString(",")
		}
		sx := strconv.Itoa(x)
		ansStr.WriteString(sx)
		ansVt = append(ansVt, x)

		hash1 = ((hash1*(base+1))%mod1 + int64(x+1)) % mod1
		hash2 = ((hash2*(base+2))%mod2 + int64(x+1)) % mod2

		if x > max1 {
			max1 = x
			maxid1 = i
		} else if x > max2 {
			max2 = x
			maxid2 = i
		}
	}

	result := &Result{
		Vector: ansVt,
		String: ansStr.String(),
		Maxid1: maxid1,
		Maxid2: maxid2,
		Hash1:  hash1,
		Hash2:  hash2,
		Alp:    alp,
	}

	return true, result
}

func (ths *pretreat) doGetDoubleSucVector(newVt []int, ds string) ([]int, string, bool) {
	// sce := Scenario.GetInstance()
	sz := len(newVt)
	// assert.True((int)(sz) == sce.GetSeqNum(), "结点向量个数错误")

	var ansVt []int
	var ansStr bytes.Buffer
	for i := (int)(0); i < sz; i++ {
		x := ths.doGetDoubleSucIndex(i, int(newVt[i]), ds)
		if x == ths.failedIndex {
			return ansVt, ansStr.String(), false
		}
		if i != 0 {
			ansStr.WriteString(",")
		}
		sx := strconv.Itoa((int)(x))
		ansStr.WriteString(sx)
		ansVt = append(ansVt, int(x))
	}
	return ansVt, ansStr.String(), true
}

func (ths *pretreat) printSucTable() {
	if !*debugSucTable {
		return
	}
	sce := Scenario.GetInstance()
	sequences := sce.GetSequences()
	alpSet := sce.GetAlpSet()
	for i, seq := range sequences {
		glog.V(100).Infoln("======================================")
		glog.V(100).Infoln(seq)
		for _, alp := range alpSet {
			logInfo := fmt.Sprintf("%c: ", alp)
			for j := 0; j < len(seq); j++ {
				as := ths.GetSucIndex(i, j, alp) + 1
				tmp := fmt.Sprintf("%v ", as)
				logInfo += tmp
			}
			glog.V(100).Infoln(logInfo)
		}
	}
	glog.V(100).Infoln("======================================")
}
