package types

type CreditCardRequest struct {
	CreditCardNumber string
}

type CreditCardResponse struct {
	Issuer string
	Valid  bool
}
