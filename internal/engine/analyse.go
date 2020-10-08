package engine

import (
	"errors"
	"log"
	"os"
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
	testfile := "/home/marius/test.exe"

	// reading the file
	file, err := os.Open(testfile)
	if err != nil {
		log.Fatal("Could not open file: ")
	}
	// make sure we always close the file once we are done scanning it
	defer file.Close()

	return nil, nil
}
