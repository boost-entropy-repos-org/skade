package engine

// An adapter for datastorage
// will add support for postgresql first,
// maybe some other solutions follow
type ReportStorage interface {
	Find(FileName string) (*Report, error)
	Store(report *Report) error
}
