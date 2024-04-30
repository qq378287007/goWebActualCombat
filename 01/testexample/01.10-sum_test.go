package testexample

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	array := []int{6, 8, 10}
	ret := Min(array)
	fmt.Println(ret)
}

// go mod init testexample & go mod tidy

// go test
// go test -v -run TestMin
// go test -v -run="Test"
// go test -v -o testexample.exe & testexample.exe
