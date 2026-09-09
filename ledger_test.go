package main

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	result := doSomething()
	if result != 10 {
		t.Errorf(`Expected to get 1`)
	}
}
