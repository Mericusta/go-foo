package redisfoo

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Mericusta/go-stp"
	"github.com/go-redis/redis/v8"
)

func connect(url, password string, DB int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       DB,
	})
	if rdb == nil {
		panic(rdb)
	}
	return rdb
}

func ping(url, password string, DB int) string {
	rdb := connect(url, password, DB)
	cmd := rdb.Ping(context.Background())
	if cmd == nil {
		panic(cmd)
	}
	result, err := cmd.Result()
	if err != nil {
		panic(err)
	}
	return result
}

func zaddFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	r, e := rdb.ZAdd(context.Background(), "LGR_S1_TEST", &redis.Z{
		Score: 122, Member: 122,
	}).Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}

func zrankFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	r, e := rdb.ZRank(context.Background(), "liyunpeng_dev_local_LGR_S6_G1", "13472").Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}

func zrevrankFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	r, e := rdb.ZRevRank(context.Background(), "liyunpeng_dev_local_LGR_S6_G1", "13488").Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}

func getFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	r, e := rdb.Get(context.Background(), "NON_EXISTS_KEY").Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}

func hsetFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	m := make(map[string]interface{})
	im := make(map[int]int)
	for index := 1; index != 10; index++ {
		im[index] = index
	}
	for k, v := range im {
		m[fmt.Sprintf("field%v", k)] = fmt.Sprintf("value%v", v)
	}
	r, e := rdb.HSet(context.Background(), "HSET_TEST", m).Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}

func hgetallFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	m := make(map[string]interface{})
	m["field"] = "value"
	r, e := rdb.HGetAll(context.Background(), "HSET_TEST").Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}

var (
	lockScript = `
		redis.replicate_commands()
		local lockerKey = KEYS[1]
		local timestamp = ARGV[1]
		local expireSeconds = ARGV[2]

		local locked = redis.call('EXISTS', lockerKey)
		if locked == 1 then
			return 1
		end

		local setResultTable = redis.call('SET', lockerKey, timestamp)
		if setResultTable['ok'] ~= 'OK' then
			return 2
		end

		local expireResult = redis.call('EXPIRE', lockerKey, expireSeconds)
		if expireResult ~= 1 then
			return 3
		end

		return 0
	`
	unlockScript = `
		redis.replicate_commands()
		local lockerKey = KEYS[1]

		local locked = redis.call('EXISTS', lockerKey)
		if locked == 0 then
			return 1
		end

		local delResult = redis.call('DEL', lockerKey)
		if delResult ~= 1 then
			return 2
		end

		return 0
	`
)

func distributedLockerFoo(url, password string, DB int) {
	count := 10
	wg := &sync.WaitGroup{}
	wg.Add(count)
	for index := 0; index != count; index++ {
		go func(_index int, _wg *sync.WaitGroup) {
			defer _wg.Done()
			fmt.Printf("goroutine %v, running\n", _index)
			rdb := connect(url, password, DB)
			redisKey, nowTS, overtimeSeconds := "TEST_LOCKER", time.Now().Unix(), 3
			locker := stp.NewDistributedRedisLocker()

			result := locker.Lock(context.Background(), rdb, redisKey, overtimeSeconds, nowTS)
			switch result {
			case 0:
				fmt.Printf("goroutine %v, got locker, need unlock\n", _index)
			case 1:
				fmt.Printf("goroutine %v, other goroutine got locker, need waiting\n", _index)
				return
			case 2:
				fmt.Printf("goroutine %v, execute command 'set %v %v' failed\n", _index, redisKey, nowTS)
				return
			case 3:
				fmt.Printf("goroutine %v, execute command 'expire %v %v' failed\n", _index, redisKey, 10)
				return
			default:
				fmt.Printf("goroutine %v, execute lock script got unexpected value %v\n", _index, result)
				return
			}

			// do something
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))

			result = locker.Unlock(context.Background(), rdb, redisKey)
			switch result {
			case 0:
				fmt.Printf("goroutine %v, release locker\n", _index)
			case 1:
				fmt.Printf("goroutine %v, locker already released\n", _index)
				return
			case 2:
				fmt.Printf("goroutine %v, execute command 'del %v %v' failed\n", _index, redisKey, nowTS)
				return
			default:
				fmt.Printf("goroutine %v, execute unlock script got unexpected value %v\n", _index, result)
				return
			}
		}(index, wg)
	}
	wg.Wait()
}

type rogueData struct {
	OptionEventList []*optionEvent `json:"optionEventList"`
}

type optionEvent struct {
	EventID  int     `json:"eventID"`
	TodoList []*todo `json:"todoList"`
}

type todo struct {
	TodoID int   `json:"todoID"`
	NextID []int `json:"nextID"`
}

func SearchAndFixFromRedis(url, password string, DB int) {
	fmt.Println("---------------- search from redis ----------------")
	fmt.Println()

	rdb := connect(url, password, DB)
	keys, err := rdb.Keys(context.Background(), "*_ROGUE_data").Result()
	if err != nil {
		panic(err)
	}
	for _, key := range keys {
		value, err := rdb.Get(context.Background(), key).Result()
		if err != nil {
			panic(err)
		}
		// fmt.Printf("key = %v\n, value = %v\n", key, value)
		rd := &rogueData{}
		err = json.Unmarshal([]byte(value), rd)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("key = %v\n", key)
		for _, oe := range rd.OptionEventList {
			// fmt.Printf("eventID %v\n", oe.EventID)
			todoArray := stp.NewArray(oe.TodoList)
			for _, t := range oe.TodoList {
				// fmt.Printf("todoID %v, nextID %v\n", t.TodoID, t.NextID)
				for _, nextID := range t.NextID {
					if nextID == -1 {
						continue
					}
					index := todoArray.FindIndex(func(v *todo, i int) bool {
						return v.TodoID == nextID
					})
					if index == -1 {
						fmt.Printf("wrong data, key %v\n", key)
						fmt.Println()
						goto NEXT
					}
				}
			}
		}
		// fmt.Printf("no problem data, key %v\n", key)
		// fmt.Println()
	NEXT:
	}
}
