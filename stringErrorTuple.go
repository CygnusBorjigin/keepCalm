package keepCalm

type StringErrorTuple struct {
	errorKey     string
	errorMessage map[string]interface{}
}

func CreateStringErrorTuple(key string, value map[string]interface{}) StringErrorTuple {
	return StringErrorTuple{
		errorKey:     key,
		errorMessage: value,
	}
}

func (s StringErrorTuple) GetKey() string {
	return s.errorKey
}

func (s StringErrorTuple) GetMessage() map[string]interface{} {
	return s.errorMessage
}

func (s StringErrorTuple) ShowError() (string, map[string]interface{}) {
	return s.errorKey, s.errorMessage
}
