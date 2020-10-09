package engine

import "fmt"

func isPeFile() (bool, error) {
	dosStub, err := checkDosStub()
	if err != nil {
		return false, err
	}
	//if no DOS Stub is detected, this is not a PE file
	if dosStub != true {
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
