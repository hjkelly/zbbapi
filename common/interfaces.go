package common

type SelfValidator interface {
	Validate() *ValidationError
}
