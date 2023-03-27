package main

import "testing"

func TestFoobar(t *testing.T) {
	if foobar() != 42 {
		t.Fail()
	}
}
