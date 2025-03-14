package repository

import (
	"context"
	"shift-scheduling-v2/internal/model"

	"github.com/uptrace/bun"
)

type ShiftRepository struct {
	db *bun.DB
}

func NewShiftRepository(db *bun.DB) *ShiftRepository {
	return &ShiftRepository{db: db}
}

func (r *ShiftRepository) Create(ctx context.Context, shift model.Shift) error {
	_, err := r.db.NewInsert().Model(&shift).Exec(ctx)
	return err
}

// Diğer repository fonksiyonları...
