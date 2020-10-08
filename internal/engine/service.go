package engine

type MalwareAnalysis interface {
	Scan(fileName string) (*Report, error)
}
