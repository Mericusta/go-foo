package iofoo

import (
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"time"
)

func ConcurrencyWriteFileFoo(sameWriter bool) {
	getWriter := func() io.Writer {
		f, err := os.OpenFile("./temp_file", os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		return f
	}

	if sameWriter {
		writer := getWriter()
		for index := 0; index != 10; index++ {
			_i := index
			go WriteFileFoo(_i, writer)
		}
	} else {
		for index := 0; index != 10; index++ {
			_i := index
			go WriteFileFoo(_i, getWriter())
		}
	}

	t := time.NewTimer(time.Second * 5)
	<-t.C
}

func WriteFileFoo(writerIndex int, outputFile io.Writer) {
	for index := 0; index != 1000; index++ {
		b := []byte(fmt.Sprintf("writer %v output file content at index %v\n", writerIndex, index))
		outputFile.Write(b)
	}
	fmt.Printf("writer %v done\n", writerIndex)
}

func WriteFileFooWithBuffer(writerIndex int, outputFile io.Writer) {
	buffer := bufio.NewWriterSize(outputFile, 1<<24)
	flushCount := 0

	for index := 0; index != 10000; index++ {
		b := []byte(fmt.Sprintf("writer %v output file content at index %v\n", writerIndex, index))
		if len(b) > buffer.Available() {
			flushCount++
			buffer.Flush()
			buffer.Reset(outputFile)
		}
		buffer.Write(b)
	}
	flushCount++
	buffer.Flush()
}

func tlvSocketPacketFoo(len int, value uint32) []byte {
	packetTagArray := make([]byte, len)
	binary.BigEndian.PutUint32(packetTagArray, value)
	fmt.Printf("packetTagArray = %v\n", packetTagArray)
	return packetTagArray
}

func ConcurrencyControlConsoleFoo() {
	ctx, canceler := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// also can handle signal in scan
	commandInputGoroutine := func() {
		defer func() {
			fmt.Println("stop command input goroutine")
			fmt.Println("call canceler")
			canceler()
			wg.Done()
		}()
		fmt.Printf("wait input: ")
		var (
			ok  bool
			err error
			in  *bufio.Scanner = bufio.NewScanner(os.Stdin)
		)
		for ok, err = in.Scan(), in.Err(); ok && err == nil; ok, err = in.Scan(), in.Err() {
			text := in.Text()
			fmt.Println("input:", in.Text())
			if text == "exit" {
				fmt.Println("command input goroutine exiting")
				break
			}
			fmt.Printf("wait input: ")
		}
		if !ok {
			fmt.Println("command input goroutine receive signal os.Interrupt, end scan")
		}
		if err != nil {
			fmt.Println("command input goroutine scan occurs error", err.Error())
		}
	}

	signalInputGoroutine := func() {
		defer func() {
			fmt.Println("stop signal input goroutine")
			wg.Done()
		}()
		s := make(chan os.Signal, 10)
		signal.Notify(s, os.Interrupt)
		select {
		case <-s:
			fmt.Println("signal input goroutine receive single os.Interrupt")
			break
		case <-ctx.Done():
			fmt.Println("signal input goroutine receive context cancel")
			break
		}
		signal.Stop(s)
		close(s)
	}

	go commandInputGoroutine()
	go signalInputGoroutine()

	defer func() {
		wg.Wait()
		fmt.Println("program exit")
		ticker := time.NewTicker(time.Second)
		second := 0
		for range ticker.C {
			second++
			fmt.Println("exit waiting second", second)
			if second >= 3 {
				break
			}
		}
	}()
}
