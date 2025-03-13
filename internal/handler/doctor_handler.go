package handler

import (
	"shift-scheduling-v2/internal/dto"
	"shift-scheduling-v2/internal/service"
	"shift-scheduling-v2/pkg/errorx"
	"shift-scheduling-v2/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DoctorHandler struct {
	service *service.DoctorService
}

func NewDoctorHandler(s *service.DoctorService) *DoctorHandler {
	return &DoctorHandler{service: s}
}

func (h *DoctorHandler) GetDoctorsByLocation(c *fiber.Ctx) error {
	locationID, err := strconv.ParseInt(c.Params("location_id"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	resp, err := h.service.GetDoctorsByLocation(c.Context(), locationID)
	if err != nil {
		return errorx.WithDetails(errorx.ErrInternal, err.Error())
	}

	return response.Success(c, resp)
}

func (h *DoctorHandler) GetDoctorByShiftID(c *fiber.Ctx) error {
	shiftID, err := strconv.ParseInt(c.Params("shift_id"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	resp, err := h.service.GetDoctorByShiftID(c.Context(), shiftID)
	if err != nil {
		return errorx.WithDetails(errorx.ErrNotFound, "Doktor bulunamadı")
	}

	return response.Success(c, resp)
}

func (h *DoctorHandler) GetDoctorHolidays(c *fiber.Ctx) error {
	doctorID, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	resp, err := h.service.GetDoctorHolidays(c.Context(), doctorID)
	if err != nil {
		return errorx.WithDetails(errorx.ErrInternal, err.Error())
	}

	return response.Success(c, resp)
}

func (h *DoctorHandler) GetDoctorsHolidayByLocationId(c *fiber.Ctx) error {
	locationID, err := strconv.ParseInt(c.Params("location_id"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	month, err := strconv.ParseInt(c.Query("month"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	year, err := strconv.ParseInt(c.Query("year"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	resp, err := h.service.GetDoctorsHolidayByLocationId(c.Context(), locationID, month, year)
	if err != nil {
		return errorx.WithDetails(errorx.ErrInternal, err.Error())
	}

	return response.Success(c, resp)
}

func (h *DoctorHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateDoctorDTO
	if err := c.BodyParser(&req); err != nil {
		return errorx.WithDetails(errorx.ErrInvalidRequest, "Geçersiz giriş formatı")
	}

	if err := h.service.Create(c.Context(), &req); err != nil {
		return errorx.WithDetails(errorx.ErrInternal, err.Error())
	}

	return response.Success(c, nil, "Doktor başarıyla oluşturuldu")
}

func (h *DoctorHandler) List(c *fiber.Ctx) error {
	resp, err := h.service.List(c.Context())
	if err != nil {
		return errorx.WithDetails(errorx.ErrInternal, err.Error())
	}

	return response.Success(c, resp)
}

func (h *DoctorHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	resp, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return errorx.WithDetails(errorx.ErrNotFound, "Doktor bulunamadı")
	}

	return response.Success(c, resp)
}

func (h *DoctorHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	var req dto.CreateDoctorDTO
	if err := c.BodyParser(&req); err != nil {
		return errorx.WithDetails(errorx.ErrInvalidRequest, "Geçersiz giriş formatı")
	}

	if err := h.service.Update(c.Context(), id, &req); err != nil {
		return errorx.WithDetails(errorx.ErrInternal, err.Error())
	}

	return response.Success(c, nil, "Doktor başarıyla güncellendi")
}

func (h *DoctorHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return errorx.ErrInvalidRequest
	}

	if err := h.service.Delete(c.Context(), id); err != nil {
		return errorx.WithDetails(errorx.ErrInternal, err.Error())
	}

	return response.Success(c, nil, "Doktor başarıyla silindi")
}
