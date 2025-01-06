package repository

import (
	"os"

	"github.com/otiai10/gosseract/v2"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type PDFRepository struct{}

func NewPDFRepository() *PDFRepository {
	return &PDFRepository{}
}

func (r *PDFRepository) Compress(inputPath, outputPath string) error {
	return api.OptimizeFile(inputPath, outputPath, nil)
}

func (r *PDFRepository) Merge(inputPaths []string, outputPath string) error {
	return api.MergeCreateFile(inputPaths, outputPath, false, nil)
}

func (r *PDFRepository) Split(inputPath, outputPath string) error {
	pages := []string{"1", "2"}
	return api.ExtractPagesFile(inputPath, outputPath, pages, nil)
}

func (r *PDFRepository) Ocr(inputPath, outputPath string) error {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(inputPath)
	text, err := client.Text()
	if err != nil {
		return err
	}

	err = os.WriteFile(outputPath, []byte(text), 0644)
	if err != nil {
		return err
	}

	return nil
}
