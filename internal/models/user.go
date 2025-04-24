package models

import "time"

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleAnalyst  Role = "analyst"
	RoleStandard Role = "standard"
)

type User struct {
	Base
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         Role      `json:"role" db:"role"`
	APIKey       string    `json:"-" db:"api_key"`
	LastLogin    *time.Time `json:"last_login,omitempty" db:"last_login"`
	
	// User preferences stored as JSON in database
	Preferences  map[string]interface{} `json:"preferences,omitempty" db:"preferences"`
	
	// Subscription details
	SubscriptionTier   string     `json:"subscription_tier,omitempty" db:"subscription_tier"`
	SubscriptionExpiry *time.Time `json:"subscription_expiry,omitempty" db:"subscription_expiry"`
}

// IsAdmin checks if the user has admin privileges
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// HasActiveSubscription checks if the user has an active paid subscription
func (u *User) HasActiveSubscription() bool {
	if u.SubscriptionTier == "" || u.SubscriptionTier == "free" {
		return false
	}
	
	if u.SubscriptionExpiry == nil {
		return false
	}
	
	return time.Now().Before(*u.SubscriptionExpiry)
}

// UpdateLastLogin sets the last login time to now
func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.LastLogin = &now
}