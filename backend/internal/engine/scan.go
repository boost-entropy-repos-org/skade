package engine

import (
	"io"
	"mime/multipart"
	"os"
)

// the main 'scan' function, will basically be a wrapper around the other functionalities
func (a *analysisService) Scan(FileName string) (*Report, error) {
	a.logger.Debug("Skade Entrypoint")
	var err error
	//start by opening the file
	suspiciousFile, err = os.Open(FileName)
	if err != nil {
		a.logger.Error("Could not open file:")
	}

	// make sure we close the file again once we are done
	// defer suspiciousFile.Close()

	// then get all the Bytes from the file
	// using susBytes because 'bytes' is libarry name
	// susBytes, err = getBytesFromFile()
	if err != nil {
		a.logger.Error("Could not get raw bytes from file")
	}

	return nil, nil
}

func (a *analysisService) ScanBytes(susBytes []byte) (*Report, error) {
	a.logger.Debug("Starting Scan")
	return nil, nil
}

func (a *analysisService) ScanFile(file io.Reader) (*Report, error) {
	return nil, nil
}

func (a *analysisService) ScanFileUpload(file multipart.File) (*Report, error) {
	return nil, nil
}
