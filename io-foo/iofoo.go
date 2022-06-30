package iofoo

import (
	"bufio"
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
