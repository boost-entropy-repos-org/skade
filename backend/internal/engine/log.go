package engine

// Adapter for different loggers
// Currently only using zap
// but we want to follow the hexagon structure
type Logger interface {
	Debug(message string)
	Info(message string)
	Error(message string)
}
