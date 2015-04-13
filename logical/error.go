package logical

// ForbiddenError is forbidden error
type ForbiddenError error

// NewForbiddenError will create a ForbiddenError
func NewForbiddenError() (err ForbiddenError) {
	return *new(ForbiddenError)
}
