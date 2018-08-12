package main

import (
	"bytes"
	"fmt"
)

const utf8LetterToNumberOffset int32 = 97

func VigenereEncrypt(message, key string) string {
	fmt.Println("Encrypting")
	messageRune := bytes.Runes([]byte(message))
	keyRune := bytes.Runes([]byte(key))
	var cipherText []rune
	for i := range messageRune {
		encryptedValue := ((messageRune[i] - utf8LetterToNumberOffset + keyRune[i] - utf8LetterToNumberOffset) % 26) + utf8LetterToNumberOffset
		cipherText = append(cipherText, encryptedValue)
	}
	return string(cipherText)
}
