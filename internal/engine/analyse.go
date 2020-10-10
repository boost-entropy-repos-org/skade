package engine

import (
	"bufio"
	//"encoding/hex"
	"errors"
	"os"
)

var (
	ErrUnkownFileType = errors.New("Unkown File Type")
	ErrFileNotFound   = errors.New("File Not found")
)

var (
	FileName       string
	suspiciousFile *os.File
	susBytes       []byte
)

type analysisService struct {
	logger Logger
}

func NewAnalysisService(logger Logger) AnalysisService {
	return &analysisService{
		logger,
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
