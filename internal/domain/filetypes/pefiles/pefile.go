package pefiles

import ()

var (
	offset int
)

// Check if the File is in valid PE format
func IsPeFile(susBytes []byte) (bool, error) {
	MZ, err := checkMzSignature(susBytes)
	if err != nil {
		return false, err
	}
	//if no MZ Signature is detected, this is not a PE file
	if MZ != true {
		return false, nil
	}

	// get the offset (from 0x3c) needed for further checks
	getOffset(susBytes)

	//next we check if the MS-DOS Stub is present
	dosStub, err := checkDosStub(susBytes)
	if err != nil {
		return false, err
	}

	// the DOS Stub is mandatory for a PE file,
	// if its missing the file might be corrupted
	if dosStub != true {
		return false, nil
	}
	return true, nil
}

//here we check if the MS-DOS Stub is present and compliant
func checkDosStub(susBytes []byte) (bool, error) {
	// for simplicity sake, for now we only look at the string and not the full dos stub
	stubbytes := susBytes[78:120]
	dosString := string(stubbytes)
	// technically, a program can have a different DOS stub and be a valid PE file
	// that would be considered an annomaly tho
	if dosString != "This program cannot be run in DOS mode" {
		return false, nil
	}
	return true, nil
}

// Every Windows executeable has to start with MZ for the first 2 bytes
// Thus, if this file does not start with MZ it cannot be a pe file
func checkMzSignature(susBytes []byte) (bool, error) {
	MZ_bytes := susBytes[0:2]
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

func getOffset(susBytes []byte) {
	// the offset of a pe file is always placed at position 0x3c (or 60 in Decimal)
	// we will need this offset for to identify other parts of the pe format correctly
	offset_byte := susBytes[60]
	// since endianness does not matter for a single byte
	// we can convert from byte to int like this
	offset = int(offset_byte)
}
