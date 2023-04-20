package ppts

import (
	"github.com/lqs/sqlingo"
	"monster-base-backend/pkg"
)

type PromptsRepository interface {
}

type promptsRepository struct {
	db sqlingo.Database
}

func NewPromptsRepository() PromptsRepository {
	return &promptsRepository{
		db: pkg.GetDatabase(),
	}
}
