package netfoo

import (
	"context"
	"encoding/binary"
	"encoding/json"
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
		listener, err := net.Listen("tcp", "127.0.0.1:6677")
		if err != nil {
			panic(err)
		}
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
				err := connection.Close() // 本端关闭 connector
				fmt.Printf("server: connection close type 1, error: %v\n", err)
			} else {
				err := connection.Close() // 远端关闭 connector 之后本地也需要关闭 connector
				fmt.Printf("server: connection close type not 1, error: %v\n", err)
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

		// 远端关闭 connector 之后，本地也需要关闭
		// 本地关闭 connector 之后再次关闭会 use of closed network connection
		err = c.Close()
		if err != nil {
			if opError, ok := err.(*net.OpError); ok && opError.Err == net.ErrClosed {
				fmt.Printf("client: multi-close local connection\n")
			} else {
				fmt.Printf("client: close local connection occurs error: %v\n", err)
			}
		} else {
			fmt.Printf("client: close local connection\n")
		}
		wg.Done()
	}(ctx)

	time.Sleep(time.Second * 3)
	fmt.Printf("main: cancel\n")
	canceler()
	time.Sleep(time.Second * 3)
	wg.Wait()
}

// 关闭后又 connect 时的表现
func CloseAndReconnectFoo(reconnectCount int) {
	ctx, canceler := context.WithCancel(context.Background())

	// server
	go func(ctx context.Context) {
		linkMap := make(map[net.Conn]struct{})
		listener, _ := net.Listen("tcp", "127.0.0.1:6666")
		for {
			connection, acceptError := listener.Accept()
			if acceptError != nil {
				if acceptError.(*net.OpError).Err == net.ErrClosed {
					return
				}
				continue
			}

			fmt.Printf("link: new link [%v -> %v] is accepted\n", connection.RemoteAddr().String(), connection.LocalAddr().String())
			linkMap[connection] = struct{}{}

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
						return
					}
					fmt.Printf("link: read packets %v\n", packets)
				}
			}(connection)
		}
	}(ctx)

	time.Sleep(time.Second)

	// client
	go func(ctx context.Context) {
		var c net.Conn
		var err error
		connectCount := 0

	CONNECT:
		c, err = net.Dial("tcp", "127.0.0.1:6666")
		if err != nil {
			panic(fmt.Sprintf("dial occurs error: %v", err))
		}
		packet := make([]byte, 4)
		binary.BigEndian.PutUint32(packet, 1024)
		_, err = c.Write(packet)
		if err != nil {
			panic(fmt.Sprintf("write occurs error: %v", err))
		}

		time.Sleep(time.Second)
		fmt.Printf("client: connection closed by client\n")
		c.Close()

		if connectCount < reconnectCount {
			fmt.Printf("client: reconnect")
			connectCount++
			goto CONNECT
		}

		<-ctx.Done()
	}(ctx)

	time.Sleep(time.Second * 10)
	fmt.Printf("main: cancel\n")
	canceler()
	time.Sleep(time.Second * 10)
}

// ----------------------------------------------------------------

// ┌─────┬────────┬───────┐
// │ Tag │ Length │ Value │
// ├─────┼────────┼───────┤
// │  4  │   4    │       │
// └─────┴────────┴───────┘

const (
	// TLV 格式数据包中数据的标识的值的占位长度
	TLVPacketDataTagSize = 4

	// TLV 格式数据包中数据的长度的值的占位长度
	TLVPacketDataLengthSize = 4
)

func Pack[T any](connection net.Conn, msgID uint32, msgData any) error {
	msgValueByte, err := json.Marshal(msgData)
	if len(msgValueByte) == 0 {
		return fmt.Errorf("marshal msg %v %v got empty slice", msgID, msgData)
	}
	if err != nil {
		return err
	}

	msgByteDataLength := len(msgValueByte)
	tlvPacketLength := TLVPacketDataTagSize + TLVPacketDataLengthSize + msgByteDataLength
	tlvPacket := make([]byte, tlvPacketLength)

	// tlvPackMsg[0,TLVPacketDataTagSize]
	binary.BigEndian.PutUint32(tlvPacket, uint32(msgID))

	// tlvPackMsg[TLVPacketDataTagSize,TLVPacketDataTagSize+TLVPacketDataLengthSize]
	binary.BigEndian.PutUint32(tlvPacket[TLVPacketDataTagSize:], uint32(msgByteDataLength))

	// tlvPackMsg[TLVPacketDataTagSize+TLVPacketDataLengthSize:]
	copy(tlvPacket[TLVPacketDataTagSize+TLVPacketDataLengthSize:], msgValueByte)

	writeLength, writeError := connection.Write(tlvPacket)
	if writeError != nil {
		return writeError
	} else if writeLength != tlvPacketLength {
		return fmt.Errorf("write msg %v %v length %v not equal packet length %v", msgID, msgData, writeLength, msgByteDataLength)
	}

	return nil
}

func Unpack[T any](connection net.Conn) (uint32, any, error) {
	tagBytes := make([]byte, TLVPacketDataTagSize)
	_, readTagError := connection.Read(tagBytes)
	if readTagError != nil {
		return 0, nil, readTagError
	}
	tag := binary.BigEndian.Uint32(tagBytes)

	lengthBytes := make([]byte, TLVPacketDataLengthSize)
	_, readLengthError := connection.Read(lengthBytes)
	if readLengthError != nil {
		return 0, nil, readLengthError
	}
	length := binary.BigEndian.Uint32(lengthBytes)

	valueBytes := make([]byte, int(length))
	readValueLength, readValueError := connection.Read(valueBytes)
	if readValueError != nil {
		return 0, nil, readValueError
	} else if readValueLength != int(length) {
		return 0, nil, fmt.Errorf("read msg %v %v length %v not equal packet length %v", tag, valueBytes, readValueLength, length)
	}

	msg := new(T)
	err := json.Unmarshal(valueBytes, msg)
	if err != nil {
		return 0, nil, err
	}

	return tag, msg, nil
}

type user struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func tlvFoo() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	t := time.Now().Unix()
	u := &user{ID: 1, Name: "TLV-FOO"}
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			cancel()
			wg.Done()
		}()
		listener, err := net.Listen("tcp", "127.0.0.1:6666")
		if err != nil {
			panic(err)
		}
		connection, connectError := listener.Accept()
		if connectError != nil {
			panic(connectError)
		}
		_t, _uAny, err := Unpack[user](connection)
		if err != nil {
			panic(err)
		}
		_u, ok := _uAny.(*user)
		if !ok {
			panic(_uAny)
		}
		if _t != uint32(t) || _u.ID != u.ID || _u.Name != u.Name {
			panic("not equal")
		}
	}()

	go func() {
		defer wg.Done()
		connection, err := net.Dial("tcp", "127.0.0.1:6666")
		if err != nil {
			panic(err)
		}
		err = Pack[user](connection, uint32(t), u)
		if err != nil {
			panic(err)
		}
		<-ctx.Done()
	}()

	wg.Wait()
}

// ----------------------------------------------------------------

type socketBuffer struct {
	b []byte
	s int
}

func NewSocketBuffer(s int) *socketBuffer {
	return &socketBuffer{b: make([]byte, s), s: s}
}

func (sb *socketBuffer) Get(l int) []byte {
	if sb.s < l {
		return make([]byte, l)
	}
	return sb.b[:l]
}

var socketBufferPool sync.Pool = sync.Pool{
	New: func() any {
		return NewSocketBuffer(1024)
	},
}

func BufferPack[T any](connection net.Conn, buffer *socketBuffer, msgID uint32, msgData any) error {
	msgValueByte, err := json.Marshal(msgData)
	if len(msgValueByte) == 0 {
		return fmt.Errorf("marshal msg %v %v got empty slice", msgID, msgData)
	}
	if err != nil {
		return err
	}

	msgByteDataLength := len(msgValueByte)
	tlvPacketLength := TLVPacketDataTagSize + TLVPacketDataLengthSize + msgByteDataLength
	// tlvPacket := make([]byte, tlvPacketLength)
	tlvPacket := buffer.Get(tlvPacketLength)

	// tlvPackMsg[0,TLVPacketDataTagSize]
	binary.BigEndian.PutUint32(tlvPacket, uint32(msgID))

	// tlvPackMsg[TLVPacketDataTagSize,TLVPacketDataTagSize+TLVPacketDataLengthSize]
	binary.BigEndian.PutUint32(tlvPacket[TLVPacketDataTagSize:], uint32(msgByteDataLength))

	// tlvPackMsg[TLVPacketDataTagSize+TLVPacketDataLengthSize:]
	copy(tlvPacket[TLVPacketDataTagSize+TLVPacketDataLengthSize:], msgValueByte)

	writeLength, writeError := connection.Write(tlvPacket)
	if writeError != nil {
		return writeError
	} else if writeLength != tlvPacketLength {
		return fmt.Errorf("write msg %v %v length %v not equal packet length %v", msgID, msgData, writeLength, msgByteDataLength)
	}

	return nil
}

func BufferUnpack[T any](connection net.Conn, buffer *socketBuffer) (uint32, any, error) {
	// tagBytes := make([]byte, TLVPacketDataTagSize)
	tagBytes := buffer.Get(TLVPacketDataTagSize)
	_, readTagError := connection.Read(tagBytes)
	if readTagError != nil && readTagError != io.EOF {
		return 0, nil, readTagError
	}
	tag := binary.BigEndian.Uint32(tagBytes)

	// lengthBytes := make([]byte, TLVPacketDataLengthSize)
	lengthBytes := buffer.Get(TLVPacketDataLengthSize)
	_, readLengthError := connection.Read(lengthBytes)
	if readLengthError != nil {
		return 0, nil, readLengthError
	}
	length := binary.BigEndian.Uint32(lengthBytes)

	// valueBytes := make([]byte, int(length))
	valueBytes := buffer.Get(int(length))
	readValueLength, readValueError := connection.Read(valueBytes)
	if readValueError != nil {
		return 0, nil, readValueError
	} else if readValueLength != int(length) {
		return 0, nil, fmt.Errorf("read msg %v %v length %v not equal packet length %v", tag, valueBytes, readValueLength, length)
	}

	msg := new(T)
	err := json.Unmarshal(valueBytes, msg)
	if err != nil {
		return 0, nil, err
	}

	return tag, msg, nil
}

func tlvNetBufferFoo() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	t := time.Now().Unix()
	u := &user{ID: 1, Name: "TLV-FOO"}
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			cancel()
			wg.Done()
		}()
		listener, err := net.Listen("tcp", "127.0.0.1:6666")
		if err != nil {
			panic(err)
		}
		connection, connectError := listener.Accept()
		if connectError != nil {
			panic(connectError)
		}
		b := socketBufferPool.Get().(*socketBuffer)
		_t, _uAny, err := BufferUnpack[user](connection, b)
		if err != nil {
			panic(err)
		}
		_u, ok := _uAny.(*user)
		if !ok {
			panic(_uAny)
		}
		if _t != uint32(t) || _u.ID != u.ID || _u.Name != u.Name {
			panic("not equal")
		}
	}()

	go func() {
		defer wg.Done()
		connection, err := net.Dial("tcp", "127.0.0.1:6666")
		if err != nil {
			panic(err)
		}
		b := socketBufferPool.Get().(*socketBuffer)
		err = BufferPack[user](connection, b, uint32(t), u)
		if err != nil {
			panic(err)
		}
		<-ctx.Done()
	}()

	wg.Wait()
}
