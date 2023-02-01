package fastlcs

import (
	"MLCS_GO/src/strategy"
)

// Factory 实现工厂方法创建Astar对象
type Factory struct{}

// CreateStrategy for sth.
func (ths *Factory) CreateStrategy() strategy.Strategy {
	return &FastLCS{}
}
