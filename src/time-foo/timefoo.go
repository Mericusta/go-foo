package timefoo

import (
	"fmt"
	"time"
)

func tickerAndSleep(count int) {
	counter := 0
	ticker := time.NewTicker(time.Second)
	time.Sleep(time.Second)
	for {
		fmt.Printf("for\n")
		select {
		case <-ticker.C:
			counter++
			fmt.Printf("counter++ = %v\n", counter)
			fmt.Printf("sleep 1s\n")
			time.Sleep(time.Second)
			fmt.Printf("sleep done\n")
			if counter == count {
				ticker.Stop()
				fmt.Printf("count done, ticker stop\n")
				return
			}
			fmt.Printf("ticker done\n")
		}
	}
}
