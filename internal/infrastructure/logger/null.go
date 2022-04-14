package logger

type nullLogger struct {
}

func (nullLogger) Log(...interface{}) error {
	return nil
}

func NewNullLogger() Logger {
	return nullLogger{}
}
