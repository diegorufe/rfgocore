package logger

import "log"

// Default log level is debug 0100
var logLevel uint8 = 0x08

// Trace : method for print Trace logger
func Trace(data interface{}) {
	if IsTraceEnabled() {
		log.Println(data)
	}
}

// IsTraceEnabled : method for check trace enabled
func IsTraceEnabled() bool {
	return (logLevel & 0x08) > 0
}

// Debug : method for print Debug logger
func Debug(data interface{}) {
	if IsDebugEnabled() {
		log.Println(data)
	}
}

// IsDebugEnabled : method for check debug enabled
func IsDebugEnabled() bool {
	return (logLevel & 0x0C) > 0
}

// Info : method for print info logger
func Info(data interface{}) {
	if IsInfoEnabled() {
		log.Println(data)
	}
}

// IsInfoEnabled : method for check info enabled
func IsInfoEnabled() bool {
	return (logLevel & 0x0E) > 0
}

// Error : method for print error logger
func Error(data interface{}) {
	if IsErrorEnabled() {
		log.Println(data)
	}
}

// IsErrorEnabled : method for check error enabled
func IsErrorEnabled() bool {
	return (logLevel & 0x0F) > 0
}

// Fatal : method for print fatal logger
func Fatal(data interface{}) {
	log.Fatalln(data)
}
