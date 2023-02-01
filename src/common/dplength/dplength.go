package dplength

import (
	Scenario "MLCS_GO/src/scenario"
)

// Max for sth.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getDp(s1, s2 string) [][]int {
	n, m := len(s1), len(s2)
	var ans [][]int
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

var (
	myDpTable = make(map[int]map[int][][]int)
)

// LCS for sth.
func LCS(p1, p2, id1, id2 int) int {
	scenario := Scenario.GetInstance()
	itr, ok1 := myDpTable[p1]
	if ok1 {
		result, ok2 := itr[p2]
		if ok2 {
			return result[id1][id2]
		} else {
			dp := getDp(scenario.GetIDSequence(p1), scenario.GetIDSequence(p2))
			tmp := make(map[int][][]int)
			tmp[p2] = dp
			myDpTable[p1] = tmp
			return dp[id1][id2]
		}
	}
	myDpTable[p1] = make(map[int][][]int)
	dp := getDp(scenario.GetIDSequence(p1), scenario.GetIDSequence(p2))
	tmp := make(map[int][][]int)
	tmp[p2] = dp
	myDpTable[p1] = tmp
	return dp[id1][id2]
}

// DP for sth.
func DP(s1, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)

	dp[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
