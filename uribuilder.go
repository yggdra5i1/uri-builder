package uribuilder

import (
	"net/url"
)

type builder struct {
	uri  string
	safe bool
}

type InvalidURI struct {
	err string
}

func (e *InvalidURI) Error() string {
	return e.err
}

func New(uri string, safe bool) *builder {
	return &builder{
		uri:  uri,
		safe: safe,
	}
}

func (b *builder) String() string {
	return b.uri
}

func (b *builder) ToURI() *url.URL {
	u, err := url.Parse(b.uri)
	if err != nil {
		panic(&InvalidURI{err: err.Error()})
	}
	return u
}
