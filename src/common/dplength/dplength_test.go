package dplength

import (
	"fmt"
	"testing"
)

func TestDplength(t *testing.T) {
	/*
		ACTAGCTA
		TCAGGTAT
	*/
	ans := LCS("TGC", "CGAT")
	fmt.Println(ans)
}
