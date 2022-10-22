package main

import (
	concurrencyfoo "go-foo/src/concurrency-foo"
)

func main() {
	concurrencyfoo.GoroutineCommunicateByBufferChannelWithLittleStructFoo(10)
}
