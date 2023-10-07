package validator

// Validatable Object that is able to validate itself
type Validatable interface {
	// Validate that this object is correctly configured and usable
	Validate() error
}
