package utils

import "log"

type Logger struct {
}

func (logger *Logger) Printf(s string, params ...interface{}) {
	log.Printf(s, params...)
}
