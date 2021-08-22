package tool

import (
	"fmt"
	"testing"
)

func TestPrintStackTrace(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(PrintStackTrace(err))
		}
	}()

	f1()
}

func f1() {
	f2()
}

func f2() {
	f3()
}

func f3() {
	var m map[int]int
	m[0] = 0
}
