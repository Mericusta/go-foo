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

func zoneDifference() {
	nowUnix := time.Now().Unix()
	nowLocalUnix := time.Now().Local().Unix()
	fmt.Println("nowUnix", nowUnix, "nowLocalUnix", nowLocalUnix, "equal", nowUnix == nowLocalUnix)
	nowUTCStr := time.Now().UTC().String()
	nowLocalStr := time.Now().Local().String()
	fmt.Println("nowUTCStr", nowUTCStr)
	fmt.Println("nowLocalStr", nowLocalStr)
}
