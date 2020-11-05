package httpresponse

type HttpResponseSuccess struct {
	status      int
	location    string
	description string
	payload     interface{}
}

type HttpResponseReject struct {
	status      int
	location    string
	description string
	payload     interface{}
	err         error
}

// type Methods interface {
// 	setSuccessPayload(status int, location string, description string, payload interface{})
// 	setRejectPayload(status int, location string, description string, e error)
// }