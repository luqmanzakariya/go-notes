package main

import (
	// "github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	// "programmerzamannow/belajar-golang-restful-api/app"
	// "programmerzamannow/belajar-golang-restful-api/controller"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/middleware"
	// "programmerzamannow/belajar-golang-restful-api/repository"
	// "programmerzamannow/belajar-golang-restful-api/service"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
