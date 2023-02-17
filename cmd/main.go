package main

import gcfoo "go-foo/src/gc-foo"

// main 这是 main 函数注释的第一行
// main 这是 main 函数注释的第二行
func main() {
	gcfoo.AvoidGCScanByUintptr(400)
}
