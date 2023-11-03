// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"llm_plateform_resourcemanagement/internal/entity"
)

type (
	Test interface {
		History(context.Context) ([]entity.Test, error)
		Add(context.Context, entity.Test) error
	}

	TestMemory interface {
		Store(context.Context, entity.Test) error
		GetHistory(context.Context) ([]entity.Test, error)
	}
)
