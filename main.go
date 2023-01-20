package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kahfikurniaaji/belajar-golang-dependency-injection/helper"
	"github.com/kahfikurniaaji/belajar-golang-dependency-injection/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
