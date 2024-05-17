package main

import (
	"github.com/catosaurusrex2003/golang_crud/bootstrap"
	"github.com/catosaurusrex2003/golang_crud/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitializeApp(app)
}
