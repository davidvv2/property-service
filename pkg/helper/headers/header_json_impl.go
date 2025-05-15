package headers

import "net/http"

func NewDefaultJSON() *HeaderImpl {
	header := HeaderImpl{
		header: make(http.Header),
	}
	header.Add("Content-Type", "application/json")
	return &header
}
