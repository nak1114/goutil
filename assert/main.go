package assert

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var t = (*testing.T)(nil)
var message = "initial"
var submessage = []string{}
var count = 0

// Set test title
func Set(tt *testing.T, mes string) {
	message = mes
	count = 0
	t = tt
}

func printMessage() {
	fmt.Println(message, count)
}

// FileContent is expect file content.
func FileContent(fname string, content string) {
	count++
	if _, err := os.Stat(fname); err != nil {
		printMessage()
		t.Errorf("file not found.(%s)\n", fname)
		return
	}
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		printMessage()
		t.Errorf("file can't read.(%s)\n", fname)
		return
	}
	sbuf := string(buf)
	if sbuf != content {
		printMessage()
		fmt.Println("actual:", sbuf)
		fmt.Println("expect:", content)
		t.Errorf("content not equal.\n")
		return
	}
	return
}

// Eq is general equal assertion
func Eq(actual interface{}, expected interface{}) {
	count++
	if actual != expected {
		printMessage()
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}
}
