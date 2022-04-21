package extractorfoo

import (
	"fmt"
	"go-foo/utility"
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

func ExtractGoFilePackage(fileContent []byte) string {
	submatch := GoPackageRegexp.FindSubmatch(fileContent)
	if GoPackageRegexpSubmatchNameIndex >= len(submatch) {
		panic("can not match package")
	}
	return string(submatch[GoPackageRegexpSubmatchNameIndex])
}

func ExtractGoFileImportPackage(fileContent []byte) (map[string]string, bool) {
	importScopeContent, singleWithBrackets := ExtractGoFileImportScopeContent(fileContent)
	if len(importScopeContent) == 0 {
		return nil, false
	}
	// fmt.Printf("import scope content: |%v|\n", string(fileContent[importScopeBeginIndexSlice[1]+1:importScopeBeginIndexSlice[1]+1+importScopeLength]))
	importAliasPathMap := make(map[string]string)
	for _, eachImportString := range strings.Split(string(importScopeContent), "\n") {
		for _, submatchSlice := range GoImportRegexp.FindAllStringSubmatch(strings.TrimSpace(eachImportString), -1) {
			if GoImportRegexpSubmatchAliasIndex == -1 || len(submatchSlice[GoImportRegexpSubmatchAliasIndex]) == 0 {
				importAliasPathMap[filepath.Base(submatchSlice[GoImportRegexpSubmatchPathIndex])] = submatchSlice[GoImportRegexpSubmatchPathIndex]
			} else {
				importAliasPathMap[submatchSlice[GoImportRegexpSubmatchAliasIndex]] = submatchSlice[GoImportRegexpSubmatchPathIndex]
			}
		}
	}
	return importAliasPathMap, singleWithBrackets
}

func ExtractGoFileImportScopeContent(fileContent []byte) ([]byte, bool) {
	importScopeBeginExpression := `import\s*\(`
	importScopeBeginRegexp := regexp.MustCompile(importScopeBeginExpression)
	importScopeBeginIndexSlice := importScopeBeginRegexp.FindIndex(fileContent)
	if len(importScopeBeginIndexSlice) == 0 {
		singleImportScopeExpression := `import\s+(?P<IMPORT_CONTENT>(\w+\s+)?"\S+")`
		singleImportScopeRegexp := regexp.MustCompile(singleImportScopeExpression)
		singleImportScopeSubmatchImportContentIndex := singleImportScopeRegexp.SubexpIndex("IMPORT_CONTENT")
		subMatchSlice := singleImportScopeRegexp.FindSubmatch(fileContent)
		if len(subMatchSlice) > 0 {
			return subMatchSlice[singleImportScopeSubmatchImportContentIndex], false
		} else {
			return nil, false
		}
	} else {
		return utility.GetScopeContentBetweenPunctuationMarks(fileContent, importScopeBeginIndexSlice[1]-1), true
	}
}

func ExtractGoFileSingleImportScope(fileContent []byte) []byte {
	singleImportScopeExpression := `import\s+(?P<IMPORT_CONTENT>(\w+\s+)?"\S+")`
	singleImportScopeRegexp := regexp.MustCompile(singleImportScopeExpression)
	singleImportScopeIndexSlice := singleImportScopeRegexp.FindIndex(fileContent)
	if len(singleImportScopeIndexSlice) > 0 {
		return fileContent[singleImportScopeIndexSlice[0]:singleImportScopeIndexSlice[1]]
	}
	return nil
}

// ----------------------------------------------------------------

var (
	GO_IMPORT_SCOPE_TEMPLATE                         = "import (\n[IMPORT_CONTENT]\n)"
	GO_IMPORT_CONTENT_TEMPLATE                       = "\t[IMPORT_ALIAS] \"[IMPORT_PATH]\"\n"
	GO_IMPORT_MAKE_UP_REPLACE_KEYWORD_IMPORT_CONTENT = "[IMPORT_CONTENT]"
	GO_IMPORT_MAKE_UP_REPLACE_KEYWORD_IMPORT_ALIAS   = "[IMPORT_ALIAS]"
	GO_IMPORT_MAKE_UP_REPLACE_KEYWORD_IMPORT_PATH    = "[IMPORT_PATH]"
)

func MakeUpGoFileImportScope(importPkgMap map[string]string) []byte {
	builder := strings.Builder{}
	for importPkgAlias, importPkgPath := range importPkgMap {
		replaceAlias := importPkgAlias
		if filepath.Base(importPkgPath) == importPkgAlias {
			replaceAlias = ""
		}
		importContent := strings.Replace(GO_IMPORT_CONTENT_TEMPLATE, GO_IMPORT_MAKE_UP_REPLACE_KEYWORD_IMPORT_PATH, importPkgPath, -1)
		importContent = strings.Replace(importContent, GO_IMPORT_MAKE_UP_REPLACE_KEYWORD_IMPORT_ALIAS, replaceAlias, -1)
		builder.WriteString(importContent)
	}
	// fmt.Printf("%v\n", strings.Replace(GO_IMPORT_SCOPE_TEMPLATE, GO_IMPORT_MAKE_UP_REPLACE_KEYWORD_IMPORT_CONTENT, builder.String(), -1))
	return []byte(strings.Replace(GO_IMPORT_SCOPE_TEMPLATE, GO_IMPORT_MAKE_UP_REPLACE_KEYWORD_IMPORT_CONTENT, builder.String(), -1))
}
