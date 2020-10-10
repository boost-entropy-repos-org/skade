package engine

import (
	"fmt"
)

var (
	offset int
)

func (a *analysisService) isPeFile() (bool, error) {
	a.logger.Debug("isPe test")
	MZ, err := checkMzSignature()
	if err != nil {
		return false, err
	}
	//if no MZ Signature is detected, this is not a PE file
	if MZ != true {
		return false, nil
	}
	getOffset()
	return true, nil
}

//here we check if the MS-DOS Stub is present and compliant
func checkDosStub() (bool, error) {
	stubbytes := susBytes[0:8]
	fmt.Println(stubbytes)
	return true, nil
}

// Every Windows executeable has to start with MZ for the first 2 bytes
// Thus, if this file does not start with MZ it cannot be a pe file
func checkMzSignature() (bool, error) {
	MZ_bytes := susBytes[0:2]
	fmt.Println(MZ_bytes)
	// 77 Decimal Ascii representation of M
	if MZ_bytes[0] != 77 {
		return false, nil
	}
	// 90 Decimal Ascii representation of M
	if MZ_bytes[1] != 90 {
		return false, nil
	}
	return true, nil
}

func getOffset() {
	// the offset of a pe file is always placed at position 0x3c (or 60 in Decimal)
	// we will need this offset for to identify other parts of the pe format correctly
	offset_byte := susBytes[60]
	// since endianness does not matter for a single byte
	// we can convert from byte to int like this
	offset = int(offset_byte)
	//	Logger.Debug("testing here")
	fmt.Println("offset bytes:")
	fmt.Println(offset_byte)
	fmt.Println("offset:")
	fmt.Println(offset)
}
