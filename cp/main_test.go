package cp

import (
	"fmt"
	"os"
	"testing"
	//"os/exec"
	//"runtime"
)

var message = "initial"
var count = 0

func set(mes string) {
	message = mes
	count = 0
}
func eq(t *testing.T, actual interface{}, expected interface{}) {
	count++
	if actual != expected {
		fmt.Println(message, count)
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestDir(t *testing.T) {
	e := Dir("./misc/test", "./misc/test2")
	fmt.Println(e)
	e = File("./misc/test/hoge.txt", "./misc/")
	fmt.Println(e)
	e = Dir("./misc/test_dup", "./misc/test2")
	fmt.Println(e)

}
