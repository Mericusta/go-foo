package astfoo

import (
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
