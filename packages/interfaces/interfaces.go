package interfaces

import (
	"io"
	"net/http"
)

type Validator interface {
	IsValid(sequence DigitSequence) bool
	Type() string
}

type DigitSequence interface {
	SetSequence(sequence string)
	GetSequence() []int
	HasCorrectLength() bool
}

type Configuration interface {
	GetConfiguration() map[string]string
}

type Handler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type ErrorHandler[T any] interface {
	CheckMethod(method string)
	CheckBody(body io.ReadCloser)
	HasError() bool
	GetMessage() string
	GetCode() int
	GetParsed() T
}
