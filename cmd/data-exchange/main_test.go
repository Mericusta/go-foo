package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	data_exchange "go-foo/cmd/data-exchange/pb"
	"reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

func Test_ProtoMarshalFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{v: newRobotsFightData()},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ProtoMarshalFoo(tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ProtoMarshalFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_ProtoUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		b []byte
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{
				b: func() []byte {
					b, _ := proto.Marshal(newRobotsFightData())
					return b
				}(),
				v: &data_exchange.RobotsFightData{},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ProtoUnmarshalFoo(tt.args.b, tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ProtoUnmarshalFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_JsonMarshalFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{v: newRobotsFightData()},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := JsonMarshalFoo(tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("JsonMarshalFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_JsonUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		b []byte
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{
				b: func() []byte {
					b, _ := json.Marshal(newRobotsFightData())
					return b
				}(),
				v: newRobotsFightData(),
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := JsonUnmarshalFoo(tt.args.b, tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("JsonUnmarshalFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_GobMarshalFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{v: newRobotsFightData()},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := GobMarshalFoo(tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("GobMarshalFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_GobUnmarshalFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		b []byte
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{
				b: func() []byte {
					var buffer bytes.Buffer
					gob.NewEncoder(&buffer).Encode(newRobotsFightData())
					return buffer.Bytes()
				}(),
				v: newRobotsFightData(),
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := GobUnmarshalFoo(tt.args.b, tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("GobUnmarshalFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_ProtoFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{v: newRobotsFightData()},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := ProtoFoo(tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("ProtoFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_JsonFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{v: newRobotsFightData()},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := JsonFoo(tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("JsonFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}

func Test_GobFoo_17ed74d8fa7607506f6248e12df7da56(t *testing.T) {
	type args struct {
		v *data_exchange.RobotsFightData
	}
	tests := []struct {
		name  string
		args  args
		want0 error
	}{
		{
			"test case 1",
			args{v: newRobotsFightData()},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got0 := GobFoo(tt.args.v)
			if !reflect.DeepEqual(got0, tt.want0) {
				t.Errorf("GobFoo() got0 = %v, want0 %v", got0, tt.want0)
			}
		})
	}
}
