package models

type LocationType string

const (
	LocationTypeFarm    LocationType = "farm"
	LocationTypeMarket  LocationType = "market"
	LocationTypePort    LocationType = "port"
	LocationTypeRegion  LocationType = "region"
	LocationTypeCountry LocationType = "country"
)

type Location struct {
	Base
	Name        string      `json:"name" db:"name"`
	Region      string      `json:"region" db:"region"`
	Type        LocationType `json:"type" db:"type"`
	Latitude    float64     `json:"latitude" db:"latitude"`
	Longitude   float64     `json:"longitude" db:"longitude"`
	ParentID    *int64      `json:"parent_id,omitempty" db:"parent_id"`
	Description string      `json:"description" db:"description"`
	
	// Related data
	Parent     *Location   `json:"parent,omitempty" db:"-"`
	Children   []Location  `json:"children,omitempty" db:"-"`
	Prices     []Price     `json:"prices,omitempty" db:"-"`
}

// FullPath returns the complete location path (e.g., "Lima, Peru")
func (l *Location) FullPath() string {
	if l.Parent == nil || l.ParentID == nil {
		return l.Name
	}
	return l.Name + ", " + l.Parent.FullPath()
}

// DistanceTo calculates the approximate distance to another location in kilometers
// using the Haversine formula
func (l *Location) DistanceTo(other *Location) float64 {
	// Simplified implementation - in a real app, use a proper geospatial library
	// This is just a placeholder for the concept
	return 0.0 // Replace with actual calculation
}