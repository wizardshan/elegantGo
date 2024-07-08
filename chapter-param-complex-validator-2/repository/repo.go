package repository

import (
	"elegantGo/chapter-param-complex-validator-2/repository/ent"
)

type repo struct {
	db *ent.Client
}
