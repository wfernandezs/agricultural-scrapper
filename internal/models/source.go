package models

import "time"

type SourceType string

const (
	SourceTypeGovernment  SourceType = "government"
	SourceTypeMarket      SourceType = "market"
	SourceTypeAssociation SourceType = "association"
	SourceTypeCompany     SourceType = "company"
	SourceTypeMedia       SourceType = "media"
	SourceTypeOther       SourceType = "other"
)

type Source struct {
	Base
	Name 			string `json:"name" db:"name"`
	URL 			string `json:"url" db:"url"`
	Type 			SourceType `json:"type" db:"type"`
	Description 	string `json:"description" db:"description"`
	Reliability 		float64 `json:"reliability" db:"reliability"`
	UpdateFrequency 	string `json:"update_frequency" db:"update_frequency"`
	LastScraped     *time.Time `json:"last_scraped,omitempty" db:"last_scraped"`
	ScraperConfig map[string]interface{} `json:"scraper_config" db:"scraper_config"`

		// Related data
	Prices []Price `json:"prices,omitempty" db:"-"`
}

func (s *Source) ShouldScape() bool {
	if s.LastScraped == nil {
		return true
	}

	now := time.Now()
	switch s.UpdateFrequency {
	case "hourly":
		return now.Sub(*s.LastScraped) > time.Hour
	case "daily":
		return now.Sub(*s.LastScraped) > 24*time.Hour
	case "weekly":
		return now.Sub(*s.LastScraped) > 7*24*time.Hour
	case "monthly":
		return now.Sub(*s.LastScraped) > 30*24*time.Hour
	default:
		return now.Sub(*s.LastScraped) >= 24*time.Hour // Default to daily
	}
}

func (s *Source) UpdateLastScraped() {
	now := time.Now()
	s.LastScraped = &now
}
