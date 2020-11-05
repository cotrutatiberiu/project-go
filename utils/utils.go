package utils

type Singleton struct {
	Logger *Logger
}

var Instance *Singleton = nil

func New() *Singleton {
	if Instance == nil {
		Instance = &Singleton{&Logger{}}
	}
	return Instance
}
