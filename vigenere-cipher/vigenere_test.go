package main

import (
	"reflect"
	"testing"
)

func TestVigenereReturnsSlice(t *testing.T) {
	cipherText := VigenereEncrypt("", "")
	kind := reflect.TypeOf(cipherText).Kind()
	if kind != reflect.String {
		t.Errorf("Expected string return value, got %s", kind)
	}
}

func TestVigenereEncryptionOfMessageOfSameLengthAsKey(t *testing.T) {
	cipherText := VigenereEncrypt("test", "ball")
	expectedEncryptedText := "uede"
	if cipherText != expectedEncryptedText {
		t.Errorf("Encryption error for string of equal length to key, expected %s, got %s", expectedEncryptedText, cipherText)
	}
}

func TestVigenereEncryptionOfMessageLongerThanKey(t *testing.T) {
	cipherText := VigenereEncrypt("testmessagethatisnoneofyourbusiness", "racetrack")
	expectedEncryptedText := "keuxfvsukxevltkiuxfngsypowbsuumgvsu"
	if cipherText != expectedEncryptedText {
		t.Errorf("Encryption error for string of greater length than key, expected %s, got %s", expectedEncryptedText, cipherText)
	}
}
