package repository

import (
	"context"
	"fmt"
	"shift-scheduling-v2/internal/model"
	"shift-scheduling-v2/pkg/cache"
	"time"

	"github.com/uptrace/bun"
)

const (
	doctorCacheKeyPrefix = "doctor:"
	doctorCacheDuration  = 24 * time.Hour
)

type DoctorRepository struct {
	db *bun.DB
}

func NewDoctorRepository(db *bun.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) Create(ctx context.Context, doctor *model.Doctor) error {
	_, err := r.db.NewInsert().Model(doctor).Exec(ctx)
	if err != nil {
		return fmt.Errorf("veritabanı insert hatası: %v", err)
	}
	return nil
}

func (r *DoctorRepository) GetByID(ctx context.Context, id int64, relations ...string) (*model.Doctor, error) {
	cacheKey := fmt.Sprintf("%s%d", doctorCacheKeyPrefix, id)

	var doctor model.Doctor
	err := cache.Get(ctx, cacheKey, &doctor)
	if err == nil {
		return &doctor, nil
	}

	query := r.db.NewSelect().Model(&doctor).Where("id = ?", id)
	for _, relation := range relations {
		query = query.Relation(relation)
	}

	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	if err = cache.Set(ctx, cacheKey, &doctor, doctorCacheDuration); err != nil {
		return &doctor, nil
	}

	return &doctor, nil
}

func (r *DoctorRepository) GetByShiftID(ctx context.Context, shiftID int64) (*model.Doctor, error) {
	var doctor model.Doctor
	err := r.db.NewSelect().Model(&doctor).
		Join("INNER JOIN shifts ON shifts.doctor_id = doctor.id").
		Where("shifts.id = ?", shiftID).
		Scan(ctx)
	return &doctor, err
}

func (r *DoctorRepository) GetByLocation(ctx context.Context, locationID int64) ([]model.Doctor, error) {
	var doctors []model.Doctor
	err := r.db.NewSelect().Model(&doctors).
		Join("INNER JOIN doctor_shift_locations dsl ON dsl.doctor_id = doctor.id").
		Where("dsl.location_id = ?", locationID).
		Relation("User").
		Scan(ctx)
	return doctors, err
}

func (r *DoctorRepository) GetHolidaysByDoctor(ctx context.Context, doctorID int64) ([]model.Holiday, error) {
	var holidays []model.Holiday
	err := r.db.NewSelect().Model(&holidays).
		Where("doctor_id = ?", doctorID).
		Scan(ctx)
	return holidays, err
}

func (r *DoctorRepository) GetHolidaysByLocation(ctx context.Context, locationID int64, month, year int64) ([]model.Holiday, error) {
	var holidays []model.Holiday
	query := r.db.NewSelect().Model(&holidays).
		Relation("Doctor.User").
		Relation("Location").
		Where("location_id = ?", locationID)

	if month != 0 && year != 0 {
		query = query.Where("EXTRACT(YEAR FROM holiday_date) = ? AND EXTRACT(MONTH FROM holiday_date) = ?", year, month)
	}

	err := query.Order("holiday_date ASC").Scan(ctx)
	return holidays, err
}

func (r *DoctorRepository) List(ctx context.Context, relations ...string) ([]model.Doctor, int, error) {
	var doctors []model.Doctor
	query := r.db.NewSelect().Model(&doctors)

	for _, relation := range relations {
		query = query.Relation(relation)
	}

	err := query.Scan(ctx)
	return doctors, len(doctors), err
}

func (r *DoctorRepository) Update(ctx context.Context, doctor *model.Doctor) error {
	_, err := r.db.NewUpdate().Model(doctor).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	// Cache'i temizle
	cacheKey := fmt.Sprintf("%s%d", doctorCacheKeyPrefix, doctor.ID)
	_ = cache.Delete(ctx, cacheKey)

	return nil
}

func (r *DoctorRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.NewDelete().Model((*model.Doctor)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}

	// Cache'i temizle
	cacheKey := fmt.Sprintf("%s%d", doctorCacheKeyPrefix, id)
	_ = cache.Delete(ctx, cacheKey)

	return nil
}
