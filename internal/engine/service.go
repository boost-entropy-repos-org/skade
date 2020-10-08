package engine

type AnalysisService interface {
	Scan(fileName string) (*Report, error)
}
