package errorx

func NewServiceError(code string) *ServiceError {
	return &ServiceError{
		Code: code,
	}
}

// NewDefSE creates a new ServiceError with default values.
func NewDefSE() *ServiceError {
	return NewServiceError(EXCEPTION)
}

// NewFullError creates a new ServiceError with the specified code, message, payload, and stack trace.
func NewFullError(code string, msg string, payload any, stack error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Msg:     msg,
		Payload: payload,
		Stack:   stack,
	}
}

func NewPartialSuccess() *ServiceError {
	return NewServiceError(PartialSuccess)
}

func NewParamsError() *ServiceError {
	return NewServiceError(ParamsError)
}

func NewUnauthorized() *ServiceError {
	return NewServiceError(Unauthorized)
}

func NewBadCredentials() *ServiceError {
	return NewServiceError(BadCredentials)
}

func NewAccessDenied() *ServiceError {
	return NewServiceError(AccessDenied)
}

func NewResourceExist() *ServiceError {
	return NewServiceError(ResourceExist)
}

func NewResourceNotExist() *ServiceError {
	return NewServiceError(ResourceNotExist)
}

func NewConfigError() *ServiceError {
	return NewServiceError(ConfigError)
}
