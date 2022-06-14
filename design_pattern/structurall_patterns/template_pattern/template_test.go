package template_pattern

import "testing"

func TestTemplate(t *testing.T) {
	x := &XiHongShi{}
	doCook(x)

	c := &ChaoJiDan{}
	doCook(c)
}
