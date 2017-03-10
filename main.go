package main

import "github.com/nak1114/goutil/assert"

type TestT map[string]string

func main() {
	var a = TestT{}
	var b = TestT{}
	a["aaa"] = "bbb"
	a["ccc"] = "ddd"
	b["aaa"] = "bbb"
	b["ccc"] = "ddd"
	assert.Eq(a, b)
}
