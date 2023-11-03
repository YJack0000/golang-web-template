package usecase

import (
	"context"
	"fmt"

	"llm_plateform_resourcemanagement/internal/entity"
)

type TestUseCase struct {
	memo TestMemory
}

func New(r TestMemory) *TestUseCase {
	return &TestUseCase{
		memo: r,
	}
}

func (uc *TestUseCase) History(ctx context.Context) ([]entity.Test, error) {
	Tests, err := uc.memo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TestUseCase - History - s.memo.GetHistory: %w", err)
	}

	return Tests, nil
}

func (uc *TestUseCase) Add(ctx context.Context, Test entity.Test) error {
	err := uc.memo.Store(ctx, Test)
	if err != nil {
		return fmt.Errorf("TestUseCase - Add - s.memo.Add: %w", err)
	}

	return nil
}
