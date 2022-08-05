package strategy_pattern

// 可以随意更换策略，而不影响 Operator 的所有实现
type IStrategy interface {
	do(int, int) int
}

type add struct {
}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	startegy IStrategy
}

func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.startegy = strategy
}

func (operator *Operator) caculate(a, b int) int {
	return operator.startegy.do(a, b)
}
