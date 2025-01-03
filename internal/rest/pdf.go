package rest

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/gofiber/fiber/v2"
)

type PdfService interface {
	CompressPDF(ctx context.Context, inputPath, outputPath string) error
	MergePDF(ctx context.Context, req domain.MergePDF) error
	SplitPDF(ctx context.Context, req domain.SplitPDF) error
}

type PDFHandler struct {
	Service PdfService
}

func NewPDFHandler(app *fiber.App, svc PdfService) *PDFHandler {
	handler := &PDFHandler{
		Service: svc,
	}

	app.Post("/pdf/compress", handler.CompressPDF)
	app.Post("/pdf/merge", handler.MergePDF)
	app.Post("/pdf/split", handler.SplitPDF)

	return handler
}

func (h *PDFHandler) CompressPDF(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get file",
		})
	}

	outputDir := "./result"
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create output directory: %v", err),
		})
	}

	inputPath := filepath.Join(outputDir, file.Filename)
	outputPath := filepath.Join(outputDir, "compressed_"+file.Filename)

	if err := c.SaveFile(file, inputPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}
	defer os.Remove(inputPath)

	err = h.Service.CompressPDF(c.Context(), inputPath, outputPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to compress PDF: %v", err),
		})
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Compressed file not found",
		})
	}

	return c.SendFile(outputPath)
}

func (h *PDFHandler) MergePDF(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse multipart form",
		})
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No files uploaded",
		})
	}

	outputDir := "./result"
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create output directory: %v", err),
		})
	}

	var inputPaths []string
	for _, file := range files {
		inputPath := filepath.Join(outputDir, file.Filename)
		inputPaths = append(inputPaths, inputPath)

		if err := c.SaveFile(file, inputPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save file",
			})
		}
		defer os.Remove(inputPath)
	}

	outputPath := filepath.Join(outputDir, "merged.pdf")

	err = h.Service.MergePDF(c.Context(), domain.MergePDF{FilePaths: inputPaths, Output: outputPath})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to merge PDFs",
		})
	}

	return c.SendFile(outputPath)
}

func (h *PDFHandler) SplitPDF(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get file",
		})
	}

	outputDir := "./result"
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create output directory: %v", err),
		})
	}

	inputPath := filepath.Join(outputDir, file.Filename)

	if err := c.SaveFile(file, inputPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}
	defer os.Remove(inputPath)

	err = h.Service.SplitPDF(c.Context(), domain.SplitPDF{FilePath: inputPath, OutputDir: outputDir})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to split PDF: %v", err),
		})
	}

	files, err := filepath.Glob(filepath.Join(outputDir, "*_page_1.pdf"))
	if err != nil || len(files) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Split file 'page_1.pdf' was not created",
		})
	}

	return c.SendFile(files[0])
}
