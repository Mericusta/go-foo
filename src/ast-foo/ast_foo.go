package astfoo

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"

	stpmap "github.com/Mericusta/go-stp/map"
)

// ParseFileFoo
// @parseFilePath 待解析的文件路径
// @return        文件的包名称，导入的包路径，定义的结构体名称，定义函数/方法名称，全局常量名称，全局变量名称
func ParseFileFoo(parseFilePath string) (string, []string, []string, []string, []string, []string) {
	fileAST, err := parser.ParseFile(token.NewFileSet(), parseFilePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	fileImportPkgPath := make([]string, len(fileAST.Imports))
	for index, fileImportPkg := range fileAST.Imports {
		fileImportPkgPath[index] = fileImportPkg.Path.Value
	}

	var fileStructName, fileFuncName, fileConstantName, fileVariableName []string
	if fileAST.Scope != nil {
		keys := stpmap.Key(fileAST.Scope.Objects)
		sort.Strings(keys)
		for _, name := range keys {
			switch fileAST.Scope.Objects[name].Kind {
			case ast.Typ:
				fileStructName = append(fileStructName, name)
			case ast.Fun:
				fileFuncName = append(fileFuncName, name)
			case ast.Con:
				fileConstantName = append(fileConstantName, name)
			case ast.Var:
				fileVariableName = append(fileVariableName, name)
			}
		}

	}

	return fileAST.Name.Name, fileImportPkgPath, fileStructName, fileFuncName, fileConstantName, fileVariableName
}

// ParseDirFoo
// @parseDirPath 待解析的目录路径
// @filter       文件筛选器
// @return       目录的包名称，包路径
func ParseDirFoo(parseDirPath string, filter func(fs.FileInfo) bool) ([]string, []string) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), parseDirPath, filter, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var dirPkgs, filepath []string
	for _, pkg := range pkgs {
		dirPkgs = append(dirPkgs, pkg.Name)
		filepath = stpmap.Key(pkg.Files) // 里面是 *ast.File，同 parser.ParseFile
		sort.Strings(filepath)
	}

	return dirPkgs, filepath
}

// FormatFoo
// @parseFilePath  待解析的文件路径
// @outputFunction 待输出的函数名称
func FormatFoo(parseFilePath, outputFunction string) {
	parseFileContent, err := ioutil.ReadFile(parseFilePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("parseFileContent\n|%v|\n", string(parseFileContent))

	fileSet := token.NewFileSet()

	fileAST, err := parser.ParseFile(fileSet, parseFilePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.OpenFile(fmt.Sprintf("%v.bak", parseFilePath), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	buffer := &bytes.Buffer{}
	if fileAST.Scope != nil {
		for name, object := range fileAST.Scope.Objects {
			if object.Kind == ast.Fun && name == outputFunction {
				ast.Print(fileSet, object)
				decl := object.Decl.(*ast.FuncDecl)
				fmt.Printf("decl.End() = %v, decl.Pos() = %v, content = \n|%v|\n", decl.End(), decl.Pos(), string(parseFileContent)[decl.Pos():decl.End()])
				if declLen := decl.End() - decl.Pos(); buffer.Cap() < int(declLen) {
					buffer.Grow(int(declLen))
				}
				// decl.Doc.List[0].Text, decl.Doc.List[1].Text = decl.Doc.List[1].Text, decl.Doc.List[0].Text
				err = format.Node(buffer, fileSet, decl)
				if err != nil {
					panic(err)
				}
				outputFile.Write(buffer.Bytes())
				buffer.Reset()
				break
			}
		}
	}
}
