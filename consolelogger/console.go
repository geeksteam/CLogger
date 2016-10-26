package consolelogger

import (
	"log"
	"os"

	"github.com/geeksteam/klogger"
)

type (
	consoleLogger struct {
		prefix string

		debug    *log.Logger
		info     *log.Logger
		warning  *log.Logger
		error    *log.Logger
		critical *log.Logger
	}
)

const (
	CLR_0 = "\x1b[30;1m"
	CLR_R = "\x1b[31;1m"
	CLR_G = "\x1b[32;1m"
	CLR_Y = "\x1b[33;1m"
	CLR_B = "\x1b[34;1m"
	CLR_M = "\x1b[35;1m"
	CLR_C = "\x1b[36;1m"
	CLR_W = "\x1b[37;1m"
	CLR_N = "\x1b[0m"
)

func NewConsoleLogger() klogger.Logger {
	return &consoleLogger{
		"",
		log.New(os.Stdout, CLR_0, log.Ldate|log.Ltime),
		log.New(os.Stdout, CLR_G, log.Ldate|log.Ltime),
		log.New(os.Stdout, CLR_Y, log.Ldate|log.Ltime),
		log.New(os.Stdout, CLR_R, log.Ldate|log.Ltime),
		log.New(os.Stdout, CLR_C, log.Ldate|log.Ltime),
	}
}

func NewPrefixLogger(prefix string) klogger.Logger {
	return &consoleLogger{
		prefix,
		log.New(os.Stdout, CLR_0, log.Ldate|log.Ltime),
		log.New(os.Stdout, CLR_G, log.Ldate|log.Ltime),
		log.New(os.Stdout, CLR_Y, log.Ldate|log.Ltime),
		log.New(os.Stdout, CLR_R, log.Ldate|log.Ltime),
		log.New(os.Stdout, CLR_C, log.Ldate|log.Ltime),
	}
}

func (logger *consoleLogger) Debug(format string, args ...interface{}) {
	logger.debug.Printf(logger.prefix+"DEBUG: "+format+CLR_N, args...)
}

func (logger *consoleLogger) Info(format string, args ...interface{}) {
	logger.info.Printf(logger.prefix+"INFO:  "+format+CLR_N, args...)
}

func (logger *consoleLogger) Warning(format string, args ...interface{}) {
	logger.warning.Printf(logger.prefix+"WARN:  "+format+CLR_N, args...)
}

func (logger *consoleLogger) Error(format string, args ...interface{}) {
	logger.error.Printf(logger.prefix+"ERROR: "+format+CLR_N, args...)
}

func (logger *consoleLogger) Critical(format string, args ...interface{}) {
	logger.critical.Printf(logger.prefix+"FATAL: "+format+CLR_N, args...)
}

func init() {
	klogger.New = NewConsoleLogger
	klogger.NewWithPrefix = NewPrefixLogger
}
