package interfaces

import (
	"io"
	"net/http"

	"main.go/packages/types"
)

type Validator interface {
	IsValid(sequence CreditCardInterface) (bool, error)
}

type CreditCardInterface interface {
	SetSequence(sequence string) error
	GetSequence() []int
	HasCorrectLength() bool
	Issuers() []types.CardIssuer
}

//TODO: Either rename the interface or compose it with Card Interface

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
	Create() CreditCardInterface
}

type ErrorHandlerFactory interface {
	Create() ErrorHandler[types.CreditCardRequest]
}
