//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"go-restful/app"
	"go-restful/controller"
	"go-restful/middleware"
	"go-restful/repository"
	"go-restful/service"
	"net/http"
)

//var CategorySet = wire.NewSet(
//	repository.NewCategoryRepository,
//	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
//	service.NewCategoryService,
//	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
//	controller.NewCategoryController,
//	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
//)

var CategorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(app.NewDb,
		validator.New,
		CategorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
