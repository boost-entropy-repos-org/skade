package engine

import (
	"errors"
)

var (
	ErrUnkownFileType = errors.New("Unkown File Type")
	ErrFileNotFound   = errors.New("File Not found")
)

type analysisService struct {
	reportStorage ReportStorage
}

func NewAnalysisService(reportStorage ReportStorage) AnalysisService {
	return &analysisService{
		reportStorage,
	}
}

func (a *analysisService) Scan(FileName string) (*Report, error) {
	return nil, nil
}
