package repository

import (
	"app/chapter-param-validator-7/repository/ent"
)

type repo struct {
	db *ent.Client
}
