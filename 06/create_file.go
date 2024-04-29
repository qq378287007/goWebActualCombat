package main

import (
	"fmt"
	"os"
)

func main() {
	defer os.Remove("create_file.txt")

	fp, err := os.Create("create_file.txt")
	if err != nil {
		return
	}
	defer fp.Close()

	for i := 0; i < 5; i++ {
		outstr := fmt.Sprintf("%s:%d\r\n", "Hello Go", i)
		fp.WriteString(outstr)
		fp.Write([]byte("i love go\r\n"))
	}

}
