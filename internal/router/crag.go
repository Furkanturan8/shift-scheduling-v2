package router

import (
	"shift-scheduling-V2/internal/api"

	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
)

type CragRouter interface {
	Init(root *fiber.Router, authzMiddleware *casbin.Middleware)
}

type cragRouter struct {
	api api.CragHttpApi
}

func NewCragRouter(api api.CragHttpApi) CragRouter {
	return &cragRouter{api: api}
}

func (mr *cragRouter) Init(root *fiber.Router, authzMiddleware *casbin.Middleware) {
	router := (*root).Group("/crag")
	{
		// commands

		router.Post("", authzMiddleware.RequiresPermissions([]string{"use_service:access"}, casbin.WithValidationRule(casbin.MatchAllRule)), mr.api.AddCrag)
		router.Put("/:id", mr.api.UpdateCrag)
		router.Delete("/:id", mr.api.DeleteCrag)
		// queries
		router.Get("", mr.api.GetCrags)
		router.Get("/:id", mr.api.GetCrag)
	}

}
