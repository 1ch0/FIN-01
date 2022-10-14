package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

func main() {
	SendMessage("通知： 应用发布")
}

type Message struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func SendMessage(msg string) {
	var m Message
	m.MsgType = "text"
	m.Text.Content = msg
	jsons, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("SendMessage Marshal failed", zap.Error(err))
		return
	}
	resp := string(jsons)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "your webhook url", strings.NewReader(resp))
	if err != nil {
		log.Fatalf("SendMessage http NewRequest failed", zap.Error(err))
		return
	}
	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		log.Fatalf("SendMessage client Do failed", zap.Error(err))
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("SendMessage ReadAll Body failed", zap.Error(err))
		return
	}
	log.Println("SendMessage success", zap.String("body", string(body)))
}
