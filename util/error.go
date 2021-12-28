package util

import "fmt"

type ErrorType int

const (
	Success                    ErrorType = 200
	Error                      ErrorType = 500
	InvalidParams              ErrorType = 400
	ErrorBadRequest            ErrorType = 421
	ErrorNoPermission          ErrorType = 403
	ErrorNotFound              ErrorType = 404
	ErrorMethodNotAllow        ErrorType = 405
	ErrorInvalidParent         ErrorType = 409
	ErrorAllowDeleteWithChild  ErrorType = 410
	ErrorNotAllowDelete        ErrorType = 411
	ErrorUserDisabled          ErrorType = 412
	ErrorExistMenuName         ErrorType = 413
	ErrorExistRole             ErrorType = 414
	ErrorExistRoleUser         ErrorType = 415
	ErrorNotExistUser          ErrorType = 416
	ErrorLoginFailed           ErrorType = 422
	ErrorInvalidOldPass        ErrorType = 423
	ErrorPasswordRequired      ErrorType = 424
	ErrorTooManyRequest        ErrorType = 429
	ErrorInternalServer        ErrorType = 512
	ErrorAuthCheckTokenFail    ErrorType = 401
	ErrorAuthCheckTokenTimeout ErrorType = 402
	ErrorAuthToken             ErrorType = 408
	ErrorAuth                  ErrorType = 407
	ErrorExistEmail            ErrorType = 430
	ErrorNotExistRole          ErrorType = 431
	ErrorTokenExpired          ErrorType = 461
	ErrorTokenInvalid          ErrorType = 462
	ErrorTokenMalformed        ErrorType = 463

	// System errors
	ErrorMarshal ErrorType = iota + 1000
	ErrorUnmarshal
	ErrorDatabaseGet
	ErrorDatabaseCreate
	ErrorDatabaseUpdate
	ErrorDatabaseDelete

	ErrorInvalidPassword
)

var MsgMap = map[ErrorType]string{
	Success:                    "OK",
	InvalidParams:              "Request parameter error ",
	ErrorAuthCheckTokenFail:    "Token authentication failed",
	ErrorAuthCheckTokenTimeout: "Token time out",
	ErrorAuthToken:             "Token build failed",
	ErrorAuth:                  "Token error",
	Error:                      "Error occurred",
	ErrorInternalServer:        "Server error",
	ErrorExistEmail:            "The Email Address entered already exists in the system",
	ErrorBadRequest:            "Request error",
	ErrorInvalidParent:         "Invalid parent node",
	ErrorAllowDeleteWithChild:  "Contains children, cannot be deleted",
	ErrorNotAllowDelete:        "Resources are not allowed to be deleted",
	ErrorInvalidOldPass:        "Old password is incorrect",
	ErrorNotFound:              "Resource does not exist",
	ErrorPasswordRequired:      "Password is required",
	ErrorExistMenuName:         "Menu name already exists",
	ErrorUserDisabled:          "User is disabled, please contact administrator",
	ErrorNoPermission:          "No access",
	ErrorMethodNotAllow:        "Method is not allowed",
	ErrorTooManyRequest:        "Requests are too frequent",
	ErrorLoginFailed:           "Email or password is invalid",
	ErrorExistRole:             "Role name already exists",
	ErrorNotExistUser:          "Account is invalid",
	ErrorExistRoleUser:         "The role has been given to the user and is not allowed to be deleted",
	ErrorNotExistRole:          "Role user is disabled, please contact administrator",
	ErrorTokenExpired:          "Token is expired",
	ErrorTokenInvalid:          "Token is invalid",
	ErrorTokenMalformed:        "That's not even a token",
}

// GetMsg from status
func GetMsg(status int) string {
	msg, ok := MsgMap[ErrorType(status)]
	if ok {
		return msg
	}
	return MsgMap[Error]
}
func (errType ErrorType) New() error {
	return fmt.Errorf(GetMsg(int(errType)))
}

var (
	ErrTokenExpired   = ErrorTokenExpired.New()
	ErrTokenInvalid   = ErrorTokenInvalid.New()
	ErrTokenMalformed = ErrorTokenMalformed.New()
)
