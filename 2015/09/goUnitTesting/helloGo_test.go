package myAwesomePackage

import "testing"

func TestHelloGo(t *testing.T) {
	test := MyAwesomeFunction("joe")
	if test != "hello joe" {
		t.Log("Expected: 'hello joe' got: " + test) // HL
		t.Fail() // HL
	}
}