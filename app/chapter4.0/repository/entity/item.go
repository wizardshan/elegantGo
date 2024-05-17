package entity

import (
	"time"
)

type Items []*Item

type Item struct {
	ID         int
	Title      string
	CreateTime time.Time
	DeleteTime *time.Time
}
