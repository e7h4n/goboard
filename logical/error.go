package logical

// ForbiddenError is forbidden error
type ForbiddenError error

// NewForbiddenError will create a ForbiddenError
func NewForbiddenError() (err ForbiddenError) {
	return *new(ForbiddenError)
}

// PredictFailedError is predict fail error
type PredictFailedError error

// NewPredictFailedError will create a predict fail error
func NewPredictFailedError() (err PredictFailedError) {
	return *new(PredictFailedError)
}
