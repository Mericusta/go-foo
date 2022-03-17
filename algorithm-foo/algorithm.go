package algorithmfoo

import (
	"fmt"
	"strings"
	"time"
)

func ConvertCamelCase2SnakeCaseWithPhraseTest() {
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("ElfAwakeExample", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("AwakeElfExample", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("AwakeExampleElf", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("ELFAwakeExample", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("AwakeELFExample", map[string]struct{}{"elf": {}}))
	fmt.Printf("%v\n", ConvertCamelCase2SnakeCaseWithPhrase("AwakeExampleELF", map[string]struct{}{"elf": {}}))

}

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

// CalculateYearsOld 根据出生时间戳计算当前年龄
func CalculateYearsOld(birthTimstamp int64) int {
	birthTime := time.Unix(birthTimstamp, 0)
	nowTime := time.Now()
	if nowTime.Month() < birthTime.Month() || (nowTime.Month() == birthTime.Month() && nowTime.Day() < birthTime.Day()) {
		return nowTime.Year() - birthTime.Year() - 1
	}
	return nowTime.Year() - birthTime.Year()
}

func CalculateYearsOldTest() {
	// 1995.6.5 15:00
	birthTimestamp := 802335600
	fmt.Printf("now from 1995.6.5 15:00, %v\n", CalculateYearsOld(int64(birthTimestamp)))

	// 1995.3.5 15:00
	birthTimestamp = 794386800
	fmt.Printf("now from 1995.3.5 15:00, %v\n", CalculateYearsOld(int64(birthTimestamp)))

	// 1995.2.5 15:00
	birthTimestamp = 791967600
	fmt.Printf("now from 1995.2.5 15:00, %v\n", CalculateYearsOld(int64(birthTimestamp)))
}
