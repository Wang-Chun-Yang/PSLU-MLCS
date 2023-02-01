package fastlcs

import PreTreat "MLCS_GO/src/pretreat"

// Sons for sort
type Sons []*PreTreat.Result

func (ms Sons) Len() int {
	return len(ms)
}

func (ms Sons) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms Sons) Less(i, j int) bool {
	sz := len(ms[i].Vector)
	for k := 0; k < sz; k++ {
		if ms[i].Vector[k] == ms[j].Vector[k] {
			continue
		}
		return ms[i].Vector[k] < ms[j].Vector[k]
	}
	return false
}
