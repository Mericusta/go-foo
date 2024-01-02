package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	data_exchange "go-foo/cmd/data-exchange/pb"

	"google.golang.org/protobuf/proto"
)

func newRobotsFightData() *data_exchange.RobotsFightData {
	robotsFightData := &data_exchange.RobotsFightData{}
	robotsFightData.Robots = make([]*data_exchange.FightPlayerData, 0, 8)
	for i := 0; i != 5; i++ {
		robot := &data_exchange.FightPlayerData{}
		robot.GroupID = int32(i)
		robot.PlayerID = uint64(i)
		robot.RobotID = int32(i)
		robot.PlayerName = fmt.Sprintf("%v", i)
		robot.MaxHP = int32(i)
		robot.CurrentHP = int32(i)
		robot.Wave = int32(i)
		robot.Score = int32(i)
		robot.ScoreLevel = int32(i)
		robot.Finished = true
		robot.Attacked = make(map[int32]int32)
		for j := 0; j != 5; j++ {
			if i == j {
				continue
			}
			robot.Attacked[int32(j)] = 1
		}
		robotsFightData.Robots = append(robotsFightData.Robots, robot)
	}
	return robotsFightData
}

func ProtoMarshalFoo[T proto.Message](v T) error {
	_, err := proto.Marshal(v)
	return err
}

func ProtoUnmarshalFoo[T proto.Message](b []byte, v T) error {
	err := proto.Unmarshal(b, v)
	return err
}

func ProtoFoo[T proto.Message](v T) error {
	b, err := proto.Marshal(v)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}

func JsonMarshalFoo[T any](v T) error {
	_, err := json.Marshal(v)
	return err
}

func JsonUnmarshalFoo[T any](b []byte, v T) error {
	err := json.Unmarshal(b, v)
	return err
}

func JsonFoo[T any](v T) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}

func GobMarshalFoo[T any](v T) error {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(v)
	return err
}

func GobUnmarshalFoo[T any](b []byte, v T) error {
	err := gob.NewDecoder(bytes.NewReader(b)).Decode(v)
	return err
}

func GobFoo[T any](v T) error {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(v)
	if err != nil {
		return err
	}
	err = gob.NewDecoder(&buffer).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func main() {

}
