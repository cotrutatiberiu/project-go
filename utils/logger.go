package utils

import "log"

type logger struct {
}

func (logger *logger) Printf(s string, params ...interface{}) {
	log.Printf(s, params...)
}