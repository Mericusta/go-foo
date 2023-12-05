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

func InterfaceMarshalFoo() {
	s := &cargo_def.DungeonBaseInfo{
		Result:          1,
		LastTime:        2,
		LimitCount:      make(map[int32]int32),
		FinishCondition: nil,
		Roles:           map[uint64]int32{1: 1},
	}
	marshalStruct, err := proto.Marshal(s)
	if err != nil {
		panic(err)
	}

	i := func() proto.Message {
		return &cargo_def.DungeonBaseInfo{
			Result:          1,
			LastTime:        2,
			LimitCount:      make(map[int32]int32),
			FinishCondition: nil,
			Roles:           map[uint64]int32{1: 1},
		}
	}()
	marshalInterface, err := proto.Marshal(i)
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

	unmarshalStruct := &cargo_def.DungeonBaseInfo{}
	err = proto.Unmarshal(marshalStruct, unmarshalStruct)
	if err != nil {
		panic(err)
	}

	// unmarshalInterface := &cargo_def.DungeonBaseInfo{}
	var unmarshalInterface proto.Message
	err = proto.Unmarshal(marshalInterface, unmarshalInterface)
	if err != nil {
		panic(err)
	}

	switch {
	case unmarshalStruct.Result != unmarshalInterface.(*cargo_def.DungeonBaseInfo).Result:
		panic(fmt.Sprintf("Result not equal %+v %+v", unmarshalStruct, unmarshalInterface))
	case unmarshalStruct.LastTime != unmarshalInterface.(*cargo_def.DungeonBaseInfo).LastTime:
		panic(fmt.Sprintf("LastTime not equal %+v %+v", unmarshalStruct, unmarshalInterface))
	}
}

func MarshalEmptyStructFoo() (int, int) {
	emptyMsg := &cargo_def.BagDataInfo{}
	notEmptyMsg := &cargo_def.BagDataInfo{}
	notEmptyMsg.Items = make(map[uint64]*cargo_def.ItemInfo)
	notEmptyMsg.Items[1] = &cargo_def.ItemInfo{
		GUID: 1, ClassID: 1, Count: 1, OverTime: 0,
	}
	emptyMsgResult, err := proto.Marshal(emptyMsg)
	if err != nil {
		panic(err)
	}
	notEmptyMsgResult, err := proto.Marshal(notEmptyMsg)
	if err != nil {
		panic(err)
	}
	return len(emptyMsgResult), len(notEmptyMsgResult)
}
