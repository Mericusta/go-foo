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
	"github.com/tidwall/spinlock"
	"golang.org/x/sync/singleflight"
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

type SpinLocker struct {
	_    sync.Mutex
	lock uintptr
}

func (sl *SpinLocker) Lock() {
loop:
	for !atomic.CompareAndSwapUintptr(&sl.lock, 0, 1) {
		runtime.Gosched()
		goto loop
	}
}

func (sl *SpinLocker) Unlock() {
	atomic.StoreUintptr(&sl.lock, 0)
}

var localValue int

func operationLocalValue() {
	localValue++
}

var spinLocker SpinLocker
var mutex sync.Mutex
var tidwallSpinLocker spinlock.Locker

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

func TidwallSpinLockerPerformanceOnLocalOperation(gCount int) int {
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			tidwallSpinLocker.Lock()
			operationLocalValue()
			tidwallSpinLocker.Unlock()
			gp.Done()
		}()
	}
	gp.Wait()
	return localValue
}

var cache struct {
	value       string
	holderCount int32
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
func SpinLockerPerformanceOnLoadCacheFromRedis(gCount int) (string, int32) {
	cache = struct {
		value       string
		holderCount int32
	}{}
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			spinLocker.Lock()
			if len(cache.value) != 0 {
				spinLocker.Unlock()
				goto USE_CACHE
			}
			loadCacheFromRedis()
			spinLocker.Unlock()
		USE_CACHE:
			atomic.AddInt32(&cache.holderCount, 1)
			gp.Done()
		}()
	}
	gp.Wait()
	return cache.value, cache.holderCount
}

func TidwallSpinLockerPerformanceOnLoadCacheFromRedis(gCount int) (string, int32) {
	cache = struct {
		value       string
		holderCount int32
	}{}
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			tidwallSpinLocker.Lock()
			if len(cache.value) != 0 {
				tidwallSpinLocker.Unlock()
				goto USE_CACHE
			}
			loadCacheFromRedis()
			tidwallSpinLocker.Unlock()
		USE_CACHE:
			atomic.AddInt32(&cache.holderCount, 1)
			gp.Done()
		}()
	}
	gp.Wait()
	return cache.value, cache.holderCount
}

// MutexLockerPerformanceOnLoadCacheFromRedis 互斥锁在从 redis 读取缓存时的性能表现
func MutexLockerPerformanceOnLoadCacheFromRedis(gCount int) (string, int32) {
	cache = struct {
		value       string
		holderCount int32
	}{}
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			if len(cache.value) != 0 {
				goto USE_CACHE
			}
			mutex.Lock()
			if len(cache.value) != 0 {
				mutex.Unlock()
				goto USE_CACHE
			}
			loadCacheFromRedis()
			mutex.Unlock()
		USE_CACHE:
			atomic.AddInt32(&cache.holderCount, 1)
			gp.Done()
		}()
	}
	gp.Wait()
	return cache.value, cache.holderCount
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

func TidwallSpinLockerPerformanceOnBlockingGoroutine(gCount int, d time.Duration) {
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			tidwallSpinLocker.Lock()
			blockingGoroutine(d)
			tidwallSpinLocker.Unlock()
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

func TidwallSpinLockerPerformanceOnChannelReceiver(gCount int, tickerDuration time.Duration, tickerMax int) {
	go channelSender(tickerDuration, tickerMax)
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			tidwallSpinLocker.Lock()
			channelReceiver()
			tidwallSpinLocker.Unlock()
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

func TidwallSpinLockerPerformanceOnHttpRequest(gCount int) {
	go httpServer()
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != 100; index++ {
		go func() {
			tidwallSpinLocker.Lock()
			httpClient()
			tidwallSpinLocker.Unlock()
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

func getValueFromRedisByKey(k string) string {
	v, err := redisClient.Get(redisCtx, k).Result()
	if err != nil {
		panic(err)
	}
	return v
}

var gsf singleflight.Group

func SingleFlightPerformanceOnLoadCacheFromRedis(gCount int) (string, int32) {
	cache = struct {
		value       string
		holderCount int32
	}{}
	gp := sync.WaitGroup{}
	gp.Add(gCount)
	for index := 0; index != gCount; index++ {
		go func() {
			_, err, _ := gsf.Do(redisCacheKey, func() (interface{}, error) {
				cache.value = getValueFromRedisByKey(redisCacheKey)
				return cache.value, nil
			})
			if err != nil {
				panic(err)
			}
			atomic.AddInt32(&cache.holderCount, 1)
			gp.Done()
		}()
	}
	gp.Wait()
	return cache.value, cache.holderCount
}

// PerformanceOnLoadCacheFromRedis

// 1075014 SingleFlight
// 20387928 65535 g
// 30443000 10w g
// 1042071 MutexLocker
// 16209227 65535 g
// 23575100 10w g
// 3689323 SpinLocker
// 17550989 65535 g
// 26109480 10w g
// 3005724 tidwall spinlocker

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
