package file

type responseJson struct {
	ID string `json:"id"`
}

func newResponseJson(id string) *responseJson {
	return &responseJson{
		ID: id,
	}
}
