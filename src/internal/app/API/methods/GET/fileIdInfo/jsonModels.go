package fileIdInfo

type response struct {
	ID                string `json:"id"`
	Filename          string `json:"filename"`
	Processing        bool   `json:"processing"`
	ProcessingSuccess bool   `json:"processing_success"`
}

func newResponse(id string, filename string, processing bool, processingsuccess bool) *response {
	return &response{
		ID:                id,
		Filename:          filename,
		Processing:        processing,
		ProcessingSuccess: processingsuccess,
	}
}
