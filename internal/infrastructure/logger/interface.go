package logger

type Logger interface {
	Log(...interface{}) error
}
