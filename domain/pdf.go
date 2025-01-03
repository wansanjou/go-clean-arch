package domain

type CompressPDF struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type MergePDF struct {
	FilePaths []string `json:"filePaths"`
	Output    string   `json:"output"`
}

type SplitPDF struct {
	FilePath   string `json:"filePath"`
	OutputDir  string `json:"outputDir"`
	SplitPages []int  `json:"splitPages"`
}
