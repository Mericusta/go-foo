package workerfoo

import (
	"fmt"
	"go-foo/utility"
	"io"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

// 以文件为单位提取
// 以包为单位整合

type GoPackageInfo struct {
	Name    string                 // 包名：package workerfoo -> workerfoo
	Project string                 // 所属项目名：project/A/B/C -> project
	Path    string                 // 包相对路径：project/A/B/C -> A/B/C
	FileMap map[string]*GoFileInfo // 包中文件信息，键是 GoFileInfo.Path
}

// ImportPath 生成引入该包时的引入别名和引入路径：project + A/B/C -> import workerfoo "go-foo/worker-foo"
func (g *GoPackageInfo) ImportPath() string {
	if filepath.Base(g.Path) == g.Name {
		return filepath.ToSlash(fmt.Sprintf("%v/%v", g.Project, g.Path))
	} else {
		return filepath.ToSlash(fmt.Sprintf("%v %v/%v", g.Name, g.Project, g.Path))
	}
}

type GoFileInfo struct {
	Path                string                   // 相对项目根目录的路径：worker-foo/package.go
	StructDefinitionMap map[string]*GoStructInfo // 该文件定义的结构体，键是 GoStructInfo.Name
	// ImportStruct        map[string]map[string]struct{} // 该文件引入的外部结构体，键是包名 GoPackageInfo.Name，值是 GoStructInfo.Name
}

// CleanFileComment 置空文件中所有注释
func CleanFileComment(r io.Reader) string {
	fileContent, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	isBlock, isComment := false, false
	firstCommentIndex, secondCommentIndex := -1, -1
	builder, commentBuffer := strings.Builder{}, strings.Builder{}
	for index, b := range fileContent {
		switch rune(b) {
		case utility.PunctuationMarkLeftDoubleQuotes:
			if !isComment {
				if !isBlock {
					isBlock = true
				} else {
					isBlock = false
				}
			}
		case '/':
			if !isBlock {
				if firstCommentIndex == -1 {
					firstCommentIndex = index
				} else if secondCommentIndex == -1 {
					secondCommentIndex = index
					isComment = true
					commentBuffer.Reset()
				}
			}
		case '\n':
			if isComment {
				isComment = false
				firstCommentIndex = -1
				secondCommentIndex = -1
				commentBuffer.Reset()
			}
		}

		if !isComment {
			if firstCommentIndex != -1 && secondCommentIndex == -1 {
				if commentBuffer.Len() > 0 {
					// just one /, clear comment buffer
					builder.WriteString(commentBuffer.String())
					builder.WriteByte(b)
					firstCommentIndex = -1
					commentBuffer.Reset()
				} else {
					// first match /
					commentBuffer.WriteByte(b)
				}
			} else {
				builder.WriteByte(b)
			}
		}
	}

	return builder.String()
}

var (
	GO_PACKAGE_EXPRESSION            string = `package\s+(?P<NAME>[[:alpha:]][_\w]+)`
	GoPackageRegexp                         = regexp.MustCompile(GO_PACKAGE_EXPRESSION)
	GoPackageRegexpSubmatchNameIndex        = GoPackageRegexp.SubexpIndex("NAME")
)

// ExtractGoFilePackage 从 go 文件中提取包名
func ExtractGoFilePackage(r io.Reader) string {
	fileContent, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	submatch := GoPackageRegexp.FindSubmatch(fileContent)
	if GoPackageRegexpSubmatchNameIndex >= len(submatch) {
		panic("can not match package")
	}
	return string(submatch[GoPackageRegexpSubmatchNameIndex])
}

var (
	GO_IMPORT_EXPRESSION             string = `((?P<ALIAS>\w+)\s+)?"(?P<PATH>[/_\.\w-]+)"`
	GoImportRegexp                          = regexp.MustCompile(GO_IMPORT_EXPRESSION)
	GoImportRegexpSubmatchAliasIndex        = GoImportRegexp.SubexpIndex("ALIAS")
	GoImportRegexpSubmatchPathIndex         = GoImportRegexp.SubexpIndex("PATH")
)

// ExtractGoFileImportPackage 从 go 文件中提取引入的包别名及其引入路径
func ExtractGoFileImportPackage(r io.Reader) map[string]string {
	fileContent, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	importScopeBeginExpression := `import\s*\(`
	importScopeBeginRegexp := regexp.MustCompile(importScopeBeginExpression)
	importScopeBeginIndexSlice := importScopeBeginRegexp.FindIndex(fileContent)
	importScopeBeginRune := rune(fileContent[importScopeBeginIndexSlice[1]-1])
	importScopeEndRune := utility.GetAnotherPunctuationMark(importScopeBeginRune)
	importScopeLength := utility.CalculatePunctuationMarksContentLength(
		string(fileContent[importScopeBeginIndexSlice[1]+1:]),
		importScopeBeginRune,
		importScopeEndRune,
		utility.InvalidScopePunctuationMarkMap,
	)
	if importScopeLength == 0 {
		panic("Error: file can not find import scope end index")
	}
	// fmt.Printf("import scope content: |%v|\n", string(fileContent[importScopeBeginIndexSlice[1]+1:importScopeBeginIndexSlice[1]+1+importScopeLength]))
	importAliasPathMap := make(map[string]string)
	for _, eachImportString := range strings.Split(string(fileContent[importScopeBeginIndexSlice[1]+1:importScopeBeginIndexSlice[1]+1+importScopeLength]), "\n") {
		for _, submatchSlice := range GoImportRegexp.FindAllStringSubmatch(strings.TrimSpace(eachImportString), -1) {
			if GoImportRegexpSubmatchAliasIndex != -1 {
				importAliasPathMap[submatchSlice[GoImportRegexpSubmatchAliasIndex]] = submatchSlice[GoImportRegexpSubmatchPathIndex]
			} else {
				importAliasPathMap[filepath.Base(submatchSlice[GoImportRegexpSubmatchPathIndex])] = submatchSlice[GoImportRegexpSubmatchPathIndex]
			}
		}
	}
	return importAliasPathMap
}
