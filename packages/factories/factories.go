package factories

import (
	"main.go/packages/credit_card"
	"main.go/packages/error_handlers"
	"main.go/packages/interfaces"
	"main.go/packages/types"
)

type ErrorHandlerFactory struct {
}

func (*ErrorHandlerFactory) Create() interfaces.ErrorHandler[types.CreditCardRequest] {
	return error_handlers.NewErrorHandler()
}

type CreditCardFactory struct{}

func (c *CreditCardFactory) NewCreditCard() interfaces.DigitSequence {
	return credit_card.NewCreditCard()
}
