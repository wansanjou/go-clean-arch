package service

// func TestCompressPDF(t *testing.T) {
// 	mockRepo := new(mocks.PdfRepository)
// 	mockRepo.On("Compress", "input.pdf", "output.pdf").Return(nil)
// 	service := NewService(mockRepo)
// 	ctx := context.Background()
// 	err := service.CompressPDF(ctx, "input.pdf", "output.pdf")
// 	assert.NoError(t, err)

// 	mockRepo.AssertExpectations(t)
// }

// func TestMergePDF(t *testing.T) {
// 	mockRepo := new(mocks.PdfRepository)
// 	req := domain.MergePDF{
// 		FilePaths: []string{"file1.pdf", "file2.pdf"},
// 		Output:    "merged.pdf",
// 	}
// 	mockRepo.On("Merge", req.FilePaths, req.Output).Return(nil)
// 	service := NewService(mockRepo)
// 	ctx := context.Background()
// 	err := service.MergePDF(ctx, req)
// 	assert.NoError(t, err)

// 	mockRepo.AssertExpectations(t)
// }

// func TestSplitPDF(t *testing.T) {
// 	mockRepo := new(mocks.PdfRepository)
// 	req := domain.SplitPDF{
// 		FilePath:  "input.pdf",
// 		OutputDir: "output_dir",
// 	}
// 	mockRepo.On("Split", req.FilePath, req.OutputDir).Return(nil)
// 	service := NewService(mockRepo)
// 	ctx := context.Background()
// 	err := service.SplitPDF(ctx, req)
// 	assert.NoError(t, err)

// 	mockRepo.AssertExpectations(t)
// }
