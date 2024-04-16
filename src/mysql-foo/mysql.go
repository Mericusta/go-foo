package mysqlfoo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_MAX_OPEN_CONNECTIONS = 1000
	MYSQL_MAX_IDLE_CONNECTIONS = 100 // 并发协程数 = 保活链接数 / 10
	MYSQL_DATABASE_FOO         = "root:yunpeng.li@VanePlus950605@tcp(192.168.2.203:3306)/foo"
)

func connect(URI string) *sql.DB {
	db, openError := sql.Open("mysql", URI)
	if db == nil || openError != nil {
		panic(openError)
	}

	db.SetMaxOpenConns(MYSQL_MAX_OPEN_CONNECTIONS)
	db.SetMaxIdleConns(MYSQL_MAX_IDLE_CONNECTIONS)

	return db
}

func ping(url string) {
	db := connect(url)
	if pingError := db.Ping(); pingError != nil {
		panic(pingError)
	}
}

var INSERT_OR_UPDATE_SQL string = `
	INSERT INTO insert_or_update_table VALUES (?, ?, ?, ?) AS new(primary_key, new_field1, new_field2, new_condition1) ON DUPLICATE KEY UPDATE field1 = IF(condition1 < new.new_condition1, new.new_field1, field1), field2 = IF(condition1 < new.new_condition1, new.new_field2, field2), condition1 = IF(condition1 < new.new_condition1, new.new_condition1, condition1)
`

func insertOrUpdateFoo(url string) {
	db := connect(url)
	key := time.Now().Unix()
	for index := 0; index != 10; index++ {
		field1, field2, condition1 := index, -index, index
		result, err := db.Exec(INSERT_OR_UPDATE_SQL, key, field1, field2, condition1)
		if err != nil {
			panic(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			panic(err)
		}
		fmt.Printf("rowsAffected %v\n", rowsAffected)
	}
}

func concurrentlyInsertOrUpdateFoo(url string, goroutineCount int) {
	db := connect(url)
	wg := &sync.WaitGroup{}
	key := time.Now().Unix()
	wg.Add(goroutineCount)
	for g := 0; g != goroutineCount; g++ {
		go func(_g int, _wg *sync.WaitGroup) {
			defer _wg.Done()
			for index := 0; index != goroutineCount; index++ {
				field1, field2, condition1 := index, -index, index
				result, err := db.Exec(INSERT_OR_UPDATE_SQL, key, field1, field2, condition1)
				if err != nil {
					panic(err)
				}
				rowsAffected, err := result.RowsAffected()
				if err != nil {
					panic(err)
				}
				if rowsAffected > 0 {
					fmt.Printf("_g %v, rowsAffected %v\n", _g, rowsAffected)
				}
			}
		}(g, wg)
	}
	wg.Wait()
}

// --------------------------------

type prayResult struct {
	r map[int32]int64
}

func (p *prayResult) GetItem() map[int32]int64 {
	return p.r
}

func BatchInsertPrayRecordData() {
	db := connect("user@password@tcp(192.168.2.203:3306)/db")

	playerID, prayID, prayTimestamp := 1, 1, time.Now().Unix()
	prayResult := []*prayResult{
		{
			r: map[int32]int64{
				10: 1,
			},
		},
		{
			r: map[int32]int64{
				11: 1,
			},
		},
		{
			r: map[int32]int64{
				12: 1,
			},
		},
		{
			r: map[int32]int64{
				13: 1,
			},
		},
		{
			r: map[int32]int64{
				14: 1,
			},
		},
		{
			r: map[int32]int64{
				15: 1,
			},
		},
		{
			r: map[int32]int64{
				16: 1,
			},
		},
		{
			r: map[int32]int64{
				17: 1,
			},
		},
		{
			r: map[int32]int64{
				18: 1,
			},
		},
		{
			r: map[int32]int64{
				19: 1,
			},
		},
	}
	var INSERT_PRAY_RECORDS_SQL string = `
		INSERT INTO pray_records (player_id, pray_id, reward_content, reward_at) VALUES 
	`
	var INSERT_PRAY_VALUE_FORMAT string = `(%v, %v, '%v', %v)`

	sqlBuilder, valueBuilder := strings.Builder{}, strings.Builder{}
	sqlBuilder.WriteString(INSERT_PRAY_RECORDS_SQL)
	for index, prayReward := range prayResult {
		if index != 0 {
			sqlBuilder.WriteRune(',')
		}
		sqlBuilder.WriteRune(' ')
		if len(prayReward.GetItem()) == 0 {
			panic("prayReward is empty")
		}
		prayContent, marshalError := json.Marshal(prayReward.GetItem())
		if len(prayContent) == 0 || marshalError != nil {
			panic(marshalError)
		}
		_, writeError := valueBuilder.WriteString(fmt.Sprintf(INSERT_PRAY_VALUE_FORMAT, playerID, prayID, prayContent, prayTimestamp))
		if writeError != nil {
			panic(writeError)
		}
		sqlBuilder.WriteString(valueBuilder.String())
		valueBuilder.Reset()
	}

	_, execError := db.Exec(sqlBuilder.String())
	fmt.Printf("sqlBuilder.String() = %v\n", sqlBuilder.String())
	if execError != nil {
		panic(execError)
	}
	db.Close()
}

// --------------------------------

func queryNilFoo() {
	db := connect(MYSQL_DATABASE_FOO)
	if db == nil {
		panic(MYSQL_DATABASE_FOO)
	}

	r, e := db.Query("select * from insert_or_update_table where primary_key = -1")
	fmt.Println("r =", r)
	fmt.Println("e =", e)
}
