package utility

import "strings"

var (
	punctuationMarkMap map[rune]rune = map[rune]rune{
		'(': ')', ')': '(',
		'{': '}', '}': '{',
		'[': ']', ']': '[',
	}
	PunctuationMarkLeftDoubleQuotes   rune = '"'
	PunctuationMarkRightDoubleQuotes  rune = '"'
	PunctuationMarkLeftInverseQuotes  rune = '`'
	PunctuationMarkRightInverseQuotes rune = '`'
	PunctuationMarkLeftSingleQuotes   rune = '\''
	PunctuationMarkRightSingleQuotes  rune = '\''
	InvalidScopePunctuationMarkMap         = map[rune]rune{
		PunctuationMarkLeftDoubleQuotes:  PunctuationMarkRightDoubleQuotes,
		PunctuationMarkLeftInverseQuotes: PunctuationMarkRightInverseQuotes,
		PunctuationMarkLeftSingleQuotes:  PunctuationMarkRightSingleQuotes,
	}
)

// GetAnotherPunctuationMark 获取标点符号的另一对
func GetAnotherPunctuationMark(r rune) rune {
	if markRune, hasMark := punctuationMarkMap[r]; hasMark {
		return markRune
	}
	return ' '
}

// CalculatePunctuationMarksContentLength 计算成对符号的内容长度
// @contentAfterLeftPunctuationMark 待计算的字符串，不包括起始符号
// @leftPunctuationMark 符号左边界字符
// @rightPunctuationMark 符号右边界字符
// @invalidScopePunctuationMarkMap 排除计算的边界符号
// @return
func CalculatePunctuationMarksContentLength(contentAfterLeftPunctuationMark string, leftPunctuationMark, rightPunctuationMark rune, invalidScopePunctuationMarkMap map[rune]rune) int {
	length := 0
	leftCount := 1
	rightCount := 0
	isValid := true
	var invalidScopePunctuationMark rune = -1
	strings.IndexFunc(contentAfterLeftPunctuationMark, func(r rune) bool {
		length++

		// end invalid scope
		if !isValid && r == invalidScopePunctuationMark {
			isValid = true
			invalidScopePunctuationMark = -1
			return false
		}

		// in invalid scope
		if !isValid {
			return false
		}

		// begin invalid scope
		if punctuationMark, isInvalidScopePunctuationMark := invalidScopePunctuationMarkMap[r]; isValid && isInvalidScopePunctuationMark {
			isValid = false
			invalidScopePunctuationMark = punctuationMark
			return false
		}

		// out invalid scope
		if r == leftPunctuationMark {
			leftCount++
		} else if r == rightPunctuationMark {
			rightCount++
		}

		if leftCount == rightCount {
			length-- // cut right punctuation mark len
		}
		return leftCount == rightCount
	})
	return length
}
