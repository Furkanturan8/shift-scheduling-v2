package handler

import (
	"errors"
	"shift-scheduling-v2/internal/dto"
	"shift-scheduling-v2/internal/model"
	"shift-scheduling-v2/internal/service"
	"shift-scheduling-v2/pkg/errorx"
	"shift-scheduling-v2/pkg/response"
	"strings"

	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ShiftHandler struct {
	shiftService  *service.ShiftService
	doctorService *service.DoctorService
}

func NewShiftHandler(s *service.ShiftService, d *service.DoctorService) *ShiftHandler {
	return &ShiftHandler{shiftService: s, doctorService: d}
}

func (h ShiftHandler) Create(ctx *fiber.Ctx) error {
	var vm dto.ShiftCreateRequest
	if err := ctx.BodyParser(&vm); err != nil {
		return errorx.ErrInvalidRequest
	}

	shift := vm.ToDBModel(model.Shift{})

	err := h.shiftService.CreateShift(ctx.Context(), shift)
	if err != nil {
		return err
	}

	return response.Success(ctx, nil, "Shift created successfully")
}

func (h ShiftHandler) AssignShiftsForMonth(ctx *fiber.Ctx, doctors []model.Doctor, locationID int, startOfMonth time.Time, endOfMonth time.Time) error {
	// 1. Tüm doktorların tatil günlerini al
	holidayMap := make(map[int64][]dto.DoctorHolidayDTO)
	for _, doctor := range doctors {
		holidays, err := h.doctorService.GetDoctorHolidays(ctx.Context(), doctor.ID)
		if err != nil {
			return err
		}
		holidayMap[doctor.ID] = holidays
	}

	// 2. Nöbet günlerini dağıt
	shiftAssignments := make(map[int64]int) // Her doktorun aldığı nöbet sayısı
	var unassignedDays []time.Time          // Atanamayan günlerin listesi

	for shiftDate := startOfMonth; shiftDate.Before(endOfMonth); shiftDate = shiftDate.AddDate(0, 0, 1) {
		var selectedDoctorID int64
		for _, doctor := range doctors {
			// Doktorun tatilde olup olmadığını kontrol et
			isHoliday := false
			for _, holiday := range holidayMap[doctor.ID] {
				if holiday.HolidayDate.Equal(shiftDate) {
					isHoliday = true
					break
				}
			}
			if isHoliday {
				continue
			}

			// Doktorun nöbet limitine ulaşıp ulaşmadığını kontrol et
			if shiftAssignments[doctor.ID] >= doctor.ShiftLimit {
				continue
			}

			// Doktorun zaten o tarihte atanmış bir nöbeti olup olmadığını kontrol et
			isAssigned, err := h.shiftService.IsDoctorAssignedToShift(ctx.Context(), doctor.ID, shiftDate)
			if err != nil {
				return err
			}
			if isAssigned {
				continue
			}

			// Doktor uygun, nöbet atamasını yap
			selectedDoctorID = doctor.ID
			shiftAssignments[doctor.ID]++
			break
		}

		// Eğer uygun doktor bulunmadıysa, o gün boş geçilecek
		if selectedDoctorID == 0 {
			unassignedDays = append(unassignedDays, shiftDate)
			continue
		}

		// Nöbeti oluştur
		shift := model.Shift{
			DoctorID:   selectedDoctorID,
			LocationID: int64(locationID),
			ShiftDate:  shiftDate,
		}
		err := h.shiftService.CreateShift(ctx.Context(), shift)
		if err != nil {
			return err
		}
	}

	// todo -> 3. Atanamayan günleri, eksik nöbeti olan doktorlara dağıt

	// İşlem tamamlandığında, eğer boş kalan günler varsa uyarı ver
	if len(unassignedDays) > 0 {
		var unassignedDaysStr []string
		for _, day := range unassignedDays {
			unassignedDaysStr = append(unassignedDaysStr, day.Format("2006-01-02"))
		}
		return errors.New("Aşağıdaki günlerde doktor atanamadı: " + strings.Join(unassignedDaysStr, ", "))
	}

	return nil
}

func (h ShiftHandler) AutoAssignShifts(ctx *fiber.Ctx) error {
	var vm dto.AutoAssignShiftDTO
	if err := ctx.BodyParser(&vm); err != nil {
		return errorx.ErrInvalidRequest
	}

	shift := vm.ToDBModel(model.ShiftsStatus{})
	year := shift.Year
	month := shift.Month
	locationID := shift.LocationID

	doctors, err := h.shiftService.GetDoctorsByLocation(ctx.Context(), locationID)
	if err != nil {
		return err
	}
	if len(doctors) == 0 {
		return errorx.ErrNotFound
	}

	shiftStatus, err := h.shiftService.GetShiftStatus(ctx.Context(), year, month, int(locationID))
	if err != nil {
		return err
	}

	if shiftStatus.Done {
		return errorx.ErrInvalidRequest
	}

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	err = h.shiftService.AssignShiftsForMonth(ctx.Context(), doctors, int(locationID), startOfMonth, endOfMonth)
	if err != nil {
		return err
	}

	err = h.shiftService.MarkShiftStatusAsDone(ctx.Context(), year, month, int(locationID))
	if err != nil {
		return err
	}

	return response.Success(ctx, nil, "Shifts assigned successfully")
}

func (h ShiftHandler) ResetShifts(ctx *fiber.Ctx) error {
	var vm dto.AutoAssignShiftDTO
	if err := ctx.BodyParser(&vm); err != nil {
		return errorx.ErrInvalidRequest
	}

	shift := vm.ToDBModel(model.ShiftsStatus{})
	year := shift.Year
	month := shift.Month
	locationID := shift.LocationID

	err := h.shiftService.ResetShiftsForMonth(ctx.Context(), year, month, int(locationID))
	if err != nil {
		return err
	}

	return response.Success(ctx, nil, "Shifts reset successfully")
}

func (h ShiftHandler) GetShiftByDate(ctx *fiber.Ctx) error {
	param := ctx.Params("date")
	date, err := time.Parse("2006-01-02", param)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	shift, err := h.shiftService.GetShiftByDate(ctx.Context(), date)
	if err != nil {
		return err
	}

	shiftListVM := dto.ShiftResponse{}.ToResponseModel(*shift)

	return response.Success(ctx, shiftListVM, "Shift retrieved successfully")
}

func (h ShiftHandler) GetTodayShifts(ctx *fiber.Ctx) error {
	loc, _ := time.LoadLocation("Europe/Istanbul")
	now := time.Now().In(loc)

	todayStr := fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day())
	today, _ := time.ParseInLocation("2006-01-02", todayStr, loc)

	shifts, err := h.shiftService.GetTodayShifts(ctx.Context(), today)
	if err != nil {
		return err
	}

	shiftListVM := make([]dto.ShiftListWithDetailsDTO, len(shifts))
	for i, shift := range shifts {
		shiftListVM[i] = dto.ShiftListWithDetailsDTO{}.ToResponseModel(shift)
	}

	return response.Success(ctx, shiftListVM, "Today's shifts retrieved successfully")
}

func (h ShiftHandler) GetAllShiftsWithDetails(ctx *fiber.Ctx) error {
	shifts, err := h.shiftService.GetAllShiftsWithDetails(ctx.Context())
	if err != nil {
		return errorx.ErrInternal
	}

	shiftListVM := make([]dto.ShiftListWithDetailsDTO, len(shifts))
	for i, shift := range shifts {
		shiftListVM[i] = dto.ShiftListWithDetailsDTO{}.ToResponseModel(shift)
	}

	return response.Success(ctx, shiftListVM, "All shifts with details retrieved successfully")
}

func (h ShiftHandler) GetShiftsByLocationID(ctx *fiber.Ctx) error {
	param := ctx.Params("location_id")
	locationID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	query := ctx.Query("month")
	month, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	query2 := ctx.Query("year")
	year, err := strconv.ParseInt(query2, 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	shifts, err := h.shiftService.GetShiftsByLocationID(ctx.Context(), locationID, month, year)
	if err != nil {
		return errorx.ErrInternal
	}

	shiftListVM := make([]dto.ShiftListWithDetailsDTO, len(shifts))
	for i, shift := range shifts {
		shiftListVM[i] = dto.ShiftListWithDetailsDTO{}.ToResponseModel(shift)
	}

	return response.Success(ctx, shiftListVM, "Shifts by location ID retrieved successfully")
}

func (h ShiftHandler) GetAllShifts(ctx *fiber.Ctx) error {
	shifts, err := h.shiftService.GetAllShift(ctx.Context())
	if err != nil {
		return errorx.ErrInternal
	}

	shiftListVM := make([]dto.ShiftResponse, len(*shifts))
	for i, shift := range *shifts {
		shiftListVM[i] = dto.ShiftResponse{}.ToResponseModel(shift)
	}

	return response.Success(ctx, shiftListVM, "All shifts retrieved successfully")
}

func (h ShiftHandler) GetByShiftID(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	shift, err := h.shiftService.GetShiftByID(ctx.Context(), id)
	if err != nil {
		if err.Error() == "record not found" {
			return errorx.ErrNotFound
		}
		return errorx.ErrInternal
	}

	shiftListVM := dto.ShiftResponse{}.ToResponseModel(*shift)

	return response.Success(ctx, shiftListVM, "Shift retrieved successfully")
}

func (h ShiftHandler) DeleteShift(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	err = h.shiftService.DeleteShift(ctx.Context(), id)
	if err != nil {
		return errorx.ErrInternal
	}

	return response.Success(ctx, nil, "Shift deleted successfully")
}

func (h ShiftHandler) UpdateShift(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	m, err := h.shiftService.GetShiftByID(ctx.Context(), id)
	if err != nil {
		return err
	}

	var vm dto.ShiftCreateRequest
	if err := ctx.BodyParser(&vm); err != nil {
		return errorx.ErrInvalidRequest
	}

	updatedShift := vm.ToDBModel(*m)
	err = h.shiftService.UpdateShift(ctx.Context(), updatedShift)
	if err != nil {
		return err
	}

	return response.Success(ctx, nil, "Shift updated successfully")
}

func (h ShiftHandler) GetShiftsStatus(ctx *fiber.Ctx) error {
	shiftsStatus, err := h.shiftService.GetShiftsStatus(ctx.Context())
	if err != nil {
		return errorx.ErrInternal
	}

	return response.Success(ctx, shiftsStatus, "Shifts status retrieved successfully")
}

func (h ShiftHandler) GetShiftLocations(ctx *fiber.Ctx) error {
	shiftLocations, err := h.shiftService.GetShiftLocations(ctx.Context())
	if err != nil {
		return errorx.ErrInternal
	}

	shiftLocationsVM := make([]dto.ShiftLocationDTO, len(shiftLocations))
	for i, location := range shiftLocations {
		shiftLocationsVM[i] = dto.ShiftLocationDTO{}.ToResponseModel(location)
	}

	return response.Success(ctx, shiftLocationsVM, "Shift locations retrieved successfully")
}
