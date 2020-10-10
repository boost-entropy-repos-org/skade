package engine

import (
	"bufio"
	//"encoding/hex"
	"errors"
	"fmt"
	"log"
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

func (a *analysisService) Scan(FileName string) (*Report, error) {
	var err error
	//start by opening the file
	suspiciousFile, err = os.Open(FileName)
	if err != nil {
		log.Fatal("Could not open file")
	}

	//then get all the susBytes from the file
	susBytes, err = getBytesFromFile()
	if err != nil {
		log.Fatal("failed to read get susBytes from file, this is fatal")
	}
	defer suspiciousFile.Close()

	isPe, err := isPeFile()
	if err != nil {
		log.Fatal("error or sth")
	}
	if isPe {
		fmt.Println("yay a pe")
	} else {
		fmt.Println("nay not a pe")
	}
	thisdoesnotmatter := susBytes[0:8]
	fmt.Println(thisdoesnotmatter)
	return nil, nil
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
