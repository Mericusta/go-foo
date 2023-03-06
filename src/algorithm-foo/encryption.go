package algorithmfoo

import (
	"fmt"
	"io/ioutil"
	"os"
)

const uint8Max = 255

func encryptionFoo(inputFilePath, outputFilePath, operate string) {
	// var inputFilePath string
	// flag.StringVar(&inputFilePath, "input", "", "specify an input file")

	// var outputFilePath string
	// flag.StringVar(&outputFilePath, "output", "", "specify an output file")

	// var operate string
	// flag.StringVar(&operate, "operate", "encode", "specify an operation, encode or decode, default operation is encode")

	// flag.Parse()

	if inputFilePath == "" {
		fmt.Println("ERROR: input file required")
		return
	}
	if outputFilePath == "" {
		fmt.Println("ERROR: output file required")
		return
	}

	inputFile, osOpenError := os.Open(inputFilePath)
	if osOpenError != nil {
		fmt.Printf("open file %v occurs error, %v\n", inputFile, osOpenError)
		return
	}

	outputFile, osCreateError := os.Create(outputFilePath)
	if osCreateError != nil {
		fmt.Printf("create %v error, %v", outputFilePath, osCreateError)
		return
	}

	operationResult := true

	if operate == "encode" {
		operationResult = encode(inputFile, outputFile)
	} else if operate == "decode" {
		operationResult = decode(inputFile, outputFile)
	} else {
		fmt.Printf("unknown operation %v\n", operate)
		return
	}

	if operationResult {
		fmt.Printf("NOTE: %v successfully\n", operate)
	} else {
		fmt.Printf("NOTE: %v failed\n", operate)
	}
}

func encode(toEncryptFile, encryptedFile *os.File) bool {
	toEncryptByteList, toEncryptFileReadAllError := ioutil.ReadAll(toEncryptFile)
	if toEncryptFileReadAllError != nil {
		fmt.Printf("ERROR: read input file occurs error, %v", toEncryptFileReadAllError)
		return false
	}
	if len(toEncryptByteList) == 0 {
		fmt.Println("ERROR: read input file, null content")
		return false
	}

	encryptedByteList := make([]byte, 0)
	for _, toEncryptByte := range toEncryptByteList {
		encryptedByteList = append(encryptedByteList, toEncryptByte^uint8Max)
	}

	_, encryptedFileWriteError := encryptedFile.Write(encryptedByteList)
	if encryptedFileWriteError != nil {
		fmt.Printf("ERROR: write output file occurs error, %v", encryptedFileWriteError)
		return false
	}

	return true
}

func decode(encryptedFile, decryptedFile *os.File) bool {
	encryptedByteList, encryptedFileReadAllError := ioutil.ReadAll(encryptedFile)
	if encryptedFileReadAllError != nil {
		fmt.Printf("ERROR: read input file occurs error, %v", encryptedFileReadAllError)
		return false
	}
	if len(encryptedByteList) == 0 {
		fmt.Println("ERROR: read input file, null content")
		return false
	}

	decryptByteList := make([]byte, 0)
	for _, encryptedByte := range encryptedByteList {
		decryptByteList = append(decryptByteList, encryptedByte^uint8Max)
	}

	_, decryptedFileWriteError := decryptedFile.Write(decryptByteList)
	if decryptedFileWriteError != nil {
		fmt.Printf("ERROR: write output file occurs error, %v", decryptedFileWriteError)
		return false
	}

	return true
}
