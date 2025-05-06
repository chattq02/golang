package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(c *gin.Context) (string, bool) {
	// Authorization: Bearer token
	authHeader := c.GetHeader("Authorization")

	if strings.HasPrefix(authHeader, "Bearer") { // Kiểm tra xem chuỗi authHeader có bắt đầu bằng "Bearer" hay không.
		return strings.TrimPrefix(authHeader, "Bearer"), true //Loại bỏ tiền tố "Bearer" khỏi chuỗi authHeader, chỉ giữ lại phần token phía sau.
	}

	return "", false // No token found in Authorization header
}