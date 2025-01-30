package fileId

type request struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func newRequest() *request {
	return &request{}
}

type response struct {
	Success bool `json:"success"`
}

func newResponse(success bool) *response {
	return &response{
		success,
	}
}
