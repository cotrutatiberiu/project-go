package httpresponse

type Success struct {
	Status      int         `json:"status"`
	Location    string      `json:"location"`
	Description string      `json:"description"`
	Payload     interface{} `json:"payload"`
}

type Reject struct {
	Status      int         `json:"status"`
	Location    string      `json:"location"`
	Description string      `json:"description"`
	Payload     interface{} `json:"payload"`
	Error       error       `json:"error"`
}