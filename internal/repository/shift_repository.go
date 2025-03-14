package repository

import (
	"context"
	"shift-scheduling-v2/internal/model"
	"time"

	"github.com/uptrace/bun"
)

type ShiftRepository struct {
	db *bun.DB
}

func NewShiftRepository(db *bun.DB) *ShiftRepository {
	return &ShiftRepository{db: db}
}

func (r *ShiftRepository) GetShiftStatus(ctx context.Context, year int, month int, locationID int) (*model.ShiftsStatus, error) {
	var shiftStatus model.ShiftsStatus
	err := r.db.NewSelect().
		Model(&shiftStatus).
		Where("year = ? AND month = ? AND location_id = ?", year, month, locationID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &shiftStatus, nil
}

func (r *ShiftRepository) UpdateShiftStatus(ctx context.Context, shiftStatus *model.ShiftsStatus) error {
	_, err := r.db.NewUpdate().
		Model(shiftStatus).
		WherePK().
		Exec(ctx)
	return err
}

func (r *ShiftRepository) DeleteShiftsForMonth(ctx context.Context, year int, month int, locationID int) error {
	_, err := r.db.NewDelete().
		Model((*model.Shift)(nil)).
		Where("EXTRACT(YEAR FROM shift_date) = ? AND EXTRACT(MONTH FROM shift_date) = ? AND location_id = ?", year, month, locationID).
		Exec(ctx)
	return err
}

func (r *ShiftRepository) IsDoctorAssignedToShift(ctx context.Context, doctorID int64, shiftDate time.Time) (bool, error) {
	exists, err := r.db.NewSelect().
		Model((*model.Shift)(nil)).
		Where("doctor_id = ? AND shift_date = ?", doctorID, shiftDate).
		Exists(ctx)
	return exists, err
}

func (r *ShiftRepository) Create(ctx context.Context, shift model.Shift) error {
	_, err := r.db.NewInsert().Model(&shift).Exec(ctx)
	return err
}

func (r *ShiftRepository) GetShiftByDate(ctx context.Context, date time.Time) (*model.Shift, error) {
	var shift model.Shift
	err := r.db.NewSelect().
		Model(&shift).
		Where("shift_date = ?", date).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *ShiftRepository) GetTodayShifts(ctx context.Context, date time.Time) ([]model.Shift, error) {
	var shifts []model.Shift
	err := r.db.NewSelect().
		Model(&shifts).
		Relation("Doctor").
		Relation("Doctor.User").
		Relation("Location").
		Where("shift_date = ?", date).
		Scan(ctx)
	return shifts, err
}

func (r *ShiftRepository) GetAllShiftsWithDetails(ctx context.Context) ([]model.Shift, error) {
	var shifts []model.Shift
	err := r.db.NewSelect().
		Model(&shifts).
		Relation("Doctor").
		Relation("Doctor.User").
		Relation("Location").
		Scan(ctx)
	return shifts, err
}

func (r *ShiftRepository) GetShiftsByLocationID(ctx context.Context, locationID int64, month int64, year int64) ([]model.Shift, error) {
	var shifts []model.Shift
	query := r.db.NewSelect().
		Model(&shifts).
		Relation("Doctor").
		Relation("Doctor.User").
		Relation("Location").
		Where("location_id = ?", locationID)

	if month != 0 && year != 0 {
		query = query.Where("EXTRACT(YEAR FROM shift_date) = ? AND EXTRACT(MONTH FROM shift_date) = ?", year, month)
	}

	err := query.Scan(ctx)
	return shifts, err
}

func (r *ShiftRepository) GetAllShift(ctx context.Context) (*[]model.Shift, error) {
	var shifts []model.Shift
	err := r.db.NewSelect().
		Model(&shifts).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &shifts, nil
}

func (r *ShiftRepository) GetShiftByID(ctx context.Context, id int64) (*model.Shift, error) {
	var shift model.Shift
	err := r.db.NewSelect().
		Model(&shift).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *ShiftRepository) DeleteShift(ctx context.Context, id int64) error {
	_, err := r.db.NewDelete().
		Model((*model.Shift)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (r *ShiftRepository) UpdateShift(ctx context.Context, shift model.Shift) error {
	_, err := r.db.NewUpdate().
		Model(&shift).
		WherePK().
		Exec(ctx)
	return err
}

func (r *ShiftRepository) GetShiftsStatus(ctx context.Context) ([]model.ShiftsStatus, error) {
	var shiftsStatus []model.ShiftsStatus
	err := r.db.NewSelect().
		Model(&shiftsStatus).
		Scan(ctx)
	return shiftsStatus, err
}

func (r *ShiftRepository) GetShiftLocations(ctx context.Context) ([]model.ShiftLocation, error) {
	var locations []model.ShiftLocation
	err := r.db.NewSelect().
		Model(&locations).
		Scan(ctx)
	return locations, err
}

func (r *ShiftRepository) GetDoctorsByLocation(ctx context.Context, locationID int64) ([]model.Doctor, error) {
	var doctors []model.Doctor
	err := r.db.NewSelect().
		Model(&doctors).
		Where("location_id = ?", locationID).
		Scan(ctx)
	return doctors, err
}

func (r *ShiftRepository) AssignShiftsForMonth(ctx context.Context, doctors []model.Doctor, locationID int, startOfMonth time.Time, endOfMonth time.Time) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Nöbet atama işlemleri...
	// Bu kısım oldukça karmaşık ve uzun olduğu için, service katmanında kalabilir
	// Ya da ayrı bir "assignment" servisi oluşturulabilir

	return tx.Commit()
}
