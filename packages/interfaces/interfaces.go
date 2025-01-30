package interfaces

type Validator interface {
	IsValid(sequence DigitSequence) bool
}

type DigitSequence interface {
	SetSequence(sequence string)
	GetSequence() []int
	HasCorrectLength() bool
}

//TODO: implement Configuration Interface
