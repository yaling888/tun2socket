package log

type Logger interface {
	D(format string, args ...interface{})
	I(format string, args ...interface{})
	W(format string, args ...interface{})
	E(format string, args ...interface{})
}

var logger Logger

func SetLogger(log Logger) {
	logger = log
}

func D(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.D(format, args...)
}

func I(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.I(format, args...)
}

func W(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.W(format, args...)
}

func E(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	logger.E(format, args...)
}
