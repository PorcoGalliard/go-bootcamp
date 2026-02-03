package response

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, status int, msg string, data interface{}) {
	resp := ApiResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}
	c.JSON(status, resp)
}

func SendError(c *gin.Context, status int, msg string) {
	resp := ApiResponse{
		Success: false,
		Message: msg,
		Data:    nil,
	}
	c.JSON(status, resp)
}
