package extractorfoo

import (
	"fmt"
	"go-foo/utility"
	"io"
	"regexp"
	"strings"
)

type GoInterfaceInfo struct {
	Name                   string
	FunctionDeclarationMap map[string]*GoFunctionDeclaration
}

var (
	GO_INTERFACE_DECLARATION_SCOPE_BEGIN_EXPRESSION string = `type\s+(?P<NAME>\w+)\s+interface\s+\{`
	GoInterfaceDeclarationScopeBeginRegexp                 = regexp.MustCompile(GO_INTERFACE_DECLARATION_SCOPE_BEGIN_EXPRESSION)
	GoInterfaceRegexpSubmatchNameIndex                     = GoInterfaceDeclarationScopeBeginRegexp.SubexpIndex("NAME")
	GoInterfaceDeclarationScopeBeginRune                   = '{'
)

func ExtractGoFileInterfaceDeclaration(r io.Reader) map[string]*GoInterfaceInfo {
	fileContent := CleanFileComment(r)

	fileInterfaceDeclarationMap := make(map[string]*GoInterfaceInfo)
	for _, interfaceDeclarationScopeBeginIndexSlice := range GoInterfaceDeclarationScopeBeginRegexp.FindAllStringIndex(string(fileContent), -1) {
		submatchSlice := GoInterfaceDeclarationScopeBeginRegexp.FindStringSubmatch(string(fileContent[interfaceDeclarationScopeBeginIndexSlice[0]:interfaceDeclarationScopeBeginIndexSlice[1]]))
		interfaceName := submatchSlice[GoInterfaceRegexpSubmatchNameIndex]
		if interfaceName != "GoInterfaceDeclaration" {
			continue
		}
		fileInterfaceDeclarationMap[interfaceName] = &GoInterfaceInfo{
			Name:                   interfaceName,
			FunctionDeclarationMap: make(map[string]*GoFunctionDeclaration),
		}

		// {
		// 	fmt.Println()
		// 	fmt.Printf("interfaceDeclarationScopeBeginIndexSlice = |%v|\n", interfaceDeclarationScopeBeginIndexSlice)
		// 	fmt.Printf("interfaceDeclarationScope = |%v|\n", string(fileContent[interfaceDeclarationScopeBeginIndexSlice[0]:interfaceDeclarationScopeBeginIndexSlice[1]]))
		// 	fmt.Printf("interfaceName = %v\n", interfaceName)
		// 	fmt.Println()
		// 	return nil
		// }

		interfaceDeclarationScopeBeginRune := rune(fileContent[interfaceDeclarationScopeBeginIndexSlice[1]-1])
		interfaceDeclarationScopeEndRune := utility.GetAnotherPunctuationMark(interfaceDeclarationScopeBeginRune)
		interfaceDeclarationLength := utility.CalculatePunctuationMarksContentLength(
			string(fileContent[interfaceDeclarationScopeBeginIndexSlice[1]+1:]),
			interfaceDeclarationScopeBeginRune,
			interfaceDeclarationScopeEndRune,
			utility.InvalidScopePunctuationMarkMap,
		)
		if interfaceDeclarationLength < 0 {
			fmt.Printf("Error: interface %v can not find interface end index\n", interfaceName)
			continue
		}

		// {
		// 	fmt.Println()
		// 	fmt.Printf("interface content = |%v|", string(fileContent[interfaceDeclarationScopeBeginIndexSlice[1]:interfaceDeclarationScopeBeginIndexSlice[1]+interfaceDeclarationLength]))
		// 	fmt.Println()
		// 	return nil
		// }

		for _, lineContent := range strings.Split(string(fileContent[interfaceDeclarationScopeBeginIndexSlice[1]:interfaceDeclarationScopeBeginIndexSlice[1]+interfaceDeclarationLength]), "\n") {
			trimSpaceString := strings.TrimSpace(lineContent)
			if len(trimSpaceString) == 0 {
				continue
			}

			submatchSlice := GoFunctionDeclarationScopeBeginRegexp.FindStringSubmatch(trimSpaceString)
			if len(submatchSlice) == 0 {
				continue
			}
			fmt.Printf("|%v|\n", trimSpaceString)
			// signature
			functionName := submatchSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchNameIndex]
			// params
			// fmt.Printf("submatchSlice = %v\n", submatchSlice)
			// fmt.Printf("submatchSlice[%v] = %v\n", GoStructRegexpSubmatchNameIndex, submatchSlice[GoStructRegexpSubmatchNameIndex])
			// fmt.Printf("submatchSlice[%v] = %v\n", GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex, submatchSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex])
			submatchIndexSlice := GoFunctionDeclarationScopeBeginRegexp.FindStringSubmatchIndex(trimSpaceString)
			// fmt.Printf("submatchIndexSlice = %v\n", submatchIndexSlice)
			// fmt.Printf("submatchIndexSlice[%v] = %v, %v\n", GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2, submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2], string(trimSpaceString[submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2]]))
			paramsScopeBeginRune := rune(trimSpaceString[submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2]])
			paramsScopeEndRune := utility.GetAnotherPunctuationMark(paramsScopeBeginRune)
			paramsScopeBeginIndex := submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2+1]
			paramsScopeLength := utility.CalculatePunctuationMarksContentLength(
				trimSpaceString[submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2+1]:],
				paramsScopeBeginRune,
				paramsScopeEndRune,
				utility.InvalidScopePunctuationMarkMap,
			)
			paramsScopeEndIndex := paramsScopeBeginIndex + paramsScopeLength + 1
			var paramsScopeContent string
			if paramsScopeLength > 0 {
				paramsScopeContent = trimSpaceString[paramsScopeBeginIndex:paramsScopeEndIndex]
			}
			fmt.Printf("paramsScopeContent = |%v|\n", paramsScopeContent)
			// returns
			var returnsScopeContent string
			if len(trimSpaceString[paramsScopeEndIndex+1:]) > 0 {
				returnsScopeContent = strings.TrimSpace(trimSpaceString[paramsScopeEndIndex+1:])
			}
			fmt.Printf("returnsScopeContent = |%v|\n", returnsScopeContent)
			fileInterfaceDeclarationMap[interfaceName].FunctionDeclarationMap[functionName] = &GoFunctionDeclaration{
				FunctionSignature: functionName,
				// TypeDeclaration:   ExtractGoVariableTypeDeclaration(submatchSlice[GoVariableDeclarationRegexpSubmatchTypeIndex]),
			}
		}
	}

	return nil
}
