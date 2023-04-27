package httpfoo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
