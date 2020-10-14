package engine

type Interactor interface {
	ParseFile(FileName string) ([]byte, error)
	User(Name string)
}
