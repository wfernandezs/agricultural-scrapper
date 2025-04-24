package models

import "time"

type ProductCategory string

const (
    CategoryFruit      ProductCategory = "fruit"
    CategoryVegetable  ProductCategory = "vegetable"
    CategoryGrain      ProductCategory = "grain"
    CategoryCoffee     ProductCategory = "coffee"
    CategoryCacao      ProductCategory = "cacao"
    CategoryOther      ProductCategory = "other"
)

type Certification string

const (
    CertOrganic       Certification = "organic"
    CertFairTrade     Certification = "fair_trade"
    CertRainforest    Certification = "rainforest_alliance"
    CertGlobalGAP     Certification = "global_gap"	
)

type Product struct {
	Base
	Name 	 string `json:"name" db:"name"`
	Category ProductCategory `json:"category" db:"category"`
	Variety string `json:"variety" db:"variety"`
	Certification []Certification `json:"certification" db:"certification"`
	SeasonStart  time.Month `json:"season_start" db:"season_start"`
	SeasonEnd    time.Month `json:"season_end" db:"season_end"`
	Unit 	 string `json:"unit" db:"unit"`
}

func (p *Product) IsInSeason(date time.Time) bool {
	month := date.Month()
	if p.SeasonStart > p.SeasonEnd {
		return month >= p.SeasonStart || month <= p.SeasonEnd
	}
	return month >= p.SeasonStart && month <= p.SeasonEnd
}