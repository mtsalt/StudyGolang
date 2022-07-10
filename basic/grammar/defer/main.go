package main

import (
	"fmt"
	"os"
)

func openTest(filepath string) {
	fp, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer fp.Close()

	buf := make([]byte, 64)
	for {
		n, err := fp.Read(buf)

		if n == 0 {
			break
		}

		if err != nil {
			return
		}

		fmt.Println(string(buf))
	}
}

func main() {
	openTest("sample.txt")
}
