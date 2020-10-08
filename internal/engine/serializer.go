package engine

type ReportSerialiser interface {
	Decode(input []byte) (*Report, error)
	Encode(input *Report) ([]byte, error)
}
