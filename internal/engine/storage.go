package engine

// An adapter for datastorage
// will add support for postgresql first,
// maybe some other solutions follow
type Repository interface {
	Find(FileName string) (*Report, error)
	Store(report *Report) error
	CreateUser(username string, email string) error
}
