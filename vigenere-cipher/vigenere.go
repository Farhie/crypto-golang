package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func VigenereEncrypt(message, key string) string {
	fmt.Println("Encrypting using Vigenere cipher")
	messageRune := bytes.Runes([]byte(message))
	keyRune := bytes.Runes([]byte(strings.Repeat(key, utf8.RuneCountInString(message))))
	var cipherText []rune

	for i := range messageRune {
		encryptedValue := computerEncryptedValue(messageRune[i], keyRune[i])
		cipherText = append(cipherText, encryptedValue)
	}

	return string(cipherText)
}

func computerEncryptedValue(messageByte, keyByte rune) rune {
	var utf8ToAsciiOffset int32 = 97
	return ((messageByte - utf8ToAsciiOffset + keyByte - utf8ToAsciiOffset) % 26) + utf8ToAsciiOffset
}
