package models

import "time"

type Base struct {
	ID int64 `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (b *Base) BeforeCreat(t time.Time) {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
}

func (b *Base) BeforeUpdate(t time.Time) {
	now := time.Now()
	b.UpdatedAt = now
	if b.DeletedAt != nil {
		b.DeletedAt = nil
	}
}

func (b *Base) BeforeDelete(t time.Time) {
	now := time.Now()
	b.DeletedAt = &now
}

func (b *Base) IsDeleted() bool {
	return b.DeletedAt != nil
}