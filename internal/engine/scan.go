package engine

import (
	"os"
)

// the main 'scan' function, will basically be a wrapper around the other functionalities
func (a *analysisService) Scan(FileName string) (*Report, error) {
	a.logger.Debug("Starting Scan")
	var err error
	//start by opening the file
	suspiciousFile, err = os.Open(FileName)
	if err != nil {
		a.logger.Error("Could not open file:")
	}

	//make sure we close the file again once we are done
	defer suspiciousFile.Close()

	// then get all the Bytes from the file
	// using susBytes because 'bytes' is libarry name
	susBytes, err = getBytesFromFile()
	if err != nil {
		a.logger.Error("Could not get raw bytes from file")
	}

	isPe, err := a.isPeFile()
	if err != nil {
		a.logger.Error("Error during PE check")
	}
	if isPe {
		a.logger.Info("The file has been identified as a PE file")
	} else {
		a.logger.Debug("The file is not a PE file")
	}
	return nil, nil
}
