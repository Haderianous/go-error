package errors

func Forbidden(err ...error) ErrorModel {
	return New(err...).WithType(TypeForbidden)
}

func NotFound(err ...error) ErrorModel {
	return New(err...).WithType(TypeNotFound)
}

func UnAuthorized(err ...error) ErrorModel {
	return New(err...).WithType(TypeUnAuthorized)
}

func ServiceUnavailable(err ...error) ErrorModel {
	return New(err...).WithType(TypeUnAvailable)
}

func UnProcessable(err ...error) ErrorModel {
	return New(err...).WithType(TypeUnProcessable)
}

func Duplicated(err ...error) ErrorModel {
	return New(err...).WithType(TypeDuplicate)
}

func BadRequest(err ...error) ErrorModel {
	return New(err...).WithType(TypeBadRequest)
}

func Accepted(err ...error) ErrorModel {
	return New(err...).WithType(TypeAccepted)
}
