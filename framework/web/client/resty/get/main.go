package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit":   "20",
			"sort":    "name",
			"order":   "asc",
			"random":  strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/search_result")
	if err != nil {
	}
	// Request.SetQueryString method
	resp, err = client.R().
		SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/show_product")

	// 解析返回的内容，内容是json解析到struct
	resp, err = client.R().
		//SetResult(result).
		ForceContentType("application/json").
		Get("v2/alpine/mainfestes/latest")

	fmt.Println("  Body       :\n", resp)
}
