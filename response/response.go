package response

import (
	"github.com/gin-gonic/gin"
)

// ErrorDetail untuk error spesifik dalam request
type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

// ErrorResponse adalah format standar error response
type ErrorResponse struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details,omitempty"`
}

// SendError mengirimkan response error dengan format standar
func SendError(c *gin.Context, httpStatus int, code, message string, details []ErrorDetail) {
	c.JSON(httpStatus, gin.H{
		"error": ErrorResponse{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}
