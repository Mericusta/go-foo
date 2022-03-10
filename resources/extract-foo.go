package resources

import extractorfoo "go-foo/extractor-foo"

type renameInt int

type ExtractGoStruct struct {
	i   int
	f   float64
	s   string
	sr  []rune
	msb map[string]byte
	ps  *ExtractGoStruct
	ip  *extractorfoo.GoStructMemberDefinition
	// as  struct {
	// 	*ExtractGoStruct
	// 	*extractorfoo.GoStructMemberDefinition
	// }
}

type EmptyExtractGoStruct struct {
}

type OneLineEmptyExtractGoStruct struct{}

var commentAndDivide1 int = 1 / 2
var commentAndDivide2 int = 1 / 2
