package memo

import (
	"context"
	"fmt"

	"llm_plateform_resourcemanagement/internal/entity"
)

const _defaultEntityCap = 64

type TestMemory struct {
	entities []entity.Test
}

func New() *TestMemory {
	return &TestMemory{}
}

func (r *TestMemory) GetHistory(ctx context.Context) ([]entity.Test, error) {
	entities := make([]entity.Test, 0, _defaultEntityCap)

	for _, e := range r.entities {
		entities = append(entities, e)
	}

	return entities, nil
}

func (r *TestMemory) Store(ctx context.Context, t entity.Test) error {
	r.entities = append(r.entities, t)
	fmt.Println(r.entities)
	return nil
}
