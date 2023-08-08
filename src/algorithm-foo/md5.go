package algorithmfoo

import (
	"crypto/md5"
	"fmt"
)

func MD5UsageFoo() {
	basicSeed := "channel-env-version-count-random"
	for i := 0; i < 10; i++ {
		tokenSeed := fmt.Sprintf("%v-%v", basicSeed, i)
		fmt.Printf("tokenSeed = %v\n", tokenSeed)
		tokenByte := md5.Sum([]byte(tokenSeed))
		fmt.Printf("tokenByte = %v\n", tokenByte)
		fmt.Printf("tokenByte = %x\n", tokenByte)
	}
}
