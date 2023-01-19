package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kahfikurniaaji/belajar-golang-restful-api/app"
	"github.com/kahfikurniaaji/belajar-golang-restful-api/controller"
	"github.com/kahfikurniaaji/belajar-golang-restful-api/helper"
	"github.com/kahfikurniaaji/belajar-golang-restful-api/middleware"
	"github.com/kahfikurniaaji/belajar-golang-restful-api/repository"
	"github.com/kahfikurniaaji/belajar-golang-restful-api/service"
	"net/http"
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
