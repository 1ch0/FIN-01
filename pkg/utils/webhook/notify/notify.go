package notify

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func SendMsg(webHook, msg string) error {
	content := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`

	client := resty.New().SetTimeout(5 * time.Second).SetDisableWarn(true)
	_, err := client.R().SetHeaders(map[string]string{"accept": "application/json", "Content-Type": "application/json"}).
		SetBody(content).
		Post(webHook)
	if err != nil {
		return err
	}
	return nil
}
