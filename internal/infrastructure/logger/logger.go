package logger

import (
	"os"

	kitlogger "github.com/go-kit/log"
)

func NewLogger() Logger {
	w := kitlogger.NewSyncWriter(os.Stderr)
	l := kitlogger.NewLogfmtLogger(w)
	l = kitlogger.With(l, "sb", kitlogger.DefaultTimestampUTC, "caller", kitlogger.DefaultCaller)

	return l
}
