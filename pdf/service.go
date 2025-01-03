package service

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
)

type PdfRepository interface {
	Compress(inputPath, outputPath string) error
	Merge(inputPaths []string, outputPath string) error
	Split(inputPath, outputPath string) error
}

type Service struct {
	pdfRepo PdfRepository
}

func NewService(p PdfRepository) *Service {
	return &Service{pdfRepo: p}
}

func (s *Service) CompressPDF(ctx context.Context, inputPath, outputPath string) error {
	return s.pdfRepo.Compress(inputPath, outputPath)
}

func (s *Service) MergePDF(ctx context.Context, req domain.MergePDF) error {
	return s.pdfRepo.Merge(req.FilePaths, req.Output)
}

func (s *Service) SplitPDF(ctx context.Context, req domain.SplitPDF) error {
	return s.pdfRepo.Split(req.FilePath, req.OutputDir)
}
