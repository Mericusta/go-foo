package httpfoo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

func RequestExample(index int) {
	vMap := url.Values{}
	vMap.Add("order_id", "11")
	vMap.Add("sdk_token", "token_202304131605")
	vMap.Add("channel_uid", "ios_user_202304131605")
	vMap.Add("sdk_order", "order_202304121535")

	client := &http.Client{}
	url := "http://127.0.0.1:8083/pay/cb"
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Origin", "http://ios.appstore.com")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", responseBody)
}

func JustPost(d time.Duration, useResty bool) {
	var postHandler func()
	if useResty {
		postHandler = func() {
			client := resty.New()
			client.R().Post("http://127.0.0.1:8182/pay/cb/mock")
		}
	} else {
		postHandler = func() {
			client := &http.Client{}
			url := "http://127.0.0.1:8182/pay/cb"
			request, err := http.NewRequest("POST", url, nil)
			if err != nil {
				panic(err)
			}
			request.Header.Set("Origin", "http://ios.appstore.com")
			client.Do(request)
		}
	}

	if d == 0 {
		for {
			postHandler()
		}
	} else {
		t := time.NewTicker(d)
		for range t.C {
			postHandler()
		}
	}
}
