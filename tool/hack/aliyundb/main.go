package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Reference: 阿里云数据库外网地址查询工具,by 6time
//参考:https://forum.90sec.com/t/topic/1872

func main() {
	if len(os.Args) == 2 {
		name := os.Args[1]
		fmt.Println("[+]Start :", time.Now())
		dict2626(name)
		fmt.Println("[+]End  :", time.Now())
	} else {
		fmt.Println("example:\n\taliyundbburst.exe rr-1xx234x567x8xx999.mysql.rds.aliyuncs.com\n")
	}
	//name := "rr-2ze665c874z0wd537.mysql.rds.aliyuncs.com"
	//fmt.Println(strings.Replace(name, ".", "12.", 1))

}

func dict2626(www string) {
	//共计1296种组合
	const engstr = "abcdefghijklmnopqrstuvwxyz"
	const coustr = "1234567890"
	//26字母x26字母
	burstaa(www, engstr, engstr)
	//10数字x26字母
	burstaa(www, coustr, engstr)
	//26字母x10数字
	burstaa(www, engstr, coustr)
	//10数字x10数字
	burstaa(www, coustr, coustr)
}

func burstaa(www string, str1 string, str2 string) {
	for a1 := range str1 {
		for a2 := range str2 {
			//fmt.Println(string(engstr[a1]), string(engstr[a2]))
			new_www := strings.Replace(www, ".", string(str1[a1])+string(str2[a2])+".", 1)
			if GetwwwIP(new_www) {
				fmt.Println("[+]Find DB domain:", new_www)
				if NetWorkStatus(new_www) {
					fmt.Println("[+]End  :", time.Now())
					os.Exit(0)
				} else {
					fmt.Println("[!]Find DB next")
				}
			}
		}
	}
}

func GetwwwIP(name string) bool {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		//fmt.Println("Resolved err", err.Error())
		return false
	} else {
		fmt.Println("[+]Resolved address:", addr.String())
		return true
	}
}

func NetWorkStatus(www string) bool {
	cmd := exec.Command("ping", www)
	//output, _ := cmd.CombinedOutput()
	//fmt.Println("NetWorkStatus Start:", time.Now().Unix())
	err := cmd.Run()
	//fmt.Println("NetWorkStatus End  :", time.Now().Unix())
	if err != nil {
		fmt.Println("[!]Ping Status ERR", err.Error())
		return false
	} else {
		fmt.Println("[+]Ping Status OK")
		//fmt.Println(string(output))
	}
	return true
}
