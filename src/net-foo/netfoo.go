package netfoo

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 关闭 connector 时的表现
func CloseConnectorFoo(closedBy int) {
	ctx, canceler := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		listener, _ := net.Listen("tcp", "127.0.0.1:6666")
		for {
			connection, acceptError := listener.Accept()
			if acceptError != nil {
				if acceptError.(*net.OpError).Err == net.ErrClosed {
					return
				}
				continue
			}

			wg := sync.WaitGroup{}
			wg.Add(1)
			go func(c net.Conn) {
				fmt.Printf("handle read\n")
				for {
					packets := make([]byte, 4)
					_, err := c.Read(packets)
					fmt.Printf("err %v\n", err)
					fmt.Printf("err == EOF %v\n", err == io.EOF)
					fmt.Printf("err == ErrClosed %v\n", func() bool {
						netError, ok := err.(*net.OpError)
						return ok && netError.Err == net.ErrClosed
					}())
					if err != nil {
						if err == io.EOF {
							fmt.Printf("connection closed by remote\n")
						} else if opError, ok := err.(*net.OpError); ok && opError.Err == net.ErrClosed {
							fmt.Printf("connection closed by local\n")
						} else {
							fmt.Printf("connection read occurs error: %v\n", err)
							continue
						}
						wg.Done()
						return
					}
				}
			}(connection)

			<-ctx.Done()
			if closedBy == 1 {
				fmt.Printf("connection closed by server\n")
				connection.Close() // 本端关闭 connector
			}
			fmt.Printf("listener closed\n")
			listener.Close()
			wg.Wait()
		}
	}(ctx)

	time.Sleep(time.Second * 3)
	go func(ctx context.Context) {
		c, _ := net.Dial("tcp", "127.0.0.1:6666")
		if closedBy == 2 {
			fmt.Printf("connection closed by client\n")
			c.Close() // 对端关闭 connector
		}
		<-ctx.Done()
	}(ctx)

	time.Sleep(time.Second * 3)
	fmt.Printf("cancel\n")
	canceler()
	time.Sleep(time.Second * 3)
}
