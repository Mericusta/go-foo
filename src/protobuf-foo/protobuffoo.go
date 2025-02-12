package protobuffoo

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"go-foo/src/protobuf-foo/pb"
	"os"
	"os/exec"

	"google.golang.org/protobuf/proto"
)

// func UnmarshalNilPointer() {
// 	var data *cargo_def.DungeonBaseInfo
// 	data = &cargo_def.DungeonBaseInfo{
// 		Result:          1,
// 		LastTime:        2,
// 		LimitCount:      make(map[int32]int32),
// 		FinishCondition: nil,
// 		Roles:           map[uint64]int32{1: 1},
// 	}
// 	pbData, err := proto.Marshal(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	data = nil

// 	data = &cargo_def.DungeonBaseInfo{}
// 	err = proto.Unmarshal(pbData, data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("data %+v\n", data)
// 	fmt.Printf("data.Result %v\n", data.Result)
// 	fmt.Printf("data.LastTime %v\n", data.LastTime)
// 	fmt.Printf("data.LimitCount %v, is nil %v\n", data.LimitCount, data.LimitCount == nil)
// 	fmt.Printf("data.FinishCondition %v, is nil %v\n", data.FinishCondition, data.FinishCondition == nil)
// 	fmt.Printf("data.Roles %v, is nil %v\n", data.Roles, data.Roles == nil)
// }

// func InterfaceMarshalFoo() {
// 	s := &cargo_def.DungeonBaseInfo{
// 		Result:          1,
// 		LastTime:        2,
// 		LimitCount:      make(map[int32]int32),
// 		FinishCondition: nil,
// 		Roles:           map[uint64]int32{1: 1},
// 	}
// 	marshalStruct, err := proto.Marshal(s)
// 	if err != nil {
// 		panic(err)
// 	}

// 	i := func() proto.Message {
// 		return &cargo_def.DungeonBaseInfo{
// 			Result:          1,
// 			LastTime:        2,
// 			LimitCount:      make(map[int32]int32),
// 			FinishCondition: nil,
// 			Roles:           map[uint64]int32{1: 1},
// 		}
// 	}()
// 	marshalInterface, err := proto.Marshal(i)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if len(marshalStruct) != len(marshalInterface) {
// 		panic(fmt.Sprintf("len not equal %v %v", len(marshalStruct), len(marshalInterface)))
// 	}

// 	for index := 0; index != len(marshalStruct); index++ {
// 		if marshalStruct[index] != marshalInterface[index] {
// 			panic(fmt.Sprintf("index %v b not equal", index))
// 		}
// 	}

// 	unmarshalStruct := &cargo_def.DungeonBaseInfo{}
// 	err = proto.Unmarshal(marshalStruct, unmarshalStruct)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// unmarshalInterface := &cargo_def.DungeonBaseInfo{}
// 	var unmarshalInterface proto.Message
// 	err = proto.Unmarshal(marshalInterface, unmarshalInterface)
// 	if err != nil {
// 		panic(err)
// 	}

// 	switch {
// 	case unmarshalStruct.Result != unmarshalInterface.(*cargo_def.DungeonBaseInfo).Result:
// 		panic(fmt.Sprintf("Result not equal %+v %+v", unmarshalStruct, unmarshalInterface))
// 	case unmarshalStruct.LastTime != unmarshalInterface.(*cargo_def.DungeonBaseInfo).LastTime:
// 		panic(fmt.Sprintf("LastTime not equal %+v %+v", unmarshalStruct, unmarshalInterface))
// 	}
// }

// func MarshalEmptyStructFoo() (int, int) {
// 	emptyMsg := &cargo_def.BagDataInfo{}
// 	notEmptyMsg := &cargo_def.BagDataInfo{}
// 	notEmptyMsg.Items = make(map[uint64]*cargo_def.ItemInfo)
// 	notEmptyMsg.Items[1] = &cargo_def.ItemInfo{
// 		GUID: 1, ClassID: 1, Count: 1, OverTime: 0,
// 	}
// 	emptyMsgResult, err := proto.Marshal(emptyMsg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	notEmptyMsgResult, err := proto.Marshal(notEmptyMsg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return len(emptyMsgResult), len(notEmptyMsgResult)
// }

func UnmarshalUnknownStruct(dataBytes []byte, output bool) error {
	// 直接将 []byte 写入到文件里
	dataBytesFilename := "UnmarshalUnknownStruct.bin"
	err := os.WriteFile(dataBytesFilename, dataBytes, 0644)
	if err != nil {
		return err
	}

	if output {
		manuallyCmdExpression := fmt.Sprintf("protoc --decode_raw < %v", dataBytesFilename)
		fmt.Printf("please call command manually: |%v|\n", manuallyCmdExpression)
	}

	// // 调用外部命令 protoc --decode_raw < UnmarshalUnknownStruct.bin
	catPath, err := exec.LookPath("cat")
	if err != nil {
		return err
	}
	if output {
		fmt.Printf("cat path: %v\n", catPath)
	}

	protocPath, err := exec.LookPath("protoc")
	if err != nil {
		return err
	}
	if output {
		fmt.Printf("protoc path: %v\n", protocPath)
	}

	pwdPath, err := exec.LookPath("pwd")
	if err != nil {
		return err
	}
	if output {
		fmt.Printf("pwd path: %v\n", pwdPath)
	}
	pwdCmd := exec.Command(pwdPath)
	pwdCmd.Stdout = &bytes.Buffer{}
	pwdCmd.Stderr = &bytes.Buffer{}
	err = pwdCmd.Run()
	if err != nil {
		return err
	}
	// pwdStdout := pwdCmd.Stdout.(*bytes.Buffer).String()
	if output {
		fmt.Printf("pwd std out: %v\n", pwdCmd.Stdout.(*bytes.Buffer).String())
		fmt.Printf("pwd std err: %v\n", pwdCmd.Stderr.(*bytes.Buffer).String())
	}

	shPath, err := exec.LookPath("sh")
	if err != nil {
		return err
	}
	if output {
		fmt.Printf("sh path: %v\n", shPath)
	}

	cmdExpression := fmt.Sprintf("%v --decode_raw < %v", protocPath, dataBytesFilename)
	// cmdExpression := fmt.Sprintf("%v %v", catPath, dataBytesFilename)
	if output {
		fmt.Printf("cmd expression: %v\n", cmdExpression)
	}

	cmd := exec.Command(shPath, "-c", cmdExpression)
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}
	err = cmd.Run()
	if err != nil {
		if output {
			fmt.Printf("%v err: %v\n, std err: \n|\n%v|\n", protocPath, err, cmd.Stderr.(*bytes.Buffer).String())
		}
		return err
	}
	fmt.Printf("%v std out: \n|\n%v|\n", protocPath, cmd.Stdout.(*bytes.Buffer).String())
	return nil
}

// UnmarshalUnknownStructFoo 反序列化未知结构
func UnmarshalUnknownStructFoo() {
	// 构造 proto 结构体
	data := &pb.DungeonFightRole{}
	// 填入特殊数值
	data.PlayerRoleGUID = 1234
	data.PresetRoleID = 4321
	dataBytes, err := proto.Marshal(data)
	if err != nil {
		panic(err)
	}
	// 输出为 []byte，16进制数组
	fmt.Printf("dataBytes: []byte |%v|\n", dataBytes)
	// 转换为 hex 字符串
	hexString := hex.EncodeToString(dataBytes)
	// 输出为 hex 字符串
	// 可直接用于 https://protogen.marcgravell.com/decode 解析
	fmt.Printf("dataBytes: hex string |%v|\n", hexString)

	err = UnmarshalUnknownStruct(dataBytes, true)
	if err != nil {
		panic(err)
	}
}
