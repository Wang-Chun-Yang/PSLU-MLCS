package scenario

import (
	FileIO "MLCS_GO/src/common/fileio"
	"fmt"
	"math"
	"sort"
)

// Max for sth.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min for sth.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (ths *scenario) doReadSequences() {
	ths.sequences = FileIO.ReadFile()
	alpMap := make(map[byte]bool)
	ths.minSeqLength = math.MaxUint16
	ths.maxSeqLength = 0
	for _, seq := range ths.sequences {
		sz := int(len(seq))
		ths.maxSeqLength = Max(ths.maxSeqLength, sz)
		ths.minSeqLength = Min(ths.minSeqLength, sz)
		for _, alp := range seq {
			alpMap[byte(alp)] = true
		}
	}

	var keys []int
	for k := range alpMap {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		ths.alpSet = append(ths.alpSet, byte(k))
	}
	ths.seqNum = int(len(ths.sequences))

	for _, k1 := range keys {
		for _, k2 := range keys {
			s := fmt.Sprintf("%c%c", k1, k2)
			ths.doubleAlpSet = append(ths.doubleAlpSet, s)
		}
	}
}
