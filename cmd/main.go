package main

import (
	concurrencyfoo "go-foo/src/concurrency-foo"
)

// main 这是 main 函数注释的第一行
// main 这是 main 函数注释的第二行
func main() { concurrencyfoo.GoroutineCommunicateByBufferChannelWithLittleStructFoo(10) }
