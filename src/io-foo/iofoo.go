package iofoo

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
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
