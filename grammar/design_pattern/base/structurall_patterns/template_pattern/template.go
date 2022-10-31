package template_pattern

import "fmt"

type Cooker interface {
	fire()
	cooke()
	outfire()
}

type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("fire")
}

func (CookMenu) cooke() {

}

func (CookMenu) outfire() {
	fmt.Println("out fire")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("xi hong shi")
}

func (*XiHongShi) outfire() {
	fmt.Println("关火")
}

type ChaoJiDan struct {
	CookMenu
}

func (ChaoJiDan) cooke() {
	fmt.Println("chao ji dan")
}
