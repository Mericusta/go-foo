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

	"github.com/Mericusta/go-stp"
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
		keys := stp.Key(fileAST.Scope.Objects)
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
		filepath = stp.Key(pkg.Files) // 里面是 *ast.File，同 parser.ParseFile
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

	fileAST, err := parser.ParseFile(fileSet, "", parseFileContent, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.OpenFile(fmt.Sprintf("%v.bak", parseFilePath), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	ast.Print(fileSet, fileAST)

	var funcDecl *ast.FuncDecl
	ast.Inspect(fileAST, func(n ast.Node) bool {
		if n == fileAST {
			return true
		}
		if n == nil || funcDecl != nil {
			return false
		}
		decl, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}
		if decl.Recv == nil && decl.Name.String() == outputFunction {
			funcDecl = decl
			return false
		}
		return true
	})
	if funcDecl == nil {
		panic("nil")
	}

	buffer := &bytes.Buffer{}
	ast.Print(fileSet, funcDecl)
	// Note pos 需要-1
	fmt.Printf("decl.Pos() = %v, decl.End() = %v, content = \n|%v|\n", funcDecl.Pos(), funcDecl.End(), string(parseFileContent[funcDecl.Pos()-1:funcDecl.End()]))
	if declLen := funcDecl.End() - funcDecl.Pos(); buffer.Cap() < int(declLen) {
		buffer.Grow(int(declLen))
	}
	// funcDecl.Doc.List[0].Text, funcDecl.Doc.List[1].Text = funcDecl.Doc.List[1].Text, funcDecl.Doc.List[0].Text
	err = format.Node(buffer, fileSet, funcDecl)
	if err != nil {
		panic(err)
	}
	outputFile.Write(buffer.Bytes())
	buffer.Reset()
}

func MultiParseFoo() {
	f1 := "../../cmd/main.go"
	f2 := "../concurrency-foo/concurrency.go"

	f1Content, _ := os.ReadFile(f1)
	f2Content, _ := os.ReadFile(f2)

	fileSet := token.NewFileSet()
	file1AST, err := parser.ParseFile(fileSet, "", f1Content, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Print(fileSet, file1AST)

	_, err = parser.ParseFile(fileSet, "", f2Content, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Print(fileSet, file1AST)
}

// ParseExpressionFoo
// @expression        待解析的表达式
func ParseExpressionFoo(expression string) {
	exprAST, err := parser.ParseExpr(expression)
	if err != nil {
		panic(err)
	}

	err = ast.Print(token.NewFileSet(), exprAST)
	if err != nil {
		panic(err)
	}

	buffer := &bytes.Buffer{}
	err = format.Node(buffer, token.NewFileSet(), exprAST)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buffer.Bytes())
	fmt.Println()
}
