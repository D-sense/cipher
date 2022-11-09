// Package caesar provides support for encoding and decoding using Caesar Cipher principle.
// The encryption and decryption of the text is done using a key that specifies the alphabet rotation factor.
package caesar

import (
	"fmt"
	"regexp"
	"service/business/core/models"
	"service/business/system/validation"
	"strings"
	"unicode"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

// Encrypter encrypts a string using Caesar Cipher principle.
func Encrypter(data models.NewEncoding) (string, error) {
	// validate input data
	if err := validation.Check(data); err != nil {
		return "", fmt.Errorf("validating data: %w", err)
	}

	return caesarCipher(data.Text, data.Key), nil
}

// Decrypter decrypts a string using Caesar Cipher principle.
func Decrypter(data models.NewDecoding) (string, error) {
	// validate input data
	if err := validation.Check(data); err != nil {
		return "", fmt.Errorf("validating data: %w", err)
	}

	data.Key = 26 - (data.Key % 26)
	return caesarCipher(data.Text, data.Key), nil
}

// caesarCipher encrypts or decrypts a string using a key that specifies the alphabet rotation factor.
func caesarCipher(str string, key int) string {
	runes := []rune(str)
	re := regexp.MustCompile("A-Za-z")

	for i, char := range runes {
		isCurrentCharUpperCase := unicode.IsUpper(char)
		newChar := string(char)
		if isCurrentCharUpperCase {
			newChar = strings.ToLower(newChar)
		}
		index := strings.Index(alphabet, newChar)
		if index == -1 && !re.MatchString(newChar) {
			continue
		}

		newIndex := (index + key) % 26
		if isCurrentCharUpperCase {
			runes[i] = unicode.ToUpper(rune(alphabet[newIndex]))
			continue
		}

		runes[i] = rune(alphabet[newIndex])
	}

	return string(runes)
}
