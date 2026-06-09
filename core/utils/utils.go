package utils

import (
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

/* CRYPTOGRAPHY */
type CryptoCollection struct{}

func Crypt() *CryptoCollection {
	crypt := &CryptoCollection{}
	return crypt
}

func (c CryptoCollection) Hash(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(bytes), err
}

/* STRINGS */
type StringCollection struct{}

func String() *StringCollection {
	str := &StringCollection{}
	return str
}

func (str StringCollection) ToSnakeCase(s string) string {
	var result []string
	for i, r := range s {
		if unicode.IsUpper(r) {
			// If it's the first character, just add it in lowercase
			if i > 0 {
				result = append(result, "_")
			}
			result = append(result, string(unicode.ToLower(r)))
		} else {
			result = append(result, string(r))
		}
	}
	return strings.Join(result, "")
}
