package concurrencyfoo

import "sync"

var (
	TCP_SOCKET_POOL_SERVER      = "127.0.0.1"
	TCP_SOCKET_POOL_SERVER_PORT = "6379"
)

func TcpSocketPoolFoo() {
	wg := sync.WaitGroup{}
	wg.Add(2)

}
