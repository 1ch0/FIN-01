package main

import (
	"fmt"
	"strings"

	"github.com/Lofanmi/pinyin-golang/pinyin"
)

func main() {
	dict := pinyin.NewDict()
	a := "1-2我，何时能暴富？"
	s = strings.Split(a, "-")
	s := dict.Convert(`我，何时能暴富？`, "").None()
	fmt.Println(s)
}
