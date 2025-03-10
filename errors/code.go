package errors

// ErrorCodeMap is a map of error codes to error messages
var ErrorCodeMap = map[ErrorType]string{}

// GetCode get error code
func GetCode(status int) string {
	msg, ok := ErrorCodeMap[ErrorType(status)]
	if ok {
		return msg
	}
	return ErrorCodeMap[Error]
}
