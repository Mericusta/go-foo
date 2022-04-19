package extractorfoo

import "regexp"

type GoFunctionDeclaration struct {
	FunctionSignature string
	ParamsList        []*GoVariableDefinition
}

func (d *GoFunctionDeclaration) MakeUp() string {
	return ""
}

var (
	GO_FUNCTION_DECLARATION_SCOPE_BEGIN_EXPRESSION                     string = `(?P<NAME>\w+)\s*(?P<PARAMS_SCOPE_BEGIN>\()`
	GoFunctionDeclarationScopeBeginRegexp                                     = regexp.MustCompile(GO_FUNCTION_DECLARATION_SCOPE_BEGIN_EXPRESSION)
	GoFunctionDeclarationScopeBeginRegexpSubmatchNameIndex                    = GoFunctionDeclarationScopeBeginRegexp.SubexpIndex("NAME")
	GoFunctionDeclarationScopeBeginRegexpSubmatchParamsScopeBeginIndex        = GoFunctionDeclarationScopeBeginRegexp.SubexpIndex("PARAMS_SCOPE_BEGIN")
)
