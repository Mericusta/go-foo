package astfoo

import (
	"io/fs"
	"reflect"
	"testing"
)

func TestParseFileFoo(t *testing.T) {
	type args struct {
		parseFilePath string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []string
		want2 []string
		want3 []string
		want4 []string
		want5 []string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{parseFilePath: "d:\\Projects\\go-foo\\src\\algorithm-foo\\algorithm.go"},
			"algorithmfoo",
			[]string{
				"\"math\"",
				"\"strings\"",
				"\"time\"",
				"\"github.com/pjebs/optimus-go\"",
			},
			[]string{
				"AntiAddictionData",
				"BasicContext",
				"Coordinate",
				"RenderContext",
				"Size",
				"Unit",
				"antiAddictionNormalCfg",
				"antiAddictionSpecialCfg",
			},
			[]string{
				"CalculateYearsOld",
				"ConvertCamelCase2SnakeCaseWithPhrase",
				"DecodeID",
				"EncodeID",
				"NewBasicContext",
			},
			[]string{
				"AntiAddictionSpecialDateFromat",
				"PrimeNum",
				"RandNum",
			},
			[]string{
				"optimusPrime",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4, got5 := ParseFileFoo(tt.args.parseFilePath)
			if got != tt.want {
				t.Errorf("ParseFileFoo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseFileFoo() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ParseFileFoo() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("ParseFileFoo() got3 = %v, want %v", got3, tt.want3)
			}
			if !reflect.DeepEqual(got4, tt.want4) {
				t.Errorf("ParseFileFoo() got4 = %v, want %v", got4, tt.want4)
			}
			if !reflect.DeepEqual(got5, tt.want5) {
				t.Errorf("ParseFileFoo() got5 = %v, want %v", got5, tt.want5)
			}
		})
	}
}

func TestParseDirFoo(t *testing.T) {
	type args struct {
		parseDirPath string
		filter       func(fs.FileInfo) bool
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{parseDirPath: "d:\\Projects\\go-foo\\src\\struct-foo"},
			[]string{"structfoo"},
			[]string{
				"d:\\Projects\\go-foo\\src\\struct-foo\\struct.go",
				"d:\\Projects\\go-foo\\src\\struct-foo\\struct_benchmark_test.go",
				"d:\\Projects\\go-foo\\src\\struct-foo\\struct_test.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseDirFoo(tt.args.parseDirPath, tt.args.filter)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDirFoo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ParseDirFoo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFormatFoo(t *testing.T) {
	type args struct {
		parseFilePath  string
		outputFunction string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// {
		// 	"test case 1",
		// 	args{
		// 		parseFilePath:  "D:\\Projects\\go-foo\\src\\algorithm-foo\\algorithm.go",
		// 		outputFunction: "ConvertCamelCase2SnakeCaseWithPhrase",
		// 	},
		// },
		// {
		// 	"test case 2",
		// 	args{
		// 		parseFilePath:  "D:\\Projects\\go-foo\\cmd\\main.go",
		// 		outputFunction: "main",
		// 	},
		// },
		{
			"test case 3",
			args{
				parseFilePath:  "../../cmd/main.go",
				outputFunction: "main",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FormatFoo(tt.args.parseFilePath, tt.args.outputFunction)
		})
	}
}

func TestMultiParseFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MultiParseFoo()
		})
	}
}
