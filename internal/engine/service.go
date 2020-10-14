package engine

import (
	"bufio"
	//"encoding/hex"
	"os"
)

// Package wide Variables
var (
	FileName       string
	suspiciousFile *os.File
	susBytes       []byte
)

type analysisService struct {
	logger     Logger
	repo       Repository
	interactor Interactor
}

type AnalysisService interface {
	Scan(fileName string) (*Report, error)
}

func NewAnalysisService(logger Logger, repo Repository, interactor Interactor) AnalysisService {
	return &analysisService{
		logger,
		repo,
		interactor,
	}
}

func getBytesFromFile() ([]byte, error) {
	stats, err := suspiciousFile.Stat()
	if err != nil {
		return nil, err
	}
	var size int64 = stats.Size()
	susBytes := make([]byte, size)

	buffer := bufio.NewReader(suspiciousFile)
	_, err = buffer.Read(susBytes)

	return susBytes, err
}
