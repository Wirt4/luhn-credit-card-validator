package interfaces

import (
	"io"
	"net/http"

	"main.go/packages/types"
)

type Validator interface {
	IsValid(sequence DigitSequence) (bool, error)
}

type DigitSequence interface {
	SetSequence(sequence string) error
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

type Visitor interface {
	Traverse(sequence []int, tree *types.Node) //TOOO: set node as type or interface
	GetVisited() []types.CardIssuer
}

type DigitSequenceFactory interface {
	NewCreditCard() DigitSequence
}

type ErrorHandlerFactory interface {
	Create() ErrorHandler[types.CreditCardRequest]
}
