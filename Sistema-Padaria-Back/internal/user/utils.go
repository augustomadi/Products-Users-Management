package user

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateToken gera um token aleatório para o usuário
func GenerateToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
