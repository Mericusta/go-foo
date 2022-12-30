package jsonfoo

import (
	"encoding/json"
	"fmt"
)

type JsonData struct {
	CurrentLoginTimestamp int64
	BirthTimestamp        int64 `json:"birth_timestamp,omitempty"`
	LastLoginTimestamp    int64 `json:"last_login_timestamp,omitempty"`
	DailyOnlineSeconds    int64 `json:"daily_online_seconds,omitempty"`
	WeeklyOnlineSeconds   int64 `json:"weekly_online_seconds,omitempty"`
}

func JsonFoo() {
	str := "123,asd"

	s := &JsonData{}
	err := json.Unmarshal([]byte(str), s)
	if err != nil {
		panic(err)
	}

	json.Marshal(s)

	// token := "json-foo"
	// channel := "json-foo"

	// MYSQL_GAME_URI := `root:yunpeng.li@VanePlus950605@tcp(192.168.2.203:3306)/sg_gamedb_dev_anti_addiction`

	// db, err := mysqlfoo.OpenMySQLDatabase(MYSQL_GAME_URI)
	// if db == nil || err != nil {
	// 	panic(err)
	// }

	// QUERY_ANTI_ADDICTION_DATA_SQL := `SELECT anti_addiction_data FROM account WHERE token = ? AND channel = ?`
	// rows, err := db.Query(QUERY_ANTI_ADDICTION_DATA_SQL, token, channel)
	// defer func() {
	// 	if rows != nil {
	// 		rows.Close()
	// 	}
	// }()
	// if err != nil {
	// 	panic(err)
	// }
	// var jsonDataString string
	// if rows.Next() {
	// 	scanError := rows.Scan(&jsonDataString)
	// 	if scanError != nil {
	// 		panic(scanError)
	// 	}
	// 	fmt.Printf("jsonDataString = %v\n", jsonDataString)
	// 	jsonData := &JsonData{}
	// 	unmarshalErr := json.Unmarshal([]byte(jsonDataString), jsonData)
	// 	if unmarshalErr != nil {
	// 		panic(unmarshalErr)
	// 	}
	// 	fmt.Printf("jsonData = %+v\n", jsonData)
	// } else {
	// 	INSERT_ANTI_ADDICTION_DATA_SQL := `INSERT account (token, channel, anti_addiction_data) VALUES (?, ?, ?)`
	// 	antiAddictionData := &JsonData{
	// 		BirthTimestamp:      795405600, // 1995.03.17 10:00
	// 		LastLoginTimestamp:  795405900, // 1995.03.17 10:05
	// 		DailyOnlineSeconds:  0,
	// 		WeeklyOnlineSeconds: 0,
	// 	}
	// 	antiAddictionJsonData, marshalErr := json.Marshal(antiAddictionData)
	// 	if marshalErr != nil {
	// 		panic(marshalErr)
	// 	}
	// 	_, execErr := db.Exec(INSERT_ANTI_ADDICTION_DATA_SQL, token, channel, string(antiAddictionJsonData))
	// 	if execErr != nil {
	// 		panic(execErr)
	// 	}
	// }
	// db.Close()
}

func InterfaceMarshalFoo() {
	s := &JsonData{
		CurrentLoginTimestamp: 1,
		BirthTimestamp:        2,
		LastLoginTimestamp:    3,
		DailyOnlineSeconds:    4,
		WeeklyOnlineSeconds:   5,
	}
	marshalStruct, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	i := func() interface{} {
		return &JsonData{
			CurrentLoginTimestamp: 1,
			BirthTimestamp:        2,
			LastLoginTimestamp:    3,
			DailyOnlineSeconds:    4,
			WeeklyOnlineSeconds:   5,
		}
	}()
	marshalInterface, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	if len(marshalStruct) != len(marshalInterface) {
		panic(fmt.Sprintf("len not equal %v %v", len(marshalStruct), len(marshalInterface)))
	}

	for index := 0; index != len(marshalStruct); index++ {
		if marshalStruct[index] != marshalInterface[index] {
			panic(fmt.Sprintf("index %v b not equal", index))
		}
	}

	unmarshalStruct := &JsonData{}
	err = json.Unmarshal(marshalStruct, unmarshalStruct)
	if err != nil {
		panic(err)
	}

	unmarshalInterface := &JsonData{}
	err = json.Unmarshal(marshalStruct, unmarshalInterface)
	if err != nil {
		panic(err)
	}

	switch {
	case unmarshalStruct.CurrentLoginTimestamp != unmarshalInterface.CurrentLoginTimestamp:
		panic(fmt.Sprintf("CurrentLoginTimestamp not equal %+v %+v", unmarshalStruct, unmarshalInterface))
	case unmarshalStruct.BirthTimestamp != unmarshalInterface.BirthTimestamp:
		panic(fmt.Sprintf("BirthTimestamp not equal %+v %+v", unmarshalStruct, unmarshalInterface))
	case unmarshalStruct.LastLoginTimestamp != unmarshalInterface.LastLoginTimestamp:
		panic(fmt.Sprintf("LastLoginTimestamp not equal %+v %+v", unmarshalStruct, unmarshalInterface))
	case unmarshalStruct.DailyOnlineSeconds != unmarshalInterface.DailyOnlineSeconds:
		panic(fmt.Sprintf("DailyOnlineSeconds not equal %+v %+v", unmarshalStruct, unmarshalInterface))
	case unmarshalStruct.WeeklyOnlineSeconds != unmarshalInterface.WeeklyOnlineSeconds:
		panic(fmt.Sprintf("WeeklyOnlineSeconds not equal %+v %+v", unmarshalStruct, unmarshalInterface))
	}
}

type SingleChatData struct {
	Index         int64  `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	FromPlayer    uint64 `protobuf:"varint,2,opt,name=FromPlayer,proto3" json:"FromPlayer,omitempty"`
	ToPlayer      uint64 `protobuf:"varint,3,opt,name=ToPlayer,proto3" json:"ToPlayer,omitempty"`
	Content       string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	SendTimestamp int64  `protobuf:"varint,5,opt,name=SendTimestamp,proto3" json:"SendTimestamp,omitempty"`
	BlockSender   bool   `protobuf:"varint,6,opt,name=BlockSender,proto3" json:"BlockSender,omitempty"`
}
type FriendChatDataList struct {
	ChatDataList []*SingleChatData `protobuf:"bytes,1,rep,name=ChatDataList,proto3" json:"ChatDataList,omitempty"`
	UpdateAt     int64             `protobuf:"varint,2,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
}

func SliceMarshalFoo() {
	se := &SingleChatData{
		Index:         1,
		FromPlayer:    2,
		ToPlayer:      3,
		Content:       "1",
		SendTimestamp: 4,
	}
	ss := &FriendChatDataList{}
	ss.ChatDataList = append(ss.ChatDataList, se)
	j, _ := json.Marshal(&ss.ChatDataList)
	fmt.Printf("%v\n", string(j))

	_ss := &FriendChatDataList{}
	_ss.ChatDataList = make([]*SingleChatData, 0)
	e := json.Unmarshal(j, &_ss.ChatDataList)
	fmt.Printf("e %v\n", e)
	fmt.Printf("_ss.ChatDataList %+v\n", _ss.ChatDataList)
}
