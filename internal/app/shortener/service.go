package shortener

import (
	"qasr/internal/repo/mongodb"
	"qasr/internal/utils"
)

type ShortenerService struct {
	Repo *mongodb.LinkRepository
}

func NewShortenerService(repo *mongodb.LinkRepository) *ShortenerService {
	return &ShortenerService{Repo: repo}
}

func (s *ShortenerService) CreateShortLink(url string) string {
	for {
		slug, _ := utils.GenerateSecureSlug(6)
		if !s.Repo.SlugExists(slug) {
			s.Repo.Create(slug, url)
			return slug
		}
	}
}

func (s *ShortenerService) GetOriginalURL(slug string) (string, error) {
	return s.Repo.FindBySlug(slug)
}
