package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/xmppo/go-xmpp"
)

var (
	xmppHost     string = "127.0.0.1"
	xmppPort     string = "5222"
	xmppPassword string = "1111"
)

func main() {
	var (
		err              error
		jid1             = "server1"
		jid2             = "server2"
		server1PublicJID = "server1_public@conference." + xmppHost
		server2PublicJID = "server2_public@conference." + xmppHost
		p1p2PrivateJID   = "p1p2_private@conference." + xmppHost
	)

	client1 := newXmppClient(jid1)
	client2 := newXmppClient(jid2)

	_, err = client1.JoinMUCNoHistory(server1PublicJID, "server1_delegate")
	if err != nil {
		panic(err)
	}
	_, err = client2.JoinMUCNoHistory(server2PublicJID, "server2_delegate")
	if err != nil {
		panic(err)
	}
	_, err = client1.JoinMUCNoHistory(p1p2PrivateJID, "s1p1_delegate")
	if err != nil {
		panic(err)
	}
	_, err = client2.JoinMUCNoHistory(p1p2PrivateJID, "s2p2_delegate")
	if err != nil {
		panic(err)
	}

	recvFunc := func(client *xmpp.Client) {
		fmt.Printf("jid %v, go recv\n", client.JID())
		for {
			chat, err := client.Recv()
			if err != nil {
				log.Fatal(err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				fmt.Printf("jid %v, xmpp.Chat\n", client.JID())
				fmt.Printf("jid %v, type %v\n", client.JID(), v.Type)
				fmt.Printf("jid %v, text %v\n", client.JID(), v.Text)
				fmt.Printf("jid %v, remote %v\n", client.JID(), v.Remote)
				if v.Type == "error" {
					fmt.Printf("jid %v, %+v\n", client.JID(), v)
				}
			case xmpp.Presence:
				fmt.Printf("jid %v xmpp.Presence\n", client.JID())
				fmt.Printf("from %v, show %v\n", v.From, v.Show)
			}
		}
	}

	go recvFunc(client1)
	go recvFunc(client2)

	time.Sleep(time.Second * 1)
	fmt.Println()
	time.Sleep(time.Second * 1)
	fmt.Println()

	_, err = client1.Send(xmpp.Chat{
		Remote: server1PublicJID,
		Type:   "groupchat",
		Text:   "client1_send_to_server1_public",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("client1 send\n")
	time.Sleep(time.Second)
	fmt.Println()

	client2.Send(xmpp.Chat{
		Remote: server2PublicJID,
		Type:   "groupchat",
		Text:   "client2_send_to_server2_public",
	})
	fmt.Printf("client2 send\n")
	time.Sleep(time.Second)
	fmt.Println()

	client1.Send(xmpp.Chat{
		Remote: p1p2PrivateJID + "/" + "s2p2_delegate",
		Type:   "chat",
		Text:   "client1_send_to_p1p2_private",
	})
	fmt.Printf("client1 send again\n")
	time.Sleep(time.Second)
	fmt.Println()

	client2.Send(xmpp.Chat{
		Remote: p1p2PrivateJID,
		Type:   "groupchat",
		Text:   "client2_send_to_p1p2_private",
	})
	fmt.Printf("client2 send again\n")
	time.Sleep(time.Second)
	fmt.Println()

	time.Sleep(time.Second)
}

func newXmppClient(jid string) *xmpp.Client {
	// 用户自动注册
	if err := newHttpClient().registerUser(jid, xmppPassword); err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%s:%s", xmppHost, xmppPort)
	fmt.Printf("NewXmppClient jid: %v, address: %v, JID: %v, pass: %v\n", jid, address, jid, xmppPassword)
	options := xmpp.Options{
		Host:     address,
		User:     fmt.Sprintf("%v@%s", jid, xmppHost),
		Password: xmppPassword,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NoTLS:       true,
		DialTimeout: 2 * time.Second,
		Mechanism:   "PLAIN",
	}

	xmppClient, err := options.NewClient()
	if err != nil {
		panic(err)
	}

	return xmppClient
}

type httpClient struct {
	BaseURL    string // eJabberd服务器的URL
	AuthToken  string // 认证token
	httpClient *http.Client
}

func newHttpClient() *httpClient {
	// 本地测试
	baseURL := fmt.Sprintf("http://%s:5443", xmppHost)

	adminUser := "admin@localhost"
	adminStr := fmt.Sprintf("%s:%s", adminUser, xmppPassword)
	encoded := base64.StdEncoding.EncodeToString([]byte(adminStr))
	token := fmt.Sprintf("Basic %s", encoded)

	return &httpClient{
		BaseURL:    baseURL,
		AuthToken:  token,
		httpClient: &http.Client{},
	}
}

// registerUser 注册用户
func (c *httpClient) registerUser(username string, password string) error {
	url := fmt.Sprintf("http://%s:5443/api/register", xmppHost)
	params := map[string]interface{}{
		"host":     xmppHost,
		"user":     username,
		"password": password,
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	_, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	return nil
}

type (
	Options struct {
		Username string            // 用户名
		Password string            // 密码
		Headers  map[string]string // 请求头
		Timeout  time.Duration     // 超时
		Client   *http.Client
	}

	Response struct {
		Errno  int                    `json:"errno"`
		ErrMsg string                 `json:"errmsg"`
		Data   map[string]interface{} `json:"data"`
	}
)

func newOptions() *Options {
	return &Options{}
}
