package restful

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type RestfulCode int

const (
	RequestSuccess RestfulCode = 0

	ServerError        RestfulCode = 1
	RequestTimeout     RestfulCode = 2
	TooManyRequests    RestfulCode = 3
	RequestForbidden   RestfulCode = 4
	ServiceUnavailable RestfulCode = 5
	ParamsError        RestfulCode = 6

	Unauthorized        RestfulCode = 100
	PermissionDenied    RestfulCode = 101
	PermissionNotAllow  RestfulCode = 102
	AccountLocked       RestfulCode = 103
	AccountDisabled     RestfulCode = 104
	InvalidToken        RestfulCode = 105
	TokenExpired        RestfulCode = 106
	RefreshTokenInvalid RestfulCode = 107

	InvalidParams       RestfulCode = 200
	MissingParams       RestfulCode = 201
	InvalidFormat       RestfulCode = 202
	InvalidType         RestfulCode = 203
	RequestBodyTooLarge RestfulCode = 204
	IllegalRequest      RestfulCode = 205

	ResourceNotFound RestfulCode = 300
	ResourceExisted  RestfulCode = 301
	ResourceDisabled RestfulCode = 302
	ResourceLocked   RestfulCode = 303
	DataEmpty        RestfulCode = 304
	DataFormatError  RestfulCode = 305
	FileTooLarge     RestfulCode = 306
	FileUploadFailed RestfulCode = 307

	BusinessError     RestfulCode = 400
	OperateFailed     RestfulCode = 401
	StatusInvalid     RestfulCode = 402
	RepeatSubmit      RestfulCode = 403
	VerifyCodeError   RestfulCode = 404
	VerifyCodeExpired RestfulCode = 405
	PasswordError     RestfulCode = 406
	PasswordExpired   RestfulCode = 407

	DBError           RestfulCode = 500
	RedisError        RestfulCode = 501
	MQError           RestfulCode = 502
	ThirdServiceError RestfulCode = 503
	ConfigError       RestfulCode = 504
)

type Restful[T any] struct {
	Code    RestfulCode `json:"code"`
	Message string      `json:"message"`
	Data    T           `json:"data"`
}

type Pagination[T any] struct {
	Items    []T   `json:"items"`
	Total    int64 `json:"total"`
	Page     uint  `json:"page"`
	PageSize uint  `json:"page_size"`
}

func ParsePaginationRequest(c *gin.Context) (page, pagesize int) {
	pageStr, _ := c.Params.Get("page")
	pagesizeStr, _ := c.Params.Get("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	pagesize, err = strconv.Atoi(pagesizeStr)
	if err != nil {
		pagesize = 10
	}
	return
}
