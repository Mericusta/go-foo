package iofoo

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

func WriteFileFoo(writerIndex int, outputFile io.Writer) {
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
