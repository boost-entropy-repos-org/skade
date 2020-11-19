package engine

import (
	//"bufio"
	//"encoding/hex"
	"io"
)

// Package wide Variables
var (
	FileName       string
	suspiciousFile io.Reader
	susBytes       []byte
)

type analysisService struct {
	logger 	Logger
	repo	Repository
}

type AnalysisService interface {
	//Application logic
	Scan(fileName string) (*Report, error)
	ScanBytes(susBytes []byte) (*Report, error)
	ScanFile(file io.Reader) (*Report, error)
	//Database Operations
	RegisterUser(username string, email string) error
	ForgotPassword() error
	ChangePassword()
	IsValid()
}

func NewAnalysisService(logger Logger, repo Repository) AnalysisService {
	return &analysisService{
		logger,
		repo,
	}
}

//func getBytesFromFile() ([]byte, error) {
//	stats, err := suspiciousFile.Stat()
//	if err != nil {
//		return nil, err
//	}
//	var size int64 = stats.Size()
//	susBytes := make([]byte, size)
//
//	buffer := bufio.NewReader(suspiciousFile)
//	_, err = buffer.Read(susBytes)
//
//	return susBytes, err
//}
