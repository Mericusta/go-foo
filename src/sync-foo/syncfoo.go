package syncfoo

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v8"
)

func WaitGroupCallFunctionWillCaptureWhenWait() {
	f := func(index int) {
		time.Sleep(time.Second)
		fmt.Printf("%v\n", index)
	}
	w := sync.WaitGroup{}

	// wrong: will capture index
	w.Add(5000)
	for index := 0; index != 5000; index++ {
		go func() {
			f(index)
			w.Done()
		}()
	}
	w.Wait()

	// correct: will capture index
	w.Add(5000)
	for index := 0; index != 5000; index++ {
		go func(i int) {
			f(i) // use closure func argument i but not caputre outer side value index
			w.Done()
		}(index) // pass index to closure func
	}
	w.Wait()
}

type SpinLocker uint32

func (sl *SpinLocker) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *SpinLocker) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

var localValue int

func operationLocalValue() {
	localValue++
}

var spinLocker SpinLocker
var mutex sync.Mutex

// SpinLockerPerformanceOnLocalOperation 自旋锁在本地操作时的性能表现
func SpinLockerPerformanceOnLocalOperation(gCount int) int {
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			spinLocker.Lock()
			operationLocalValue()
			spinLocker.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
	return localValue
}

// MutexLockerPerformanceOnLocalOperation 互斥锁在本地操作时的性能表现
func MutexLockerPerformanceOnLocalOperation(gCount int) int {
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			mutex.Lock()
			operationLocalValue()
			mutex.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
	return localValue
}

var cache struct {
	value string
}

var redisCacheKey string = "go_foo_test_cache"
var redisCacheValue string = "Hello Spin Key"
var redisCtx context.Context = context.Background()
var redisClient *redis.Client = func() *redis.Client {
	client := redis.NewClient(func() *redis.Options {
		opt, err := redis.ParseURL("redis://:@192.168.2.203:6379/4")
		if err != nil {
			panic(err)
		}
		return opt
	}())
	_, err := client.Ping(redisCtx).Result()
	if err != nil {
		panic(err)
	}
	_, err = client.Set(redisCtx, redisCacheKey, redisCacheValue, time.Hour).Result()
	if err != nil {
		panic(err)
	}
	return client
}()

func loadCacheFromRedis() {
	cv, err := redisClient.Get(redisCtx, redisCacheKey).Result()
	if err != nil {
		panic(err)
	}
	cache.value = cv
}

// SpinLockerPerformanceOnLoadCacheFromRedis 自旋锁在从 redis 读取缓存时的性能表现
func SpinLockerPerformanceOnLoadCacheFromRedis(gCount int) string {
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			spinLocker.Lock()
			loadCacheFromRedis()
			spinLocker.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
	return cache.value
}

// MutexLockerPerformanceOnLoadCacheFromRedis 互斥锁在从 redis 读取缓存时的性能表现
func MutexLockerPerformanceOnLoadCacheFromRedis(gCount int) string {
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			mutex.Lock()
			loadCacheFromRedis()
			mutex.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
	return cache.value
}

func blockingGoroutine(d time.Duration) {
	time.Sleep(d)
}

// SpinLockerPerformanceOnBlockingGoroutine 自旋锁在协程调度时的性能表现
func SpinLockerPerformanceOnBlockingGoroutine(gCount int, d time.Duration) {
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			spinLocker.Lock()
			blockingGoroutine(d)
			spinLocker.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
}

// MutexLockerPerformanceOnBlockingGoroutine 互斥锁在协程调度时的性能表现
func MutexLockerPerformanceOnBlockingGoroutine(gCount int, d time.Duration) {
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			mutex.Lock()
			blockingGoroutine(d)
			mutex.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
}

var channel chan int = make(chan int)

func channelSender(d time.Duration, max int) {
	t := time.NewTicker(d)
	defer t.Stop()
	c := 0
	for range t.C {
		channel <- 1
		c++
		if c > max {
			return
		}
	}
}

func channelReceiver() {
	<-channel
}

func SpinLockerPerformanceOnChannelReceiver(gCount int, tickerDuration time.Duration, tickerMax int) {
	go channelSender(tickerDuration, tickerMax)
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			spinLocker.Lock()
			channelReceiver()
			spinLocker.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
}

func MutexLockerPerformanceOnChannelReceiver(gCount int, tickerDuration time.Duration, tickerMax int) {
	go channelSender(tickerDuration, tickerMax)
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			mutex.Lock()
			channelReceiver()
			mutex.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
}

func httpServer() {
	http.HandleFunc("/sync/locker/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Spin Locker"))
	})
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func httpClient() {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/sync/locker/foo", nil)
	if err != nil {
		panic(err)
	}
	req.Close = true
	c := http.Client{}
	_, err = c.Do(req)
	if err != nil {
		panic(err)
	}
}

func SpinLockerPerformanceOnHttpRequest(gCount int) {
	go httpServer()
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			spinLocker.Lock()
			httpClient()
			spinLocker.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
}

func MutexLockerPerformanceOnHttpRequest(gCount int) {
	go httpServer()
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			spinLocker.Lock()
			httpClient()
			spinLocker.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
}

// PerformanceOnLoadCacheFromRedis

// 82987129 MutexLocker
// 74011133
// 955930750 SpinLocker
// 962044650

// PerformanceOnHttpRequest

// 2800218000 MutexLocker
// 2764006800
// 3798665000 http server sleep 10 ms before response
// 12527694800 http server sleep 100 ms before response
// 2813123800 SpinLocker
// 2773097500
// 3904198800 http server sleep 10 ms before response
// 12429413300 http server sleep 100 ms before response

// // ---

// PerformanceOnLocalOperation

// 34072 MutexLocker
// 35663 SpinLocker

// PerformanceOnChannelReceiver

// 1395442600 MutexLocker
// 1012998400 SpinLocker

// PerformanceOnBlockingGoroutine

// 1570832700 MutexLocker
// 1116442600 SpinLocker
