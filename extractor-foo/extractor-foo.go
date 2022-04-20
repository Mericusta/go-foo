package extractorfoo

import (
	"fmt"
	"io/ioutil"
	"os"
)

// func ExtractorGoFileTest[]() {
// 	testFilePath := "resources/extract-foo.go"
// 	f, e := os.Open(testFilePath)
// 	if f == nil || e != nil {
// 		panic(e)
// 	}
// }

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
			fmt.Printf("\t- %v\n", structMemberDeclaration.VariableSignature)
			structMemberDeclaration.TypeDeclaration.Traversal(2)
		}
		fmt.Printf("- construct: %v\n", structInfo.Construct())
	}
}

func ExtractGoFileInterfaceDeclarationTest() {
	testFilePath := "resources/extract-foo.go"
	f, e := os.Open(testFilePath)
	if f == nil || e != nil {
		panic(e)
	}

	fileInterfaceDeclarationMap := ExtractGoFileInterfaceDeclaration(f)
	for interfaceName := range fileInterfaceDeclarationMap {
		fmt.Printf("- %v\n", interfaceName)
		// for _, functionDeclaration := range interfaceInfo.FunctionDeclarationMap {
		// 	fmt.Printf("\t- %v\n", functionDeclaration.MakeUp())
		// }
	}
}

func ExtractGoFileFunctionScopeTest() {
	testFilePath := "resources/extract-foo.go"
	GoFmtFile(testFilePath)

	f, e := os.Open(testFilePath)
	if f == nil || e != nil {
		panic(e)
	}

	testFileContent, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	ExtractGoFileFunctionDeclaration(testFileContent)
}
