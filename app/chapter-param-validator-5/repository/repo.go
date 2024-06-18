package repository

import (
	"app/chapter-param-validator-5/repository/ent"
)

type repo struct {
	db *ent.Client
}
