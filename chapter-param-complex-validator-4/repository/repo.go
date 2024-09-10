package repository

import (
	"elegantGo/chapter-param-complex-validator-4/repository/ent"
)

type repo struct {
	db *ent.Client
}
