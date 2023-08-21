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

func JustPost(url string, header map[string]string, d time.Duration, useResty, concurrency bool, concurrencyCount int) {
	var postHandler func()
	if useResty {
		postHandler = func() {
			client := resty.New()
			request := client.R()
			for k, v := range header {
				request.SetHeader(k, v)
			}
			request.Post(url)
		}
	} else {
		postHandler = func() {
			client := &http.Client{}
			request, err := http.NewRequest("POST", url, nil)
			if err != nil {
				panic(err)
			}
			for k, v := range header {
				request.Header.Set(k, v)
			}
			client.Do(request)
		}
	}

	var rangeHandler func(int)
	if d == 0 {
		rangeHandler = func(i int) {
			counter := 0
			for {
				counter++
				postHandler()
				fmt.Printf("index %v, counter = %v\n", i, counter)
			}
		}
	} else {
		rangeHandler = func(i int) {
			counter := 0
			t := time.NewTicker(d)
			for range t.C {
				counter++
				postHandler()
				fmt.Printf("index %v, counter = %v\n", i, counter)
			}
		}
	}

	if concurrency {
		for i := 0; i < concurrencyCount; i++ {
			_i := i
			go rangeHandler(_i)
		}
		select {}
	} else {
		rangeHandler(0)
	}

}
