package main

import (
	channelfoo "go-foo/channel-foo"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// channelfoo.GoRoutineExitThenCloseChannel()
	channelfoo.ListenerBlockedChannel()
}
