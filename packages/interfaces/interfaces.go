package interfaces

type Validtator interface {
	Validate(sequence DigitSequence) bool
}

type DigitSequence interface {
	SetSequence(sequence string)
	GetSequence() []int
}

//TODO: implement Configuration Interface
