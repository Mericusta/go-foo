package algorithmfoo

import "fmt"

// tri-color marking

type objectNode struct {
	no           string
	parentNode   *objectNode
	subNodeSlice []*objectNode
}

// var AllNodeMap map[string]*Node
var rootNodeMap map[string]*objectNode
var nonRootNodeMap map[string]*objectNode
var whiteCollection map[string]*objectNode
var greyCollection map[string]*objectNode
var blackCollection map[string]*objectNode

// simulate Go GC algorithm in v1.5

func GoGarbageCollectionTriColorMarking() {
	step0InitNodeTree()
	step1ContributeCollection()
	step2MarkAllNodesWhite()
	step3ScanRootNodeSliceAndMarkGrey()
	step4ScanGreyCollectionAndMarkBlack()
	step5SweepWhiteCollection()
}

func step0InitNodeTree() {
	ANode := &objectNode{no: "A"}
	BNode := &objectNode{no: "B"}
	CNode := &objectNode{no: "C"}
	DNode := &objectNode{no: "D"}
	ENode := &objectNode{no: "E"}
	FNode := &objectNode{no: "F"}
	GNode := &objectNode{no: "G"}
	HNode := &objectNode{no: "H"}

	RootNode1 := &objectNode{no: "Root1"}
	RootNode2 := &objectNode{no: "Root2"}
	RootNode3 := &objectNode{no: "Root3"}

	rootNodeMap = map[string]*objectNode{
		RootNode1.no: RootNode1, RootNode2.no: RootNode2, RootNode3.no: RootNode3,
	}

	nonRootNodeMap = map[string]*objectNode{
		ANode.no: ANode, BNode.no: BNode, CNode.no: CNode, DNode.no: DNode,
		ENode.no: ENode, FNode.no: FNode, GNode.no: GNode, HNode.no: HNode,
	}

	RootNode1.subNodeSlice = append(RootNode1.subNodeSlice, ANode)
	ANode.parentNode = RootNode1
	ANode.subNodeSlice = append(ANode.subNodeSlice, BNode)
	BNode.parentNode = ANode
	ANode.subNodeSlice = append(ANode.subNodeSlice, CNode)
	CNode.parentNode = ANode
	ANode.subNodeSlice = append(ANode.subNodeSlice, DNode)
	DNode.parentNode = ANode
	// circular reference
	DNode.subNodeSlice = append(DNode.subNodeSlice, ANode)

	RootNode3.subNodeSlice = append(RootNode3.subNodeSlice, FNode)
	FNode.parentNode = RootNode3

	GNode.subNodeSlice = append(GNode.subNodeSlice, HNode)
	HNode.parentNode = GNode

	// Root1 Root2 Root3 ...
	//   |           |
	//   A           F   E   G
	//  /|\                  |
	// B C D                 H
}

func step1ContributeCollection() {
	whiteCollection, greyCollection, blackCollection = make(map[string]*objectNode), make(map[string]*objectNode), make(map[string]*objectNode)
}

// step2MarkAllNodesWhite
// A...H mark White
func step2MarkAllNodesWhite() {
	for no, node := range nonRootNodeMap {
		whiteCollection[no] = node
	}
}

// step3ScanRootNodeSliceAndMarkGrey
// A F mark Grey
// A F remove White
func step3ScanRootNodeSliceAndMarkGrey() {
	for _, node := range rootNodeMap {
		for _, subNode := range node.subNodeSlice {
			greyCollection[subNode.no] = subNode
			delete(whiteCollection, subNode.no)
		}
	}
}

// Step4ScanGreyCollection
// B C D mark Grey, B C D remove White, A F mark Black Collection
// B C D mark Black Collection
func step4ScanGreyCollectionAndMarkBlack() {
	for len(greyCollection) > 0 {
		newGreyCollection := make(map[string]*objectNode)
		for no, node := range greyCollection {
			blackCollection[no] = node
			for _, subNode := range node.subNodeSlice {
				if _, isBlack := blackCollection[subNode.no]; !isBlack {
					newGreyCollection[subNode.no] = subNode
					delete(whiteCollection, subNode.no)
				}
			}
		}
		greyCollection = newGreyCollection
	}
}

// step5SweepWhiteCollection
// Sweep Node: E G H
func step5SweepWhiteCollection() {
	for _, node := range whiteCollection {
		fmt.Printf("sweep node %v\n", node.no)
	}
}

// ----------------------------------------------------------------
