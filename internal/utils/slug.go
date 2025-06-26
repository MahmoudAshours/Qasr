package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const slugCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateSecureSlug(n int) (string, error) {
	b := make([]byte, n)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(slugCharset))))
		if err != nil {
			return "", err
		}
		b[i] = slugCharset[num.Int64()]
	}
	return string(b), nil
}

func CleanSlug(raw string) string {
	slug := strings.ToLower(raw)
	slug = strings.TrimSpace(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, ".", "")
	slug = strings.ReplaceAll(slug, "/", "")
	return slug
}
