package shortener

import (
	"errors"
	"qasr/internal/utils"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mu       sync.RWMutex
)

func CreateShortLink(originalURL string) string {
	slug, _ := utils.GenerateSecureSlug(6)

	mu.Lock()
	defer mu.Unlock()
	urlStore[slug] = originalURL

	return slug
}

func GetOriginalURL(slug string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()

	url, exists := urlStore[slug]
	if !exists {
		return "", errors.New("not found")
	}
	return url, nil
}
