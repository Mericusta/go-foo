package astfoo

import (
	"go/build"
	"reflect"
	"testing"
)

func TestImportFoo(t *testing.T) {
	type args struct {
		path   string
		gopath string
		mode   build.ImportMode
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
		want3 string
		want4 []string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				path:   "go/build",
				gopath: "D:\\Projects\\go-foo",
				mode:   0,
			},
			"C:\\Program Files\\Go\\src\\go\\build",
			"build",
			"go/build",
			"C:\\Program Files\\Go",
			[]string{"build.go", "doc.go", "gc.go", "read.go", "syslist.go", "zcgo.go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4 := ImportFoo(tt.args.path, tt.args.gopath, tt.args.mode)
			if got != tt.want {
				t.Errorf("ImportFoo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ImportFoo() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ImportFoo() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("ImportFoo() got3 = %v, want %v", got3, tt.want3)
			}
			if !reflect.DeepEqual(got4, tt.want4) {
				t.Errorf("ImportFoo() got4 = %v, want %v", got4, tt.want4)
			}
		})
	}
}

func TestImportDirFoo(t *testing.T) {
	type args struct {
		path string
		mode build.ImportMode
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
		want3 string
		want4 []string
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				path: "D:\\Projects\\go-foo\\src\\array-foo",
				mode: 0,
			},
			"D:\\Projects\\go-foo\\src\\array-foo",
			"arrayfoo",
			".",
			"",
			[]string{"array.go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4 := ImportDirFoo(tt.args.path, tt.args.mode)
			if got != tt.want {
				t.Errorf("ImportDirFoo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ImportDirFoo() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ImportDirFoo() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("ImportDirFoo() got3 = %v, want %v", got3, tt.want3)
			}
			if !reflect.DeepEqual(got4, tt.want4) {
				t.Errorf("ImportDirFoo() got4 = %v, want %v", got4, tt.want4)
			}
		})
	}
}
