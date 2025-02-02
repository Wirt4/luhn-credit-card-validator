package factories

import (
	"main.go/packages/credit_card"
	"main.go/packages/error_handlers"
	"main.go/packages/interfaces"
	"main.go/packages/routing_number"
	"main.go/packages/types"
)

func ErrorHandlerFactory() interfaces.ErrorHandler[types.RequestPayload] {
	return error_handlers.NewErrorHandler()
}

func ContainerFactory(t string) interfaces.DigitSequence {
	switch t {
	case "luhn":
		return &credit_card.CreditCard{}
	case "routing_checksum":
		return &routing_number.RoutingNumber{}
	}
	return nil
}
