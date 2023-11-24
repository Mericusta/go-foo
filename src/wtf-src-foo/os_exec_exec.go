package wtfsrcfoo

import (
	"context"
	"fmt"
)

// What The Fuck In Go src

// release-branch.go1.21
// $GOROOT/src/os/exec/exec.go:666
// ----------------------------------------------------------------
// Cmd.Start
func WTFInCmdStart() {
	ctx, _ := context.WithTimeout(context.Background(), -1)
	select {
	case <-ctx.Done():
		fmt.Println("never done")
	default:
		fmt.Println("always default")
	}
}
