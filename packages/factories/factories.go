package factories

import (
	"main.go/packages/credit_card"
	"main.go/packages/error_handlers"
	"main.go/packages/interfaces"
	"main.go/packages/types"
)

func ErrorHandlerFactory() interfaces.ErrorHandler[types.CreditCardPayload] {
	return error_handlers.NewErrorHandler()
}

func CreditCardFactory() interfaces.DigitSequence {
	return &credit_card.CreditCard{}
}
