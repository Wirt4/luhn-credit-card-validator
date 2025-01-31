package interfaces

type Validator interface {
	IsValid(sequence DigitSequence) bool
}

type DigitSequence interface {
	SetSequence(sequence string)
	GetSequence() []int
	HasCorrectLength() bool
}

type Configuration interface {
	GetConfiguration() map[string]string
}
