package types

type CreditCardRequest struct {
	CreditCardNumber string
}

type CreditCardResponse struct {
	Issuer string
	Valid  bool
}

type CardIssuer struct {
	Min    int
	Max    int
	Issuer string
}

type Node struct {
	Children map[int]*Node
	Data     *CardIssuer
}
