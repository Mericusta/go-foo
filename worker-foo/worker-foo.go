package workerfoo

import (
	"fmt"
	"os"
)

//
func ExtractGoFileStructDeclarationTest() {
	testFilePath := "resources/extract-foo.go"
	f, e := os.Open(testFilePath)
	if f == nil || e != nil {
		panic(e)
	}

	fileStructDeclarationMap := ExtractGoFileStructDeclaration(f)
	for _, structInfo := range fileStructDeclarationMap {
		fmt.Printf("- %v\n", structInfo.Name)
		for _, structMemberDeclaration := range structInfo.MemberDeclarationMap {
			fmt.Printf("\t- %v\n", structMemberDeclaration.MemberSignature)
			structMemberDeclaration.TypeDeclaration.Traversal(2)
		}
		fmt.Printf("- construct: %v\n", structInfo.Construct())
	}
}
