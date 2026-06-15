package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	ErrCodeBadRequest       ErrorCode = "BAD_REQUEST"
	ErrCodeNotFound         ErrorCode = "NOT_FOUND"
	ErrCodeConflict         ErrorCode = "CONFLICT"
	ErrCodeInternal         ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrCodeUnauthorized     ErrorCode = "UNAUTHORIZED"
	ErrCodeNotAcceptable    ErrorCode = "NOT_ACCEPTABLE"
	ErrCodeMethodNotAllowed ErrorCode = "METHOD_NOT_ALLOWED"
	ErrCodeTooManyRequests  ErrorCode = "TOO_MANY_REQUESTS"
)

type AppError struct {
	Message string
	Code    ErrorCode
	Err     error
}

type APIResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
	Token      any    `json:"token,omitempty"`
}

func (ae *AppError) Error() string {
	return ""
}

func NewError(message string, code ErrorCode) error {

	return &AppError{
		Message: message,
		Code:    code,
	}
}

func WrapError(message string, code ErrorCode, err error) error {

	return &AppError{
		Message: message,
		Code:    code,
		Err:     err,
	}
}

func ResponseError(ctx *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		status := httpStatusFromCode(appErr.Code)
		response := gin.H{
			"error": CapitalizeFirst(appErr.Message),
			"code":  appErr.Code,
		}

		if appErr.Err != nil {
			response["detail"] = CapitalizeFirst(appErr.Err.Error())
		}

		ctx.JSON(status, response)
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
		"code":  ErrCodeInternal,
	})
}

func ResponseSuccess(ctx *gin.Context, status int, message string, data ...any) {
	resp := APIResponse{
		Status:  "success",
		Message: message,
	}

	if len(data) > 0 && data[0] != nil {
		if m, ok := data[0].(map[string]any); ok {
			if p, exist := m["pagination"]; exist {
				resp.Pagination = p
			}

			if t, exist := m["token"]; exist {
				resp.Token = t
			}

			if d, exist := m["data"]; exist {
				resp.Data = d
			} else {
				resp.Data = m
			}
		} else {
			resp.Data = data[0]
		}
	}

	ctx.JSON(status, resp)
}

func ResponseStatusCode(ctx *gin.Context, status int) {
	ctx.Status(status)
}

func ResponseValidation(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusBadRequest, data)
}

func httpStatusFromCode(code ErrorCode) int {
	switch code {
	case ErrCodeBadRequest:
		return http.StatusBadRequest
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeInternal:
		return http.StatusInternalServerError
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeNotAcceptable:
		return http.StatusNotAcceptable
	case ErrCodeTooManyRequests:
		return http.StatusTooManyRequests
	default:
		return http.StatusMethodNotAllowed
	}
}
