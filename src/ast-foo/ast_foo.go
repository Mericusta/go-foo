package astfoo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"

	stpmap "github.com/Mericusta/go-stp/map"
)

// ParseFileFoo
// @filepath 待解析的文件地址
// @return   文件的包名称，导入的包路径，定义的结构体名称，定义函数/方法名称，全局常量名称，全局变量名称
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
