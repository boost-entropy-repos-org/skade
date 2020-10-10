package engine

// Will probably add a lot more here
// goal should be to generate a report
// with as much information about an
// executeable as possible
type Report struct {
	FileName  string
	FileSize  float64
	FileType  string
	Malicious bool
}
