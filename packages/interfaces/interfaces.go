package interfaces

type Validtator interface {
	IsValid(sequence DigitSequence) bool
}

type DigitSequence interface {
	SetSequence(sequence string)
	GetSequence() []int
}

type Configuration interface {
	GetConfiguration() map[string]string
}
