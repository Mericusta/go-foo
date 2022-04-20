package extractorfoo

import (
	"fmt"
	"go-foo/utility"
	"regexp"
	"strings"
)

type GoFunctionDeclaration struct {
	FunctionSignature string
	This              *GoVariableDefinition
	ParamsList        []*GoVariableDefinition
	ReturnList        []*GoVariableDefinition
	BodyIndexSlice    []int
}

func (d *GoFunctionDeclaration) MakeUp() string {
	return ""
}

var (
	GO_FUNCTION_DECLARATION_SCOPE_BEGIN_EXPRESSION                     string = `\nfunc\s+(\((?P<THIS>\w+)\s+(?P<THIS_TYPE>(\*)?\w+)\))?\s*(?P<NAME>\w+)\s*(?P<PARAMS_SCOPE_BEGIN>\()`
	GoFunctionDeclarationScopeBeginRegexp                                     = regexp.MustCompile(GO_FUNCTION_DECLARATION_SCOPE_BEGIN_EXPRESSION)
	GoFunctionDeclarationScopeBeginRegexpSubmatchThisIndex                    = GoFunctionDeclarationScopeBeginRegexp.SubexpIndex("THIS")
	GoFunctionDeclarationScopeBeginRegexpSubmatchThisTypeIndex                = GoFunctionDeclarationScopeBeginRegexp.SubexpIndex("THIS_TYPE")
	GoFunctionDeclarationScopeBeginRegexpSubmatchNameIndex                    = GoFunctionDeclarationScopeBeginRegexp.SubexpIndex("NAME")
	GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex        = GoFunctionDeclarationScopeBeginRegexp.SubexpIndex("PARAMS_SCOPE_BEGIN")
)

func ExtractGoFileFunctionDeclaration(content []byte) map[string]*GoFunctionDeclaration {
	if len(content) == 0 {
		return nil
	}

	if GoFunctionDeclarationScopeBeginRegexpSubmatchThisIndex == -1 || GoFunctionDeclarationScopeBeginRegexpSubmatchThisTypeIndex == -1 || GoFunctionDeclarationScopeBeginRegexpSubmatchNameIndex == -1 || GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex == -1 {
		panic("sub match index is -1")
	}

	functionDeclarationMap := make(map[string]*GoFunctionDeclaration)
	for _, functionDeclarationScopeBeginSubmatchIndexSlice := range GoFunctionDeclarationScopeBeginRegexp.FindAllSubmatchIndex(content, -1) {
		fmt.Printf("function declaration scope begin = |%v|\n", strings.TrimSpace(string(content[functionDeclarationScopeBeginSubmatchIndexSlice[0]:functionDeclarationScopeBeginSubmatchIndexSlice[1]])))

		// signature
		functionName := strings.TrimSpace(string(content[functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchNameIndex*2]:functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchNameIndex*2+1]]))
		fmt.Printf("function name = |%v|\n", functionName)

		// this
		var thisDeclaration *GoVariableDefinition
		if functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchThisIndex*2] != -1 &&
			functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchThisIndex*2+1] != -1 &&
			functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchThisTypeIndex*2] != -1 &&
			functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchThisTypeIndex*2+1] != -1 {
			thisSignature := strings.TrimSpace(string(content[functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchThisIndex*2]:functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchThisIndex*2+1]]))
			thisTypeContent := strings.TrimSpace(string(content[functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchThisTypeIndex*2]:functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchThisTypeIndex*2+1]]))
			thisDeclaration = &GoVariableDefinition{
				VariableSignature: thisSignature,
				TypeDeclaration:   ExtractGoVariableTypeDeclaration(thisTypeContent),
			}
			fmt.Printf("function this = |%v|\n", thisDeclaration.VariableSignature)
			fmt.Printf("function this type = |%v|\n", thisDeclaration.TypeDeclaration.MakeUp())
		}

		// params list
		fmt.Printf("function params scope begin = |%v|\n", strings.TrimSpace(string(content[functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2]:functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2+1]])))

		ExtractorFunctionParamsList(content, functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2])

		// // params
		// // fmt.Printf("submatchSlice = %v\n", submatchSlice)
		// // fmt.Printf("submatchSlice[%v] = %v\n", GoStructRegexpSubmatchNameIndex, submatchSlice[GoStructRegexpSubmatchNameIndex])
		// // fmt.Printf("submatchSlice[%v] = %v\n", GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex, submatchSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex])
		// submatchIndexSlice := GoFunctionDeclarationScopeBeginRegexp.FindStringSubmatchIndex(trimSpaceString)
		// // fmt.Printf("submatchIndexSlice = %v\n", submatchIndexSlice)
		// // fmt.Printf("submatchIndexSlice[%v] = %v, %v\n", GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2, submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2], string(trimSpaceString[submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2]]))
		// paramsScopeBeginRune := rune(trimSpaceString[submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2]])
		// paramsScopeEndRune := utility.GetAnotherPunctuationMark(paramsScopeBeginRune)
		// paramsScopeBeginIndex := submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2+1]
		// paramsScopeLength := utility.CalculatePunctuationMarksContentLength(
		// 	trimSpaceString[submatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2+1]:],
		// 	paramsScopeBeginRune,
		// 	paramsScopeEndRune,
		// 	utility.InvalidScopePunctuationMarkMap,
		// )
		// paramsScopeEndIndex := paramsScopeBeginIndex + paramsScopeLength + 1
		// var paramsScopeContent string
		// if paramsScopeLength > 0 {
		// 	paramsScopeContent = trimSpaceString[paramsScopeBeginIndex:paramsScopeEndIndex]
		// }
		// fmt.Printf("paramsScopeContent = |%v|\n", paramsScopeContent)

		// // returns
		// var returnsScopeContent string
		// if paramsScopeEndIndex+1 < len(trimSpaceString) && len(trimSpaceString[paramsScopeEndIndex+1:]) > 0 {
		// 	returnsScopeContent = strings.TrimSpace(trimSpaceString[paramsScopeEndIndex+1:])
		// }
		// fmt.Printf("returnsScopeContent = |%v|\n", returnsScopeContent)

		fmt.Println()

		// utility.CalculatePunctuationMarksContentLength(string(content[]))

		functionDeclarationMap[functionName] = &GoFunctionDeclaration{
			FunctionSignature: functionName,
			This:              thisDeclaration,
		}
	}

	return functionDeclarationMap
}

func ExtractorFunctionParamsList(content []byte, scopeBeginIndex int) []*GoVariableDefinition {
	paramsListContent := utility.GetScopeContentBetweenPunctuationMarks(content, scopeBeginIndex)
	fmt.Printf("paramsListContent = |%v|\n", string(paramsListContent))

	splitContent := utility.RecursiveSplitUnderSameDeepPunctuationMarksContent(string(paramsListContent), utility.GetLeftPunctuationMarkList(), ",")

	var sameTypeParamSlice, paramsSlice []*GoVariableDefinition
	for _, content := range splitContent.ContentList {
		// fmt.Printf("param content = |%v|\n", strings.TrimSpace(content))
		if len(content) == 0 {
			panic("param content is empty")
		}
		paramDeclaration := &GoVariableDefinition{}
		paramContentSubmatchSlice := GoVariableDeclarationRegexp.FindStringSubmatch(strings.TrimSpace(content))
		if len(paramContentSubmatchSlice) == 0 {
			paramDeclaration.VariableSignature = strings.TrimSpace(content)
			sameTypeParamSlice = append(sameTypeParamSlice, paramDeclaration)
		} else {
			paramDeclaration.VariableSignature = paramContentSubmatchSlice[GoVariableDeclarationRegexpSubmatchNameIndex]
			paramDeclaration.TypeDeclaration = ExtractGoVariableTypeDeclaration(paramContentSubmatchSlice[GoVariableDeclarationRegexpSubmatchTypeIndex])
			for _, sameTypeParam := range sameTypeParamSlice {
				sameTypeParam.TypeDeclaration = paramDeclaration.TypeDeclaration
			}
			sameTypeParamSlice = nil
		}
		paramsSlice = append(paramsSlice, paramDeclaration)
	}

	// for index, paramDeclaration := range paramsSlice {
	// 	fmt.Printf("%v param: %v\n", index, paramDeclaration.VariableSignature)
	// 	fmt.Printf("%v param type: %v\n", index, paramDeclaration.TypeDeclaration.MakeUp())
	// }
	return paramsSlice
}

func ExtractorFunctionReturnList() []*GoVariableDefinition {
	return nil
}
