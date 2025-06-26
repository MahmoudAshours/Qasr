package analytics

import (
	"qasr/internal/repo/mongodb"
	"time"
)

type AnalyticsService struct {
	Repo *mongodb.ClickRepo
}

// 🔧 Constructor
func NewAnalyticsService(repo *mongodb.ClickRepo) *AnalyticsService {
	return &AnalyticsService{Repo: repo}
}

// 📊 Get stats for a specific slug
func (a *AnalyticsService) GetAnalytics(slug string) (map[string]interface{}, error) {
	clicks, err := a.Repo.GetClicksBySlug(slug)
	if err != nil {
		return nil, err
	}

	countries := map[string]int{}
	referrers := map[string]int{}
	browsers := map[string]int{}
	devices := map[string]int{}
	last7Days := 0
	clicksPerDay := map[string]int{}

	sevenDaysAgo := time.Now().Add(-7 * 24 * time.Hour)

	for _, c := range clicks {
		if c.Timestamp.After(sevenDaysAgo) {
			last7Days++
		}

		day := c.Timestamp.Format("2006-01-02")
		clicksPerDay[day]++

		countries[c.Country]++
		ref := c.Referrer
		if ref == "" {
			ref = "Direct"
		}
		referrers[ref]++
		browsers[c.Browser]++
		devices[c.DeviceType]++
	}

	return map[string]interface{}{
		"slug":           slug,
		"total_clicks":   len(clicks),
		"last_7_days":    last7Days,
		"clicks_per_day": clicksPerDay,
		"top_countries":  countries,
		"top_referrers":  referrers,
		"browsers":       browsers,
		"device_types":   devices,
	}, nil
}
