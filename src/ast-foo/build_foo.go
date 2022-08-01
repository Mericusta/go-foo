package astfoo

import (
	"go/build"
)

// ImportFoo
// @importPkgPath   在项目中使用 go 包时 import 的路径
// @projectRootPath 项目的绝对路径，不填写则是 $GOROOT/$GOPATH
// @return
func ImportFoo(importPkgPath, projectRootPath string, mode build.ImportMode) (string, string, string, string, []string) {
	p, err := build.Import(importPkgPath, projectRootPath, mode)
	if err != nil {
		panic(err)
	}
	return p.Dir, p.Name, p.ImportPath, p.Root, p.GoFiles
}
