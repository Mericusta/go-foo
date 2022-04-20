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
	ReturnList        []*GoTypeDeclaration // not support named return
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

		// if functionName != "ExtractMemberFunction" {
		// 	continue
		// }

		// this scope
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

		// params scope
		paramsScopeBeginRuneIndex := functionDeclarationScopeBeginSubmatchIndexSlice[GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex*2] // '(' index
		paramsScopeBeginRune := rune(content[paramsScopeBeginRuneIndex])                                                                                   // '('
		paramsScopeEndRune := utility.GetAnotherPunctuationMark(paramsScopeBeginRune)                                                                      // ')'
		paramsScopeLength := utility.CalculatePunctuationMarksContentLength(
			string(content[paramsScopeBeginRuneIndex+1:]),
			paramsScopeBeginRune, paramsScopeEndRune, utility.InvalidScopePunctuationMarkMap,
		)
		paramsScopeEndRuneIndex := paramsScopeBeginRuneIndex + 1 + paramsScopeLength        // ')' index
		paramsListContent := content[paramsScopeBeginRuneIndex+1 : paramsScopeEndRuneIndex] // between '(' and ')'
		fmt.Printf("paramsListContent = |%v|\n", string(paramsListContent))
		paramsList := ExtractorFunctionParamsList(paramsListContent)

		// returns scope
		returnsScopeBeginRuneIndex := paramsScopeEndRuneIndex + 1 // after params scope end rune index
		returnsScopeEndRuneIndex := returnsScopeBeginRuneIndex    // before body scope begin rune index
		bodyScopeBeginRuneIndex := -1                             // '{' index
		keywordStack := []string{Keyword_func}
		word := make([]byte, 0, 16)
		for rIndex := 0; rIndex != len(content[paramsScopeEndRuneIndex:]); rIndex++ {
			contentIndex := paramsScopeEndRuneIndex + rIndex
			r := rune(content[contentIndex])
			switch {
			case utility.IsCharacter(r):
				word = append(word, byte(r))
			default:
				if len(word) > 0 {
					if IsGolangScopeKeyword(string(word)) {
						keywordStack = append(keywordStack, string(word))
					}
					word = make([]byte, 0, 16)
				}
				stackLength := len(keywordStack)
				switch {
				case utility.IsSpaceRune(r):
				case r == '(':
					scopeLength := utility.CalculatePunctuationMarksContentLength(
						string(content[contentIndex+1:]),
						'(', ')', utility.InvalidScopePunctuationMarkMap,
					)
					if stackLength == 1 && keywordStack[0] == Keyword_func {
						returnsScopeBeginRuneIndex = contentIndex
					} else {
						if keywordStack[stackLength-1] == Keyword_func {
							if stackLength-2 >= 0 {
								keywordStack = keywordStack[0 : stackLength-1]
							} else {
								panic("stack length error")
							}
						}
					}
					// func (int) -> func (int)
					//      |                 |
					rIndex += (1 + scopeLength)
				case r == '[':
					scopeLength := utility.CalculatePunctuationMarksContentLength(
						string(content[contentIndex+1:]),
						'[', ']', utility.InvalidScopePunctuationMarkMap,
					)
					// map[int]int -> map[int]int
					//    |                  |
					rIndex += (1 + scopeLength)
				case r == '{':
					if stackLength == 1 && keywordStack[0] == Keyword_func {
						if content[returnsScopeBeginRuneIndex] == '(' {
							returnsScopeEndRuneIndex = contentIndex - 1
							for content[returnsScopeEndRuneIndex] != ')' {
								returnsScopeEndRuneIndex--
							}
						} else {
							returnsScopeEndRuneIndex = contentIndex
						}
						bodyScopeBeginRuneIndex = contentIndex
						goto SEARCH_END
					} else {
						scopeLength := utility.CalculatePunctuationMarksContentLength(
							string(content[contentIndex+1:]),
							'{', '}', utility.InvalidScopePunctuationMarkMap,
						)
						if keywordStack[stackLength-1] == Keyword_func {
							panic("syntax error")
						}
						// interface{} -> interface{}
						//          |               |
						// struct{ v interface{} } -> struct{ v interface{} }
						//       |                                          |
						rIndex += (1 + scopeLength)
						if stackLength-2 >= 0 {
							keywordStack = keywordStack[0 : stackLength-1]
						} else {
							panic("stack length error")
						}
					}
				}
			}
		}
	SEARCH_END:
		fmt.Printf("function returns scope = |%v|\n", string(content[returnsScopeBeginRuneIndex+1:returnsScopeEndRuneIndex]))
		returnList := ExtractorFunctionReturnList(content[returnsScopeBeginRuneIndex+1 : returnsScopeEndRuneIndex])

		// body scope
		if bodyScopeBeginRuneIndex < 0 {
			panic("body scope begin rune index is -1")
		}
		bodyLength := utility.CalculatePunctuationMarksContentLength(
			string(content[bodyScopeBeginRuneIndex+1:]),
			'{', '}', utility.InvalidScopePunctuationMarkMap,
		)
		if bodyLength < 0 {
			panic("function body length is -1")
		}
		bodyScopeEndRuneIndex := bodyScopeBeginRuneIndex + 1 + bodyLength
		fmt.Printf("function body scope = |%v|\n", string(content[bodyScopeBeginRuneIndex+1:bodyScopeEndRuneIndex]))
		fmt.Printf("body begin rune %v\n", string(content[bodyScopeBeginRuneIndex]))
		fmt.Printf("body end rune %v\n", string(content[bodyScopeEndRuneIndex]))
		fmt.Println()

		functionDeclarationMap[functionName] = &GoFunctionDeclaration{
			FunctionSignature: functionName,
			This:              thisDeclaration,
			ParamsList:        paramsList,
			ReturnList:        returnList,
			BodyIndexSlice:    []int{bodyScopeBeginRuneIndex, bodyScopeEndRuneIndex},
		}
	}

	return functionDeclarationMap
}

func ExtractorFunctionParamsList(content []byte) []*GoVariableDefinition {
	splitContent := utility.RecursiveSplitUnderSameDeepPunctuationMarksContent(string(content), utility.GetLeftPunctuationMarkList(), ",")
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

	{
		for index, paramDeclaration := range paramsSlice {
			fmt.Printf("%v param: %v\n", index, paramDeclaration.VariableSignature)
			fmt.Printf("%v param type: %v\n", index, paramDeclaration.TypeDeclaration.MakeUp())
		}
	}
	return paramsSlice
}

func ExtractorFunctionReturnList(content []byte) []*GoTypeDeclaration {
	splitContent := utility.RecursiveSplitUnderSameDeepPunctuationMarksContent(string(content), utility.GetLeftPunctuationMarkList(), ",")
	returnTypeDeclarationSlice := make([]*GoTypeDeclaration, 0)
	for _, content := range splitContent.ContentList {
		fmt.Printf("return content = |%v|\n", strings.TrimSpace(content))
		if len(content) == 0 {
			panic("return content is empty")
		}
		typeDeclaration := ExtractGoVariableTypeDeclaration(strings.TrimSpace(content))
		returnTypeDeclarationSlice = append(returnTypeDeclarationSlice, typeDeclaration)
		fmt.Printf("typeDeclaration = %v\n", typeDeclaration.MakeUp())
		typeDeclaration.Traversal(0)
	}
	return returnTypeDeclarationSlice
}
