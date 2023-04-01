package keepCalm

type SimpleError struct {
	errorKey     string
	errorMessage string
}

func CreateSimpleError(key string, value string) SimpleError {
	return SimpleError{
		errorKey:     key,
		errorMessage: value,
	}
}

func (s SimpleError) GetKey() string {
	return s.errorKey
}

func (s SimpleError) GetMessage() string {
	return s.errorMessage
}

func (s SimpleError) ShowError() (string, string) {
	return s.errorKey, s.errorMessage
}
