package astfoo

import (
	"go/build"
	"reflect"
	"testing"
)

func TestImportFoo(t *testing.T) {
	type args struct {
		importPkgPath   string
		projectRootPath string
		mode            build.ImportMode
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
				importPkgPath: "go/build",
				mode:          0,
			},
			"C:\\Program Files\\Go\\src\\go\\build",
			"build",
			"go/build",
			"C:\\Program Files\\Go",
			[]string{"build.go", "doc.go", "gc.go", "read.go", "syslist.go", "zcgo.go"},
		},
		// {
		// 	"test case 2",
		// 	args{
		// 		importPkgPath:   "go-foo/array-foo",
		// 		projectRootPath: "S:\\Projects\\go-foo",
		// 		mode:            0,
		// 	},
		// 	"s:\\Projects\\go\\go-foo\\src\\array-foo",
		// 	"arrayfoo",
		// 	"go-foo/array-foo",
		// 	"s:\\Projects\\go\\go-foo\\src",
		// 	[]string{"array.go"},
		// },
		{
			"test case 2",
			args{
				importPkgPath:   "go-foo/array-foo",
				projectRootPath: "D:\\Projects\\go-foo",
				mode:            0,
			},
			"d:\\Projects\\go-foo\\src\\array-foo",
			"arrayfoo",
			"go-foo/array-foo",
			"d:\\Projects\\go-foo\\src",
			[]string{"array.go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4 := ImportFoo(tt.args.importPkgPath, tt.args.projectRootPath, tt.args.mode)
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
