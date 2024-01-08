package models

type KTSErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}

type KTSError struct {
	KTSErrorMessage
	Status  int   `json:"-"`
	Details string `json:"details,omitempty"`
}

func (err *KTSError) WithDetails(details string) *KTSError {
	newErr := new(KTSError)
	*newErr = *err
	newErr.Details = details
	return newErr
}
