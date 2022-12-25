package main

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type FilterCond struct {
	Title        string    `filter:"like,omitempty"`
	Tags         []string  `filter:"in,omitempty"`
	MinCreatedAt time.Time `filter:"gte,omitempty"`
	MaxCreatedAt time.Time `filter:"lte,omitempty"`
	NotCreatedBy string    `filter:"ne,omitempty"`
}

// db.Scopes(generateFilter(f)).Find(&posts)

func generateFilter(f *FilterCond) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if f.Title != "" {
			db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", f.Title))
		}
		if len(f.Tags) > 0 {
			db = db.Where("tags IN ?", f.Tags)
		}
		if !f.MinCreatedAt.IsZero() {
			db = db.Where("created_at >= ?", f.MinCreatedAt)
		}
		if !f.MaxCreatedAt.IsZero() {
			db = db.Where("created_at <= ?", f.MaxCreatedAt)
		}
		if f.NotCreatedBy != "" {
			db = db.Where("created_by <> ?", f.NotCreatedBy)
		}
		return db
	}
}
