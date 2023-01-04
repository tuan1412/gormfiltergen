package models

import (
	"time"
)

type FilterCond struct {
	Title        string    `filter:"like,omitempty"`
	Tags         []string  `filter:"in,omitempty"`
	MinCreatedAt time.Time `filter:"gte,omitempty"`
	MaxCreatedAt time.Time `filter:"lte,omitempty"`
	NotCreatedBy string    `filter:"ne,omitempty"`
}

type FilterCond2 struct {
	Title string `filter:"startWith,omitempty"`
}

func BuildQuery(s1 string, s2 string) string {
	return s1 + s2
}
