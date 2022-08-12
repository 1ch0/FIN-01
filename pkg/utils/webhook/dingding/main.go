package main

import (
	"net/http"
	"strings"
)

func main() {
	SendDingMsg("test  2022年7月20日16:20:15")
}

func SendDingMsg(msg string) {
	//请求地址模板
	webHook := ``
	content := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		// handle error
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		// handle error
	}
}
