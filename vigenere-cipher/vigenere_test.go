package main

import (
	"reflect"
	"testing"
)

func TestVigenereReturnsSlice(t *testing.T) {
	cipherText := VigenereEncrypt("", "")
	kind := reflect.TypeOf(cipherText).Kind()
	if kind != reflect.Slice {
		t.Errorf("Expected slice return value")
	}
}

func TestVigenereEncryptionOfMessageOfSameLengthAsKey(t *testing.T) {
	cipherText := VigenereEncrypt("test", "ball")
	if cipherText != "vfeel" {
		t.ErrorF("Encryption error for string of equal length to key")
	}
}
