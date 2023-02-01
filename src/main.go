package main

import (
	extFlag "MLCS_GO/src/common/flag"
	"flag"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"

	"MLCS_GO/src/simulator"

	"MLCS_GO/src/strategy"
	"MLCS_GO/src/strategy/astar"
	fastlcs "MLCS_GO/src/strategy/fast_lcs"
	quickdp "MLCS_GO/src/strategy/quick_dp"
	"MLCS_GO/src/strategy/top"

	"github.com/golang/glog"
)

var myStrategy = flag.String("strategy", "Astar", "input my_strategy")

func solve() {
	var factory strategy.Factory

	switch *myStrategy {
	case "Astar":
		factory = new(astar.Factory)
	case "Top":
		factory = new(top.Factory)
	case "FastLCS":
		factory = new(fastlcs.Factory)
	case "QuickDP":
		factory = new(quickdp.Factory)
	default:
		factory = new(astar.Factory)
	}

	simulator := &simulator.Simulator{Factory: factory}
	simulator.RunFrameWork()
}

func main() {
	extFlag.Parse()
	defer glog.Flush()

	// cpu profile
	cpuFile, _ := os.Create("./cpu.prof")
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	//  Program MainForm
	solve()


	// memory profile
	memoryFile, _ := os.Create("./memory.prof")
	pprof.WriteHeapProfile(memoryFile)
}


func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)


}