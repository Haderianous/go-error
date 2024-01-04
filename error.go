package errors

import (
	"encoding/json"
	"time"
)

type ErrorModel interface {
	error
	WithError(e error) ErrorModel
	WithErrorId(id string) ErrorModel
	WithMessageId(id string) ErrorModel

	WithMessage(message string) ErrorModel
	Message() string

	WithErrorText(text string) ErrorModel
	ErrorText() string

	WithDetail(message string) ErrorModel
	WithType(errorType Type) ErrorModel

	WithErrors(errors map[string]any) ErrorModel
	Errors() []map[string]any

	WithCode(code int) ErrorModel
	Code() int

	GetError() error
	ErrorId() string
	MessageId() string
	Type() Type
	Detail() string

	Clone() ErrorModel
	IsMsgDefault() bool
	IsIdDefault() bool
	SetDefaults(bool) ErrorModel
	Is(t Type) bool
}

type data struct {
	Total   int              `json:"total"`
	PerPage int              `json:"per_page"`
	Result  []map[string]any `json:"result"`
}

type err struct {
	error
	ErrType       Type   `json:"-"`
	ErrorID       string `json:"-"`
	MessageID     string `json:"-"`
	defaultId     bool
	defaultMsg    bool
	ErrCode       int    `json:"code,omitempty"`
	ErrMessage    string `json:"message,omitempty"`
	ErrText       string `json:"error,omitempty"`
	Data          data   `json:"data,omitempty"`
	Version       string `json:"version"`
	RepresentedAt string `json:"represented_at"`
	detail        string
}

func New(e ...error) (r ErrorModel) {
	r = new(err)
	if e != nil && len(e) > 0 && e[0] != nil {
		r = r.WithError(e[0])
	}
	return r
}

func (error err) Clone() ErrorModel {
	e := error
	return e
}

func Parse(bytes []byte) (ErrorModel, error) {
	var e *err
	if err := json.Unmarshal(bytes, &e); err != nil {
		return nil, err
	}
	return e, nil
}

func (error err) WithError(e error) ErrorModel {
	if e == nil {
		return error
	}
	error.error = e
	return error.WithDetail(e.Error())
}

func (error err) GetError() error {
	return error.error
}

func (error *err) setDefaults() {
	error.Version = "v1"
	error.RepresentedAt = time.Now().Format("2006-01-02 15:04:05")
	if error.ErrorID == "" {
		switch error.ErrType {
		case TypeUnProcessable:
			error.ErrorID = "UnProcessableError"
			error.MessageID = "InvalidData"
			error.defaultId = true
			break
		case TypeNotFound:
			error.ErrorID = "NotFoundError"
			error.MessageID = "InvalidData"
			error.defaultId = true
			break
		case TypeForbidden:
			error.ErrorID = "ForbiddenError"
			error.MessageID = "InvalidUser"
			error.defaultId = true
			break
		case TypeUnAuthorized:
			error.ErrorID = "UnAuthorizedError"
			error.MessageID = "InvalidUser"
			error.defaultId = true
			break
		case TypeUnAvailable:
			error.ErrorID = "UnAvailableError"
			error.MessageID = "HttpError"
			error.defaultId = true
			break
		}
	}
	if error.ErrMessage == "" {
		switch error.ErrType {
		case TypeUnProcessable:
			error.ErrMessage = "Invalid request information"
			error.ErrText = "Invalid given data"
			error.defaultMsg = true
			break
		case TypeNotFound:
			error.ErrMessage = "Entity does not exists"
			error.ErrText = "Invalid given data"
			error.defaultMsg = true
			break
		case TypeForbidden:
			error.ErrMessage = "Access to this section is denied"
			error.ErrText = "User is not permitted"
			error.defaultMsg = true
			break
		case TypeUnAuthorized:
			error.ErrMessage = "You are not authorized"
			error.ErrText = "User is not permitted"
			error.defaultMsg = true
			break
		case TypeUnAvailable:
			error.ErrMessage = "The application is not responsive. Please try again"
			error.ErrText = "Internal server error"
			error.defaultMsg = true
			break
		}
	}
}

func (error err) WithErrorId(id string) ErrorModel {
	error.ErrorID = id
	error.defaultId = false
	return error
}

func (error err) WithMessageId(id string) ErrorModel {
	error.MessageID = id
	error.defaultId = false
	return error
}

func (error err) WithMessage(message string) ErrorModel {
	error.ErrMessage = message
	error.defaultMsg = false
	return error
}

func (error err) Message() string {
	return error.ErrMessage
}

func (error err) WithErrorText(text string) ErrorModel {
	error.ErrText = text
	error.defaultMsg = false
	return error
}

func (error err) ErrorText() string {
	return error.ErrText
}

func (error err) WithCode(code int) ErrorModel {
	error.ErrCode = code
	return error
}

func (error err) Code() int {
	return error.ErrCode
}

func (error err) WithErrors(errors map[string]any) ErrorModel {
	data := data{}
	for k, v := range errors {
		data.Result = append(data.Result, map[string]interface{}{
			"field": k,
			"error": v,
		})
	}
	error.Data = data
	return &error
}

func (error err) Errors() []map[string]any {
	return error.Data.Result
}

func (error err) WithDetail(d string) ErrorModel {
	error.detail = d
	return error
}

func (error err) WithType(t Type) ErrorModel {
	error.ErrType = t
	error.setDefaults()
	return error
}

func (error err) ErrorId() string {
	return error.ErrorID
}

func (error err) MessageId() string {
	return error.MessageID
}

func (error err) Error() string {
	var r string
	if error.ErrorID != "" {
		r += "id: " + error.ErrorID + ", "
	}
	if error.ErrMessage != "" {
		r += "message: " + error.ErrMessage + ", "
	}
	if error.error != nil {
		r += "err: " + error.ErrText + ", "
	}
	if error.detail != "" {
		r += "detail: " + error.detail + ", "
	}
	r += "represented_at: " + error.RepresentedAt
	return r
}

func (error err) Type() Type {
	return error.ErrType
}

func (error err) Detail() string {
	return error.detail
}

func (error err) IsMsgDefault() bool {
	return error.defaultMsg
}

func (error err) IsIdDefault() bool {
	return error.defaultId
}

func (error err) SetDefaults(bool) ErrorModel {
	error.defaultId = true
	error.defaultMsg = true
	return error
}

func (error err) Is(t Type) bool {
	return error.ErrType == t
}
