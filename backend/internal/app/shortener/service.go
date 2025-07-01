package shortener

import (
	"qasr/backend/internal/repo/mongodb"
	redisCache "qasr/backend/internal/repo/redis"
	"qasr/backend/internal/utils"
	"time"
)

type ShortenerService struct {
	Repo  *mongodb.LinkRepository
	Cache *redisCache.RedisCache
}

func NewShortenerService(repo *mongodb.LinkRepository, cache *redisCache.RedisCache) *ShortenerService {
	return &ShortenerService{Repo: repo, Cache: cache}
}

func (s *ShortenerService) CreateShortLink(url string) string {
	for {
		slug, _ := utils.GenerateSecureSlug(6)
		if !s.SlugExists(slug) {
			s.Repo.Create(slug, url)
			s.Cache.Set(slug, url, 30*time.Minute)
			return slug
		}
	}
}

func (s *ShortenerService) GetOriginalURL(slug string) (string, error) {
	// Try Redis first
	url, err := s.Cache.Get(slug)
	if err == nil && url != "" {
		return url, nil
	}

	// Fallback to MongoDB
	url, err = s.Repo.FindBySlug(slug)
	if err == nil {
		_ = s.Cache.Set(slug, url, 30*time.Minute)
	}
	return url, err
}

func (s *ShortenerService) SlugExists(slug string) bool {
	// Step 1: Check Redis
	url, err := s.Cache.Get(slug)
	if err == nil && url != "" {
		return true
	}
	return false
}

func (s *ShortenerService) GetBySlug(slug string) (string, error) {
	url, err := s.Cache.Get(slug)
	if err == nil && url != "" {
		return url, nil
	}
	url, err = s.Repo.FindBySlug(slug)
	if err == nil {
		_ = s.Cache.Set(slug, url, 30*time.Minute)
	}
	return url, err
}
