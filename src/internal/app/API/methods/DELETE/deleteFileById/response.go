package deleteFileById

type response struct {
	Success bool `json:"success"`
}

func newResponse(success bool) *response {
	return &response{
		success,
	}
}
