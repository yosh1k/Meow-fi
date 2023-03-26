package auth

import (
	"Meow-fi/internal/config"
	"math/rand"
	"time"
)

func RandSeq() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, config.LenRandomSalt)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func HashPass(password, randomSalt, localSalt string) string {
	return password + randomSalt + localSalt
}
