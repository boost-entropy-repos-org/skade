package engine

import (
	"fmt"
)

func isElfFile() (bool, error) {
	return true, nil
}

// Checking for the .ELF file signtaure of an ELF file
func checkElfSignature() (bool, error) {
	elfBytes := susBytes[0:4]
	fmt.Println(elfBytes)
	// 46 Decimal Ascii representation of .
	if elfBytes[0] != 46 {
		return false, nil
	}
	// 69 Decimal Ascii representation of E
	if elfBytes[1] != 69 {
		return false, nil
	}
	// 76 Decimal Ascii representation of L
	if elfBytes[2] != 76 {
		return false, nil
	}
	// 70 Decimal Ascii representation of F
	if elfBytes[3] != 70 {
		return false, nil
	}
	return true, nil
}
