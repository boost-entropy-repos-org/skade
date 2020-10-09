package engine

import (
	"fmt"
)

func isPeFile() (bool, error) {
	MZ, err := check_MZ_signatur()
	if err != nil {
		return false, err
	}
	//if no MZ Signature is detected, this is not a PE file
	if MZ != true {
		return false, nil
	}
	return true, nil
}

//here we check if the MS-DOS Stub is present and compliant
func checkDosStub() (bool, error) {
	stubbytes := bytes[0:8]
	fmt.Println(stubbytes)
	return true, nil
}

// Every Windows executeable has to start with MZ for the first 2 bytes
// Thus, if this file does not start with MZ it cannot be a pe file
func check_MZ_signatur() (bool, error) {
	MZ_bytes := bytes[0:2]
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
