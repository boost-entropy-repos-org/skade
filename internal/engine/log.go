package engine

type Logger interface {
	Debug(message string)
	Info(message string)
	Error(message string)
}
