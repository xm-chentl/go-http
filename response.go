package gohttp

type HttpResponse struct {
	Body  []byte
	Code  int
	Error error
}
