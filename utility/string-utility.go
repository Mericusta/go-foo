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

func GetLeftPunctuationMarkList() []rune {
	leftPunctuationMarkList := make([]rune, 0, 8)
	for leftPunctuationMark := range punctuationMarkMap {
		leftPunctuationMarkList = append(leftPunctuationMarkList, leftPunctuationMark)
	}
	return leftPunctuationMarkList
}

// CalculatePunctuationMarksContentLength 计算成对符号的内容长度，去除结束符号
// @contentAfterLeftPunctuationMark       待计算的字符串，不包括起始符号
// @leftPunctuationMark                   符号左边界字符
// @rightPunctuationMark                  符号右边界字符
// @invalidScopePunctuationMarkMap        排除计算的边界符号
// @return                                不包含左右边界的内容的长度
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

		return leftCount == rightCount
	})
	return length - 1 // // cut right punctuation mark len
}

// GetScopeContentBetweenPunctuationMarks 获取成对标点符号的内容
// @content                               待查找的内容
// @scopeBeginIndex                       左边界起始符号的下标
// @return                                不包含左右边界的内容
func GetScopeContentBetweenPunctuationMarks(content []byte, scopeBeginIndex int) []byte {
	scopeBeginRune := rune(content[scopeBeginIndex])
	scopeEndRune := GetAnotherPunctuationMark(scopeBeginRune)
	scopeContentLength := CalculatePunctuationMarksContentLength(
		string(content[scopeBeginIndex+1:]),
		scopeBeginRune, scopeEndRune,
		InvalidScopePunctuationMarkMap,
	)
	return content[scopeBeginIndex+1 : scopeBeginIndex+1+scopeContentLength]
}

// SplitContent 划分内容节点
type SplitContent struct {
	ContentList         []string
	SubSplitContentList []*SplitContent
}

// RecursiveSplitUnderSameDeepPunctuationMarksContent 相同深度的成对标点符号下的内容划分
// @content                 待分析的字符串，不包含最顶层左右边界
// @punctuationLeftMarkList 指定成对标点符号的左边界，一般指 (, [, {, ", ', `
// @splitter                指定分隔符
// @return
func RecursiveSplitUnderSameDeepPunctuationMarksContent(content string, leftPunctuationMarkList []rune, splitter string) *SplitContent {
	if punctuationContentNode := RecursiveTraitMultiPunctuationMarksContent(content, &PunctuationMarkInfo{
		PunctuationMark: 0,
		Index:           -1,
	}, &PunctuationMarkInfo{
		PunctuationMark: 0,
		Index:           len(content),
	}, leftPunctuationMarkList, 1, 0); punctuationContentNode != nil {
		return splitUnderSameDeepPunctuationMarksContent(punctuationContentNode, splitter, 0, 0)
	}
	return nil
}

// PunctuationContent 成对标点符号的内容节点
type PunctuationContent struct {
	Content                   string
	LeftPunctuationMark       *PunctuationMarkInfo
	RightPunctuationMark      *PunctuationMarkInfo
	SubPunctuationContentList []*PunctuationContent
}

type PunctuationMarkInfo struct {
	PunctuationMark rune
	Index           int
}

// // TraitMultiPunctuationMarksContent 混合成对标点符号的内容分类提取
// func TraitMultiPunctuationMarksContent(content string, leftPunctuationMarkList []rune, maxDeep int) *PunctuationContent {
// 	return
// }

// RecursiveTraitMultiPunctuationMarksContent 混合成对标点符号的内容分类提取
// @content 待处理内容
// @leftPunctuationMarkInfo 根节点的左标点符号
// @rightPunctuationMarkInfo 根节点的右标点符号
// @scopeLeftPunctuationMarkList 所有作为划分区域的左标点符号
// @maxDeep 待处理的最大深度
// @deep 当前深度
// @return 根节点
func RecursiveTraitMultiPunctuationMarksContent(content string, leftPunctuationMarkInfo, rightPunctuationMarkInfo *PunctuationMarkInfo, scopeLeftPunctuationMarkList []rune, maxDeep, deep int) *PunctuationContent {
	punctuationContent := &PunctuationContent{
		Content:                   content,
		LeftPunctuationMark:       leftPunctuationMarkInfo,
		RightPunctuationMark:      rightPunctuationMarkInfo,
		SubPunctuationContentList: make([]*PunctuationContent, 0),
	}

	passLeftLength := 0
	for len(content) != 0 && deep != maxDeep {
		var leftPunctuationMark rune
		var rightPunctuationMark rune
		leftPunctuationMarkIndex := len(content) - 1

		for _, toSearchLeftPunctuationMark := range scopeLeftPunctuationMarkList {
			toSearchLeftPunctuationMarkIndex := strings.IndexRune(content, toSearchLeftPunctuationMark)
			if toSearchLeftPunctuationMarkIndex != -1 && toSearchLeftPunctuationMarkIndex < leftPunctuationMarkIndex {
				leftPunctuationMarkIndex = toSearchLeftPunctuationMarkIndex
				leftPunctuationMark = toSearchLeftPunctuationMark
			}
		}
		// fmt.Printf("relative leftPunctuationMarkIndex = %v, leftPunctuationMark = %v\n", leftPunctuationMarkIndex, string(rune(leftPunctuationMark)))

		rightPunctuationMark = GetAnotherPunctuationMark(leftPunctuationMark)
		if leftPunctuationMark == 0 || rightPunctuationMark == 0 || leftPunctuationMarkIndex == len(content)-1 {
			break
		}

		afterLeftPunctuationMarkContentIndex := leftPunctuationMarkIndex + 1

		// fmt.Printf("pass CalculatePunctuationMarksContentLength = |%v|\n", content[afterLeftPunctuationMarkContentIndex:])
		length := CalculatePunctuationMarksContentLength(content[afterLeftPunctuationMarkContentIndex:], leftPunctuationMark, rightPunctuationMark, InvalidScopePunctuationMarkMap)

		// fmt.Printf("after CalculatePunctuationMarksContentLength, length = %v\n", length)

		rightPunctuationMarkIndex := leftPunctuationMarkIndex + length + 1
		if rightPunctuationMarkIndex >= len(content) {
			// fmt.Printf("rightPunctuationMarkIndex %v >= len(content) %v\n", rightPunctuationMarkIndex, len(content))
			break
		}

		// fmt.Printf("relative rightPunctuationMarkIndex = %v, rightPunctuationMark = %v\n", rightPunctuationMarkIndex, string(rune(rightPunctuationMark)))
		// fmt.Printf("pass content = |%v|\n", content[leftPunctuationMarkIndex+1:rightPunctuationMarkIndex])

		subPunctuationContent := RecursiveTraitMultiPunctuationMarksContent(content[leftPunctuationMarkIndex+1:rightPunctuationMarkIndex], &PunctuationMarkInfo{
			PunctuationMark: leftPunctuationMark,
			Index:           leftPunctuationMarkInfo.Index + 1 + passLeftLength + leftPunctuationMarkIndex,
		}, &PunctuationMarkInfo{
			PunctuationMark: rightPunctuationMark,
			Index:           leftPunctuationMarkInfo.Index + 1 + passLeftLength + rightPunctuationMarkIndex,
		}, scopeLeftPunctuationMarkList, maxDeep, deep+1)
		if subPunctuationContent != nil {
			punctuationContent.SubPunctuationContentList = append(punctuationContent.SubPunctuationContentList, subPunctuationContent)
		}

		// fmt.Printf("update content to |%v|\n", content[rightPunctuationMarkIndex+1:])
		content = content[rightPunctuationMarkIndex+1:]
		// fmt.Printf("update pass left from %v to %v\n", passLeftLength, passLeftLength+rightPunctuationMarkIndex+1)
		passLeftLength += rightPunctuationMarkIndex + 1
		// fmt.Println("--------------------------------")
	}

	return punctuationContent
}

// splitUnderSameDeepPunctuationMarksContent 相同深度的成对标点符号下的内容划分的递归算法
// @punctuationContentNode 成对标点符号的内容根节点，注意：必须是根节点，不能是某个子节点，节点深度 >= 2，分析结果中深度大于 2 的数据不正确
// @splitter 指定分隔符
// @maxDeep 递归最大深度
// @deep 当前深度
func splitUnderSameDeepPunctuationMarksContent(punctuationContentNode *PunctuationContent, splitter string, maxDeep, deep int) *SplitContent {
	splitContentNode := &SplitContent{
		ContentList:         make([]string, 0),
		SubSplitContentList: make([]*SplitContent, 0),
	}

	var offset int
	var leftIndex int
	cycle := 0
	maxCycle := len(strings.Split(punctuationContentNode.Content, splitter))
	for cycle != maxCycle {
		cycle++
		length := strings.Index(punctuationContentNode.Content[leftIndex+offset:], splitter)
		if length == -1 {
			splitContentNode.ContentList = append(splitContentNode.ContentList, punctuationContentNode.Content[leftIndex:])
			break
		}
		rightIndex := leftIndex + length + offset
		inner := false
		for _, subNode := range punctuationContentNode.SubPunctuationContentList {
			// Note: 这里用于判断的依据是子节点相对父节点的左 区间符号 的下标
			// Note: 但是节点的 区间符号 数据中记录的下标是相对于根节点的下标 -> 必须是根节点
			// Note: 所以当节点数只有2时，这个下标可以代表相对父节点（根节点）的下标 -> 节点深度 >= 2
			if subNode.LeftPunctuationMark.Index <= rightIndex && rightIndex <= subNode.RightPunctuationMark.Index {
				inner = true
				offset = subNode.RightPunctuationMark.Index - leftIndex + 1
				break
			}
		}
		if inner {
			continue
		}
		splitContentNode.ContentList = append(splitContentNode.ContentList, punctuationContentNode.Content[leftIndex:rightIndex])
		offset = 0
		leftIndex = rightIndex + len(splitter)
	}

	if deep == maxDeep {
		return splitContentNode
	}

	for _, subPunctuationContentNode := range punctuationContentNode.SubPunctuationContentList {
		if len(subPunctuationContentNode.Content) != 0 {
			if subSplitContentNode := splitUnderSameDeepPunctuationMarksContent(subPunctuationContentNode, splitter, maxDeep, deep+1); subSplitContentNode != nil {
				splitContentNode.SubSplitContentList = append(splitContentNode.SubSplitContentList, subSplitContentNode)
			}
		}
	}

	return splitContentNode
}
