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

type ProviderData struct {
	Name              string
	IINs              []int
	MaxSequenceLength int
	MinSequenceLength int
}

type Node struct {
	Children map[int]*Node
	Data     *CardIssuer
}
