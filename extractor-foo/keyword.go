package extractorfoo

const (
	Keyword_func      string = "func"
	Keyword_struct    string = "struct"
	Keyword_interface string = "interface"
)

var ScopeKeywordMap = map[string]struct{}{
	Keyword_func:      {},
	Keyword_struct:    {},
	Keyword_interface: {},
}

// var BracketScopeKeywordMap = map[string]struct{}{}
// var CurlyBracketScopeKeywordMap = map[string]struct{}{}
// var SquareBracketScopeKeywordMap = map[string]struct{}{}

func IsGolangScopeKeyword(k string) bool {
	_, has := ScopeKeywordMap[k]
	return has
}
