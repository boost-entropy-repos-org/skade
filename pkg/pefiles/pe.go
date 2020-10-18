package pefiles

func NewPEFile(fileName string) (pefile *PEFile, err error) {
	pefile = new(PEFile)
	pefile.FileName = fileName

	// end of the happy path
	return pefile, nil
}
