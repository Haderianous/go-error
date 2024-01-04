package errors

type Type string

const (
	TypeUnProcessable Type = "UNPROCESSABLE"
	TypeNotFound           = "NOT_FOUND"
	TypeUnAuthorized       = "UNAUTHORIZED"
	TypeForbidden          = "FORBIDDEN"
	TypeUnAvailable        = "UNAVAILABLE"
	TypeDuplicate          = "DUPLICATED"
	TypeBadRequest         = "BAD_REQUEST"
)

func Is(model ErrorModel, t Type) bool {
	return model.Type() == t
}
