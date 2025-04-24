package models

import "time"

type PriceType string

const (
	PriceTypeSpot    PriceType = "spot"     // Current market price
	PriceTypeContract PriceType = "contract" // Contract/future price
	PriceTypeExport   PriceType = "export"   // Export market price
	PriceTypeForecast PriceType = "forecast" // Predicted price
)

type Price struct {
	Base
	ProductId			int64 `json:"product_id" db:"product_id"`
	SourceId			int64 `json:"source_id" db:"source_id"`
	LocationId			int64 `json:"location_id" db:"location_id"`
	Value			float64 `json:"value_id" db:"value_id"`
	Currency			string `json:"currency" db:"currency"`
	Volume				float64 `json:"volume" db:"volume"`
	Unit 			string `json:"unit" db:"unit"`
	ObservationDate		time.Time `json:"observation_date" db:"observation_date"`
	PriceType			PriceType `json:"price_type" db:"price_type"`
	Quality			string `json:"quality" db:"quality"`
	IsForecast		bool `json:"is_forecast" db:"is_forecast"`
	ConfidenceLevel		float64 `json:"confidence_level" db:"confidence_level"`
}

// AdjustForInflation calculates inflation-adjusted price based on provided rate
func (p *Price) AdjustForInflation(inflationRate float64) float64 {
	return p.Value * (1 + inflationRate)
}

// ConvertCurrency converts the price to another currency using the provided exchange rate
func (p *Price) ConvertCurrency(targetCurrency string, exchangeRate float64) Price {
	converted := *p
	converted.Value = p.Value * exchangeRate
	converted.Currency = targetCurrency
	return converted
}