package entity

import (
	"time"
)

type Articles []*Article

type Article struct {
	ID          int
	Title       string
	Content     string
	TimesOfRead int
	CreateTime  time.Time
}
