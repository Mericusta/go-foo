package mysqlfoo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_MAX_OPEN_CONNECTIONS = 1000
	MYSQL_MAX_IDLE_CONNECTIONS = 100 // 并发协程数 = 保活链接数 / 10
)

func OpenMySQLDatabase(URI string) *sql.DB {
	db, openError := sql.Open("mysql", URI)
	if db == nil || openError != nil {
		panic(openError)
	}

	db.SetMaxOpenConns(MYSQL_MAX_OPEN_CONNECTIONS)
	db.SetMaxIdleConns(MYSQL_MAX_IDLE_CONNECTIONS)

	if pingError := db.Ping(); pingError != nil {
		panic(pingError)
	}

	return db
}

type prayResult struct {
	r map[int32]int64
}

func (p *prayResult) GetItem() map[int32]int64 {
	return p.r
}

func BatchInsertPrayRecordData() {
	db := OpenMySQLDatabase("root:yunpeng.li@VanePlus950605@tcp(192.168.2.203:3306)/sg_centerdb")

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
