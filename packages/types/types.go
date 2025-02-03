package types

type CreditCardRequest struct {
	CreditCardNumber string
}

type CreditCardResponse struct {
	ValidCreditCardNumber bool
}
