package utils

type singleton struct {
	Logger *logger
}

var Instance *singleton = nil

func New() *singleton {
	if Instance == nil {
		Instance = &singleton{&logger{}}
	}
	return Instance
}
