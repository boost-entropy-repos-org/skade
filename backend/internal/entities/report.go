package entities

type Report struct {
	ID        int64         `json:"id"`
	Malicious bool          `json:"malicious"`
}