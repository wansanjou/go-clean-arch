package repository

import "github.com/pdfcpu/pdfcpu/pkg/api"

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
