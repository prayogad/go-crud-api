package main

import (
	"belajar_golang_restfult_api/app"
	"belajar_golang_restfult_api/controller"
	"belajar_golang_restfult_api/helper"
	"belajar_golang_restfult_api/middleware"
	"belajar_golang_restfult_api/repository"
	"belajar_golang_restfult_api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
