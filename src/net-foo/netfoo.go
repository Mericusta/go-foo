package netfoo

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 关闭 connector 时的表现
func CloseConnectorFoo(closedBy int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
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
				fmt.Printf("link: handle read\n")
				for {
					packets := make([]byte, 4)
					_, err := c.Read(packets)
					fmt.Printf("link: err %v\n", err)
					fmt.Printf("link: err == EOF %v\n", err == io.EOF)
					fmt.Printf("link: err == ErrClosed %v\n", func() bool {
						netError, ok := err.(*net.OpError)
						return ok && netError.Err == net.ErrClosed
					}())
					if err != nil {
						if err == io.EOF {
							fmt.Printf("link: connection closed by remote\n")
						} else if opError, ok := err.(*net.OpError); ok && opError.Err == net.ErrClosed {
							fmt.Printf("link: connection closed by local\n")
						} else {
							fmt.Printf("link: connection read occurs error: %v\n", err)
							continue
						}
						wg.Done()
						return
					}
					fmt.Printf("link: read packets %v\n", packets)
				}
			}(connection)

			<-ctx.Done()
			if closedBy == 1 {
				fmt.Printf("server: connection closed by server\n")
				connection.Close() // 本端关闭 connector
			}
			fmt.Printf("server: listener closed\n")
			listener.Close()
			wg.Wait()
		}
	}(ctx)

	time.Sleep(time.Second * 3)
	go func(ctx context.Context) {
		c, _ := net.Dial("tcp", "127.0.0.1:6666")
		if closedBy == 2 {
			fmt.Printf("client: connection closed by client\n")
			c.Close() // 本端关闭
		}
		<-ctx.Done()

		time.Sleep(time.Second * 3)
		fmt.Printf("client: after sleep 3s, connection write sth\n")
		packet := make([]byte, 4)
		binary.BigEndian.PutUint32(packet, 1024)
		n, err := c.Write(packet) // 远端关闭 connection 之后这边仍然可以写入成功？需要处理
		fmt.Printf("client: write n %v err = %v\n", n, err)
		if opError, ok := err.(*net.OpError); ok && opError.Err == net.ErrClosed {
			fmt.Printf("client: connection closed by local\n")
		}
		time.Sleep(time.Second * 3)
		n, err = c.Write(packet) // 远端关闭 connection 之后这边再次写入会失败？
		// Windows: wsasend: An established connection was aborted by the software in your host machine.
		// Linux: write: broken pipe
		fmt.Printf("client: write n %v err = %v\n", n, err)
		if opError, ok := err.(*net.OpError); ok && opError.Err == net.ErrClosed {
			fmt.Printf("client: connection closed by local\n")
		}
		wg.Done()
	}(ctx)

	time.Sleep(time.Second * 3)
	fmt.Printf("main: cancel\n")
	canceler()
	time.Sleep(time.Second * 3)
	wg.Wait()
}
