package protobuffoo

import (
	"fmt"
	cargo_def "go-foo/src/protobuf-foo/pb"

	"google.golang.org/protobuf/proto"
)

func UnmarshalNilPointer() {
	var data *cargo_def.DungeonBaseInfo
	data = &cargo_def.DungeonBaseInfo{
		Result:          1,
		LastTime:        2,
		LimitCount:      make(map[int32]int32),
		FinishCondition: nil,
		Roles:           map[uint64]int32{1: 1},
	}
	pbData, err := proto.Marshal(data)
	if err != nil {
		panic(err)
	}
	data = nil

	data = &cargo_def.DungeonBaseInfo{}
	err = proto.Unmarshal(pbData, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data %+v\n", data)
	fmt.Printf("data.Result %v\n", data.Result)
	fmt.Printf("data.LastTime %v\n", data.LastTime)
	fmt.Printf("data.LimitCount %v, is nil %v\n", data.LimitCount, data.LimitCount == nil)
	fmt.Printf("data.FinishCondition %v, is nil %v\n", data.FinishCondition, data.FinishCondition == nil)
	fmt.Printf("data.Roles %v, is nil %v\n", data.Roles, data.Roles == nil)
}
