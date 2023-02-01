package strategy

// Strategy for sth.
type Strategy interface {
	Initialize()
	Execute()
}

// Factory for sth.
type Factory interface {
	CreateStrategy() Strategy
}
