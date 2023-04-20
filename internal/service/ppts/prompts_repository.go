package ppts

import (
	"context"
	"github.com/lqs/sqlingo"
	dsl "monster-base-backend/internal/generated/dsl"
	"monster-base-backend/pkg"
)

type PromptsRepository interface {
	Add(*dsl.PromptsModel, []*dsl.PromptsIndexModel, []*dsl.MediasModel) error
}

type promptsRepository struct {
	db sqlingo.Database
}

func (p promptsRepository) Add(model *dsl.PromptsModel, models []*dsl.PromptsIndexModel, model2 []*dsl.MediasModel) error {
	return p.db.BeginTx(context.Background(), nil, func(tx sqlingo.Transaction) error {
		_, err := tx.InsertInto(dsl.Prompts).Models(model).Execute()
		if err != nil {
			return err
		}

		_, err = tx.InsertInto(dsl.PromptsIndex).Models(models).Execute()
		if err != nil {
			return err
		}
		if len(model2) > 0 {
			_, err = tx.InsertInto(dsl.Medias).Models(model2).Execute()
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func NewPromptsRepository() PromptsRepository {
	return &promptsRepository{
		db: pkg.GetDatabase(),
	}
}
