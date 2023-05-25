package main

import (
	"golang-resful-api/app"
	"golang-resful-api/controller"
	"golang-resful-api/helper"
	"golang-resful-api/middleware"
	"golang-resful-api/repository"
	"golang-resful-api/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:8000",

		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
