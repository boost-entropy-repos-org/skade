package engine

type ReportStorage interface {
	Find(FileName string) (*Report, error)
	Store(report *Report) error
}
