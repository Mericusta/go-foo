package httpfoo

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/http2"
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

// http2 使用同一个 tcp链接 做异步请求 的表现
func http2AsyncRequestWithOneConnection() {
	http2ServerGoroutine := func() {
		// 加载TLS证书和私钥
		cert, err := tls.LoadX509KeyPair("http2_certificate.crt", "http2_private.key")
		if err != nil {
			log.Fatalf("server: loadkeys: %s", err)
		}

		// 创建TLS配置
		config := &tls.Config{
			Certificates: []tls.Certificate{cert},
			NextProtos:   []string{"h2", "http/1.1"}, // 支持HTTP/2
		}

		server := &http.Server{
			Addr:      ":https",
			TLSConfig: config,
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Hello, HTTP/2!"))
		})

		log.Println("Starting server on https://localhost")
		log.Fatal(server.ListenAndServeTLS("http2_certificate.crt", "http2_private.key"))
	}

	go http2ServerGoroutine()

	time.Sleep(time.Second)

	clientCount := 10
	wg := &sync.WaitGroup{}
	wg.Add(clientCount)

	// 加载自签名的CA证书
	caCert, err := ioutil.ReadFile("http2_certificate.crt")
	if err != nil {
		log.Fatalf("Error reading CA certificate: %v", err)
	}
	caCertBlock, _ := pem.Decode(caCert)
	if caCertBlock == nil || caCertBlock.Type != "CERTIFICATE" {
		log.Fatal("Failed to parse CA certificate PEM")
	}
	caCertPool := x509.NewCertPool()
	cert, err := x509.ParseCertificate(caCertBlock.Bytes)
	if err != nil {
		log.Fatalf("ParseCertificate failed: %v", err)
	}
	caCertPool.AddCert(cert)

	// 创建HTTP客户端，指定TLS配置
	client := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				// RootCAs: caCertPool, // 使用自定义的根证书池 -> 需要定向本机 host 解析，localhost -> 解析为 127.0.0.1
				InsecureSkipVerify: true,
			},
		},
	}

	for index := 0; index != clientCount; index++ {
		go func(_i int, _c *http.Client) {
			// 发送请求
			resp, err := client.Get("https://localhost")
			if err != nil {
				log.Fatalf("Error fetching: %v", err)
			}
			defer resp.Body.Close()

			// 读取响应体
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Error reading response: %v", err)
			}

			fmt.Printf("Response: %s\n", body)
			wg.Done()
		}(index, client)
	}

	wg.Wait()
}
