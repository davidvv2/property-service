package headers

import "net/http"

var _ Header = (*HeaderImpl)(nil)

type HeaderImpl struct {
	header http.Header
}

// Get implements Header.
func (hi *HeaderImpl) Get(key string) string {
	return hi.header.Get(key)
}

// Add implements Header.
func (hi *HeaderImpl) Add(key, value string) {
	hi.header.Add(key, value)
}

// List implements Header.
func (hi *HeaderImpl) List() map[string][]string {
	return hi.header
}

// Remove implements Header.
func (hi *HeaderImpl) Remove(key string) {
	hi.header.Del(key)
}

func (hi *HeaderImpl) HTTPHeader() http.Header {
	return hi.header
}
