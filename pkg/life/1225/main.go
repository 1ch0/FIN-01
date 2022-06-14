package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	APPLE = "🍎"
	STAR  = "🌟"
	BELL  = "🔔"
	TREE  = "🎄"
	DOOR  = "🚪"
	GIFT  = "🎁"
)

func main() {
	//先来个5层高的
	ct := newChristmasTree(4)
	fmt.Print(ct)
}

type ChristmasTree struct {
	floor       int
	treeBuilder strings.Builder
}

func (ct *ChristmasTree) String() string {
	return ct.treeBuilder.String()
}

func newChristmasTree(floor int) *ChristmasTree {
	ct := &ChristmasTree{floor: floor}
	ct.gen()
	return ct
}

//获取第floor层的第line行的圣诞树数量
func getLineAmount(floor, line int) int {
	return 1 + line*2 + floor*4 + int(floor/2*2)*int((floor+1)/2)
}

//随机按比例分配星星、铃铛、苹果和圣诞树
func randAppleTree() string {
	r := rand.Intn(100)
	if r < 1 {
		return STAR
	} else if r < 2 {
		return BELL
	} else if r < 10 {
		return APPLE
	} else {
		return TREE
	}
}

//Blog：www.flysnow.org
//Wechat:flysnow_org
//生成一整颗圣诞树
func (ct *ChristmasTree) gen() {
	bottomAmount := getLineAmount(ct.floor, ct.floor+4)

	//一层，一行的生成
	for floor := 0; floor < ct.floor; floor++ {
		for line := 0; line < floor+5; line++ {
			lineAmount := getLineAmount(floor, line)

			for i := (bottomAmount-lineAmount)/2 - 1; i > 0; i-- {
				ct.treeBuilder.WriteString(" ")
			}

			for i := 0; i < lineAmount; i++ {
				ct.treeBuilder.WriteString(randAppleTree())
			}

			ct.treeBuilder.WriteString("\n")
		}
	}

	//居中、生成圣诞树根
	for i := 0; i < ct.floor; i++ {
		lineAmount := ct.floor + (ct.floor+1)%2 //一个更接近层数的近似值

		for i := (bottomAmount-lineAmount)/2 - 1; i > 0; i-- {
			ct.treeBuilder.WriteString(" ")
		}

		for i := 0; i < lineAmount; i++ {
			ct.treeBuilder.WriteString(DOOR)
		}
		ct.treeBuilder.WriteString("\n")
	}

	//在圣诞树下放点礼物
	if ct.floor > 3 {
		//主要是从倒数第2行开始放
		lines := strings.Split(ct.treeBuilder.String(), "\n")
		lines[len(lines)-4] += "    " + GIFT
		lines[len(lines)-3] += "   " + GIFT + GIFT + GIFT
		lines[len(lines)-2] += "  " + GIFT + GIFT + GIFT + GIFT + GIFT
		ct.treeBuilder.Reset()
		ct.treeBuilder.WriteString(strings.Join(lines, "\n"))
	}
}
