package servers

import (
	"github.com/forgetz/go-kawaii-shop/modules/middlewares/middlewaresHandlers"
	"github.com/forgetz/go-kawaii-shop/modules/middlewares/middlewaresRepositories"
	"github.com/forgetz/go-kawaii-shop/modules/middlewares/middlewaresUsecases"
	"github.com/forgetz/go-kawaii-shop/modules/monitor/monitorHandlers"
	"github.com/forgetz/go-kawaii-shop/modules/users/usersHandlers"
	"github.com/forgetz/go-kawaii-shop/modules/users/usersRepositories"
	"github.com/forgetz/go-kawaii-shop/modules/users/usersUsecases"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
	UserModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)
	m.r.Get("/", handler.HealthCheck)

}

func (m *moduleFactory) UserModule() {
	repository := usersRepositories.UsersRepository(m.s.db)
	usecase := usersUsecases.UsersUsecase(m.s.cfg, repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)

	// group user
	// /v1/users/xxx
	router := m.r.Group("/users")

	router.Post("/signup", handler.SignUpCustomer)
}
