package algorithmfoo

import "strings"

// Abc 开头 | 中间 DONE
// ABC 开头 | 中间
// ConvertCamelCase2SnakeCaseWithPhrase 将驼峰命名法转换为蛇形命名法：XxxYyyZzz -> xxx_yyy_zzz
func ConvertCamelCase2SnakeCaseWithPhrase(camelCase string, phraseMap map[string]struct{}) string {
	allPhraseSubString := make(map[string]struct{})
	for phrase := range phraseMap {
		for index := 0; index != len(phrase); index++ {
			allPhraseSubString[phrase[0:index]] = struct{}{}
		}
	}

	builder := strings.Builder{}
	phraseBuilder := strings.Builder{}
	isFirstPhrase := true
	for _, r := range camelCase {
		if 'a' <= r && r <= 'z' {
			phraseBuilder.WriteRune(r)
		} else {
			if phraseBuilder.Len() > 0 {
				if _, isPhrase := phraseMap[phraseBuilder.String()]; isPhrase {
					if isFirstPhrase {
						isFirstPhrase = false
					} else {
						builder.WriteRune('_')
					}
					builder.WriteString(phraseBuilder.String())
					phraseBuilder.Reset()
				} else {
					if _, maybePhrase := allPhraseSubString[phraseBuilder.String()]; !maybePhrase {
						if isFirstPhrase {
							isFirstPhrase = false
						} else {
							builder.WriteRune('_')
						}
						builder.WriteString(phraseBuilder.String())
						phraseBuilder.Reset()
					}
				}
			}
			phraseBuilder.WriteRune(r + 32)
		}
	}
	builder.WriteRune('_')
	builder.WriteString(phraseBuilder.String())
	return builder.String()
}
