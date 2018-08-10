package main

import (
	"testing"
	"reflect"
)

func TestVigenereReturnsSlice(t *testing.T) {
	cipherText := VigenereEncrypt("", "")
	kind := reflect.TypeOf(cipherText).Kind()
	if kind != reflect.Slice {
		t.Errorf("Expected slice return value")
	}
}
