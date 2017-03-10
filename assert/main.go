package assert

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/kr/pretty"
)

var t = (*testing.T)(nil)
var message = "initial"
var submessage = [10]string{}
var subdepth = 0
var count = 0

// Set test title
func Set(tt *testing.T, mes string) {
	message = mes
	count = 0
	t = tt
}

// Sub set sub test title
func Sub(n int, mes string) {
	if n > 0 && n < 10 {
		submessage[n-1] = mes
		count = 0
		subdepth = n
	}
}

func printMessage() {
	fmt.Println(message, count)
	str := "  "
	for i := 0; i < subdepth; i++ {
		fmt.Println(str, submessage[i])
		str += "  "
	}
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
	if !reflect.DeepEqual(actual, expected) {
		printMessage()
		pretty.Printf("--- actual :\n%# v\n\n", actual)
		pretty.Printf("--- expect :\n%# v\n\n", expected)
		t.Errorf("Not equal\n")
	}
}

// EqRegexp is general equal assertion
func EqRegexp(actual string, expected string) {
	count++
	if !regexp.MustCompile(expected).MatchString(actual) {
		printMessage()
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}
}

// Neq is general equal assertion
func Neq(actual interface{}, expected interface{}) {
	count++
	if actual == expected {
		printMessage()
		t.Errorf("\ngot %v\nwant %v", actual, expected)
	}
}

// StubIO stubs Stdin Stdout Stderr in 'fn'.return Stdout and Stderr
func StubIO(inbuf string, fn func()) (string, string) {
	inr, inw, _ := os.Pipe()
	outr, outw, _ := os.Pipe()
	errr, errw, _ := os.Pipe()
	orgStdin := os.Stdin
	orgStdout := os.Stdout
	orgStderr := os.Stderr
	inw.Write([]byte(inbuf))
	inw.Close()
	os.Stdin = inr
	os.Stdout = outw
	os.Stderr = errw
	fn()
	os.Stdin = orgStdin
	os.Stdout = orgStdout
	os.Stderr = orgStderr
	outw.Close()
	outbuf, _ := ioutil.ReadAll(outr)
	errw.Close()
	errbuf, _ := ioutil.ReadAll(errr)

	return string(outbuf), string(errbuf)

}
