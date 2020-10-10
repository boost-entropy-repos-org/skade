package engine

import (
	"os"
)

func (a *analysisService) Scan(FileName string) (*Report, error) {
	a.logger.Debug("testing")
	var err error
	//start by opening the file
	suspiciousFile, err = os.Open(FileName)
	if err != nil {
		a.logger.Error("Could not open file")
	}

	//then get all the susBytes from the file
	susBytes, err = getBytesFromFile()
	if err != nil {
		a.logger.Error("failed to read get susBytes from file, this is fatal")
	}
	defer suspiciousFile.Close()

	isPe, err := a.isPeFile()
	if err != nil {
		a.logger.Error("error or sth")
	}
	if isPe {
		a.logger.Info("yay a pe")
	} else {
		a.logger.Info("nay not a pe")
	}
	return nil, nil
}
