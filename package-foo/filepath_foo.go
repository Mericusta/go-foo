package packagefoo

import (
	"fmt"
	"os"
)

func FileInfoForDirLink() {
	f, err := os.Open("./")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		panic(err)
	}

	for _, fInfo := range entries {
		fmt.Printf("fInfo %v %v %v\n", fInfo.Name(), fInfo.IsDir(), fInfo.Mode())
	}
}
