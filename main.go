package keepCalm

import "github.com/cygnusborjigin/movingOn"

type MainError struct {
	errorMessage map[string]interface{}
}

func NewError() (MainError, bool) {
	newErrorMessage := make(map[string]interface{})
	return MainError{
		errorMessage: newErrorMessage,
	}, true
}

func (e MainError) IsEmpty() bool {
	if len(e.errorMessage) == 0 {
		return true
	}
	return false
}

func (e MainError) GetErrorMessage() map[string]interface{} {
	return e.errorMessage
}

func (e MainError) isAtomic() bool {
	for _, value := range e.errorMessage {
		_, isString := value.(string)
		if !isString {
			return false
		}
	}
	return true
}

func (e MainError) PushMessage(newError SimpleError) *string {
	errorKey, errorValue := newError.ShowError()
	// push message into MainError
	if MovingOn.MapStringInterfaceContainKey(e.errorMessage, errorKey) {
		_, isStringSlice := e.errorMessage[errorKey].([]string)
		if !isStringSlice {
			errorMessage := "Field is not atomic"
			return &errorMessage
		}
		e.errorMessage[errorKey] = append(e.errorMessage[errorKey].([]string), errorValue)
	} else {
		newSlice := []string{errorValue}
		e.errorMessage[errorKey] = newSlice
	}
	return nil
}

func (e MainError) PushError(newError StringErrorTuple) *string {
	errorKey, errorMessage := newError.ShowError()
	if MovingOn.MapStringInterfaceContainKey(e.errorMessage, errorKey) {
		errMessage := "duplicate key"
		return &errMessage
	}
	e.errorMessage[errorKey] = errorMessage
	return nil
}
