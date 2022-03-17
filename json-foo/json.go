package jsonfoo

import (
	"encoding/json"
	"fmt"
	mysqlfoo "go-foo/mysql-foo"
)

type JsonData struct {
	BirthTimestamp      int64 `json:"birth_timestamp,omitempty"`
	LastLoginTimestamp  int64 `json:"last_login_timestamp,omitempty"`
	DailyOnlineSeconds  int64 `json:"daily_online_seconds,omitempty"`
	WeeklyOnlineSeconds int64 `json:"weekly_online_seconds,omitempty"`
}

func JsonFoo() {
	token := "json-foo"
	channel := "json-foo"

	MYSQL_GAME_URI := `root:yunpeng.li@VanePlus950605@tcp(192.168.2.203:3306)/sg_gamedb_dev_anti_addiction`

	db, err := mysqlfoo.OpenMySQLDatabase(MYSQL_GAME_URI)
	if db == nil || err != nil {
		panic(err)
	}

	QUERY_ANTI_ADDICTION_DATA_SQL := `SELECT anti_addiction_data FROM account WHERE token = ? AND channel = ?`
	rows, err := db.Query(QUERY_ANTI_ADDICTION_DATA_SQL, token, channel)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		panic(err)
	}
	var jsonDataString string
	if rows.Next() {
		scanError := rows.Scan(&jsonDataString)
		if scanError != nil {
			panic(scanError)
		}
		fmt.Printf("jsonDataString = %v\n", jsonDataString)
		jsonData := &JsonData{}
		unmarshalErr := json.Unmarshal([]byte(jsonDataString), jsonData)
		if unmarshalErr != nil {
			panic(unmarshalErr)
		}
		fmt.Printf("jsonData = %+v\n", jsonData)
	} else {
		INSERT_ANTI_ADDICTION_DATA_SQL := `INSERT account (token, channel, anti_addiction_data) VALUES (?, ?, ?)`
		antiAddictionData := &JsonData{
			BirthTimestamp:      795405600, // 1995.03.17 10:00
			LastLoginTimestamp:  795405900, // 1995.03.17 10:05
			DailyOnlineSeconds:  0,
			WeeklyOnlineSeconds: 0,
		}
		antiAddictionJsonData, marshalErr := json.Marshal(antiAddictionData)
		if marshalErr != nil {
			panic(marshalErr)
		}
		_, execErr := db.Exec(INSERT_ANTI_ADDICTION_DATA_SQL, token, channel, string(antiAddictionJsonData))
		if execErr != nil {
			panic(execErr)
		}
	}
	db.Close()
}
