package strategy_pattern

import "testing"

func TestStrategy(t *testing.T) {
	oprator := Operator{}

	oprator.setStrategy(&add{})
	result := oprator.caculate(1, 2)
	t.Log(result)

	oprator.setStrategy(&reduce{})
	result = oprator.caculate(2, 1)
	t.Log(result)
}
