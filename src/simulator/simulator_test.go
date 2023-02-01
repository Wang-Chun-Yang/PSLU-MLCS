package simulator

import (
	astarFactory "MLCS_GO/src/strategy/astar/factory"
	"testing"
)

func TestSimulator(t *testing.T) {
	simulator := &Simulator{Factory: new(astarFactory.Factory)}
	simulator.Initialize()
	simulator.PreTreat()
	simulator.MainForm()
}
