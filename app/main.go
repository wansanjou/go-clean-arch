package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/bxcodec/go-clean-arch/docs"
	"github.com/bxcodec/go-clean-arch/internal/repository"
	"github.com/bxcodec/go-clean-arch/internal/rest"
	service "github.com/bxcodec/go-clean-arch/pdf"
)

// @title PDF API
// @version 1.0
// @description This is a sample server for a PDF API.
// @host localhost:9090
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	pdfRepo := repository.NewPDFRepository()
	svc := service.NewService(pdfRepo)
	rest.NewPDFHandler(app, svc)

	address := ":9090"
	log.Fatal(app.Listen(address))
}
