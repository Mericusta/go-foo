package extractorfoo

import "regexp"

type GoVariableDefinition struct {
	VariableSignature string
	TypeDeclaration   *GoTypeDeclaration
}

var (
	// GO_VARIABLE_DECLARATION in func declaration or struct member declaration
	// in func declaration: func([param variable declaration] [param type declaration])
	// in struct member declaration: [member variable declaration] [member type declaration]
	GO_VARIABLE_DECLARATION_EXPRESSION           string = `(?P<NAME>\w+)\s+(?P<TYPE>\S+)\s*`
	GoVariableDeclarationRegexp                         = regexp.MustCompile(GO_VARIABLE_DECLARATION_EXPRESSION)
	GoVariableDeclarationRegexpSubmatchNameIndex        = GoVariableDeclarationRegexp.SubexpIndex("NAME")
	GoVariableDeclarationRegexpSubmatchTypeIndex        = GoVariableDeclarationRegexp.SubexpIndex("TYPE")
)
