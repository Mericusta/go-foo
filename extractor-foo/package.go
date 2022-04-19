package extractorfoo

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

var (
	GO_PACKAGE_EXPRESSION            string = `package\s+(?P<NAME>[[:alpha:]][_\w]+)`
	GoPackageRegexp                         = regexp.MustCompile(GO_PACKAGE_EXPRESSION)
	GoPackageRegexpSubmatchNameIndex        = GoPackageRegexp.SubexpIndex("NAME")
	GO_IMPORT_EXPRESSION             string = `((?P<ALIAS>\w+)\s+)?"(?P<PATH>[/_\.\w-]+)"`
	GoImportRegexp                          = regexp.MustCompile(GO_IMPORT_EXPRESSION)
	GoImportRegexpSubmatchAliasIndex        = GoImportRegexp.SubexpIndex("ALIAS")
	GoImportRegexpSubmatchPathIndex         = GoImportRegexp.SubexpIndex("PATH")
)

type GoPackageInfo struct {
	Name    string
	Path    string
	FileMap map[string]*GoFileInfo
}

func (gpi *GoPackageInfo) ImportPath(project string) string {
	return filepath.ToSlash(fmt.Sprintf("%v/%v", project, gpi.Path))
}

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
	if importScopeLength < 0 {
		panic("Error: file can not find import scope end index")
	}
	// fmt.Printf("import scope content: |%v|\n", string(fileContent[importScopeBeginIndexSlice[1]+1:importScopeBeginIndexSlice[1]+1+importScopeLength]))
	importAliasPathMap := make(map[string]string)
	for _, eachImportString := range strings.Split(string(fileContent[importScopeBeginIndexSlice[1]+1:importScopeBeginIndexSlice[1]+1+importScopeLength]), "\n") {
		for _, submatchSlice := range GoImportRegexp.FindAllStringSubmatch(strings.TrimSpace(eachImportString), -1) {
			if GoImportRegexpSubmatchAliasIndex == -1 || len(submatchSlice[GoImportRegexpSubmatchAliasIndex]) == 0 {
				importAliasPathMap[filepath.Base(submatchSlice[GoImportRegexpSubmatchPathIndex])] = submatchSlice[GoImportRegexpSubmatchPathIndex]
			} else {
				importAliasPathMap[submatchSlice[GoImportRegexpSubmatchAliasIndex]] = submatchSlice[GoImportRegexpSubmatchPathIndex]
			}
		}
	}
	return importAliasPathMap
}
