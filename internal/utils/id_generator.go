package utils

import (
	"math/rand"
	"unicode/utf8"
)

const (
	availableSymbols = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	idLen            = 6
)

func GenerateID() string {
	symCount := utf8.RuneCountInString(availableSymbols)
	id := ""
	for i := 0; i < idLen; i++ {
		randomIndex := rand.Intn(symCount)
		id += string(availableSymbols[randomIndex])
	}
	return id
}
