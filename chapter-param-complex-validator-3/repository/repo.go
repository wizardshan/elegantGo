package repository

import (
	"elegantGo/chapter-param-complex-validator-3/repository/ent"
)

type repo struct {
	db *ent.Client
}
