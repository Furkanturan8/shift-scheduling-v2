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
	app         *fiber.App
	authHandler *handler.AuthHandler
	userHandler *handler.UserHandler
	// Diğer handler'lar buraya eklenecek
}

func NewRouter(authHandler *handler.AuthHandler, userHandler *handler.UserHandler) *Router {
	return &Router{
		app:         fiber.New(),
		authHandler: authHandler,
		userHandler: userHandler,
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

	// Diğer route grupları buraya eklenecek
}

func (r *Router) GetApp() *fiber.App {
	return r.app
}
