package errors

var DefaultServiceUnAvaialable = ServiceUnavailable().
	WithMessageId("ApplicationIsNotResponsive").
	WithErrorId("HttpError").
	WithMessage("The application is not responsive. Please try again").
	WithErrorText("Internal Server Error").SetDefaults(true)

var DefaultNotFound = NotFound().
	WithMessageId("NotFoundError").
	WithErrorId("InvalidData").
	WithMessage("Entity does not exists").
	WithErrorText("Invalid given data").SetDefaults(true)

var DefaultUnProcessable = UnProcessable().
	WithMessageId("UnProcessableError").
	WithErrorId("InvalidData").
	WithMessage("Invalid request information").
	WithErrorText("Invalid given data").SetDefaults(true)

var DefaultForbiddenError = Forbidden().
	WithMessageId("ForbiddenError").
	WithErrorId("InvalidUser").
	WithMessage("Access to this section is denied.").
	WithErrorText("User is not permitted").SetDefaults(true)

var DefaultUnAuthorizedError = UnAuthorized().
	WithMessageId("UnAuthorizedError").
	WithErrorId("InvalidUser").
	WithMessage("You are not authorized.").
	WithErrorText("User is not permitted").SetDefaults(true)

var DefaultDuplicatedError = UnProcessable().
	WithMessageId("DuplicatedError").
	WithErrorId("InvalidData").
	WithMessage("Given data already exists.").
	WithErrorText("Invalid given data").SetDefaults(true)

var DefaultBadRequestError = BadRequest().
	WithMessageId("RequestError").
	WithErrorId("InvalidRequest").
	WithMessage("Request is invalid.").
	WithErrorText("Invalid given data").SetDefaults(true)
