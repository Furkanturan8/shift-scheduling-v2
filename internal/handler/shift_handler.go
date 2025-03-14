package handler

import (
	"shift-scheduling-v2/internal/dto"
	"shift-scheduling-v2/internal/service"
	"shift-scheduling-v2/pkg/errorx"
	"shift-scheduling-v2/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ShiftHandler struct {
	shiftService *service.ShiftService
}

func NewShiftHandler(s *service.ShiftService) *ShiftHandler {
	return &ShiftHandler{shiftService: s}
}

func (h *ShiftHandler) Create(ctx *fiber.Ctx) error {
	var vm dto.ShiftCreateRequest
	if err := ctx.BodyParser(&vm); err != nil {
		return errorx.ErrInvalidRequest
	}

	err := h.shiftService.CreateShift(ctx.Context(), vm)
	if err != nil {
		return err
	}

	return response.Success(ctx, nil, "Shift created successfully")
}

// Diğer handler fonksiyonları...
