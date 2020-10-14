package domain

// An interactor is a way for our engine to receive files
// the file might come from a webserver, it might also come from a cli
type Interactor interface {
	GetFile(FileName string) ([]byte, error)
	AuthenticateUser(Name string)
}
