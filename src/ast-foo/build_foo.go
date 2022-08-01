package astfoo

import (
	"go/build"
)

// build.Import 解析 $GOROOT 或 $GOPATH 目录中可导入的包的信息
// build.Import 可以设置 build.Default 的 GOROOT/GOPATH 以及环境变量中的 GOROOT/GOPATH 以改变路径
func ImportFoo(goPkgPath, srcPath string, mode build.ImportMode) (string, string, string, string, []string) {
	p, err := build.Import(goPkgPath, srcPath, mode)
	if err != nil {
		panic(err)
	}
	return p.Dir, p.Name, p.ImportPath, p.Root, p.GoFiles
}

// build.ImportDir 解析指定目录下的包信息
func ImportDirFoo(path string, mode build.ImportMode) (string, string, string, string, []string) {
	p, err := build.ImportDir(path, mode)
	if err != nil {
		panic(err)
	}
	return p.Dir, p.Name, p.ImportPath, p.SrcRoot, p.GoFiles
}
