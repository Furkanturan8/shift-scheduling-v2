package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"shift-scheduling-v2/internal/handler"
	"shift-scheduling-v2/internal/middleware"
)

type Router struct {
	app           *fiber.App
	authHandler   *handler.AuthHandler
	userHandler   *handler.UserHandler
	doctorHandler *handler.DoctorHandler
	shiftHandler  *handler.ShiftHandler
	// Diğer handler'lar buraya eklenecek
}

func NewRouter(a *handler.AuthHandler, u *handler.UserHandler, d *handler.DoctorHandler, s *handler.ShiftHandler) *Router {
	return &Router{
		app:           fiber.New(),
		authHandler:   a,
		userHandler:   u,
		doctorHandler: d,
	}
}

func (r *Router) SetupRoutes() {
	// Middleware'leri ekle
	r.app.Use(logger.New())
	r.app.Use(recover.New())
	r.app.Use(cors.New())

	// API versiyonu
	api := r.app.Group("/api")
	v1 := api.Group("/v1")

	// Auth routes
	auth := v1.Group("/auth")
	auth.Post("/register", r.authHandler.Register)
	auth.Post("/login", r.authHandler.Login)
	auth.Post("/refresh", r.authHandler.RefreshToken)
	auth.Post("/forgot-password", r.authHandler.ForgotPassword)
	auth.Post("/reset-password", r.authHandler.ResetPassword)
	auth.Post("/logout", middleware.AuthMiddleware(), r.authHandler.Logout)

	// User routes - Base group
	users := v1.Group("/users")

	// Normal user routes (profil yönetimi)
	userProfile := users.Group("/me")
	userProfile.Use(middleware.AuthMiddleware()) // Sadece authentication gerekli
	userProfile.Get("/", r.userHandler.GetProfile)
	userProfile.Put("/", r.userHandler.UpdateProfile)

	// Admin only routes
	adminUsers := users.Group("/")
	adminUsers.Use(middleware.AuthMiddleware(), middleware.AdminOnly()) // Admin yetkisi gerekli
	adminUsers.Get("/", r.userHandler.List)
	adminUsers.Get("/:id", r.userHandler.GetByID)
	adminUsers.Put("/:id", r.userHandler.Update)
	adminUsers.Delete("/:id", r.userHandler.Delete)

	// Doctor routes
	doctors := v1.Group("/doctors")
	adminDoctors := doctors.Group("/")
	adminDoctors.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	adminDoctors.Get("/", r.doctorHandler.List)
	adminDoctors.Get("/:id", r.doctorHandler.GetByID)
	adminDoctors.Post("/", r.doctorHandler.Create)
	adminDoctors.Put("/:id", r.doctorHandler.Update)
	adminDoctors.Delete("/:id", r.doctorHandler.Delete)
	adminDoctors.Get("/location/:location_id", r.doctorHandler.GetDoctorsByLocation)
	adminDoctors.Get("/:id/holidays", r.doctorHandler.GetDoctorHolidays)
	adminDoctors.Get("/holidays/:location_id", r.doctorHandler.GetDoctorsHolidayByLocationId)
	adminDoctors.Get("/:shift_id", r.doctorHandler.GetDoctorByShiftID)

	// Shift routes
	shifts := v1.Group("/shifts")
	adminShifts := shifts.Group("/")
	adminShifts.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	//adminShifts.Get("/", r.shiftHandler.ListShifts)
	//adminShifts.Get("/:id", r.shiftHandler.GetShiftByID)
	adminShifts.Post("/", r.shiftHandler.Create)

	// Diğer route grupları buraya eklenecek
}

func (r *Router) GetApp() *fiber.App {
	return r.app
}
