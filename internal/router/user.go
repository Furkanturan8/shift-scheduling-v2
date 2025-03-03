package router

import (
	"shift-scheduling-V2/internal/api"

	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

type UserRouter interface {
	Init(root *fiber.Router, authzMiddleware *casbin.Middleware)
}

type userRouter struct {
	api api.UserHttpApi
}

func NewUserRouter(api api.UserHttpApi) UserRouter {
	return &userRouter{api: api}
}

func (mr *userRouter) Init(root *fiber.Router, authzMiddleware *casbin.Middleware) {
	router := (*root).Group("/user")
	{
		// commands

		router.Post("", authzMiddleware.RequiresPermissions([]string{"use_service:access"}, casbin.WithValidationRule(casbin.MatchAllRule)), mr.api.AddUser)
		router.Put("/:id", mr.api.UpdateUser)
		router.Delete("/:id", mr.api.DeleteUser)
		// queries
		router.Get("", mr.api.GetUsers)
		router.Get("/:id", mr.api.GetUser)
	}

}
