package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func NewLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		with := logger.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("ip", c.ClientIP()),
			slog.String("user_agent", c.Request.UserAgent()),
		)

		c.Next()
		with.Info(
			"request completed",
			slog.Int("status", c.Writer.Status()),
			slog.Int("size", c.Writer.Size()),
		)
	}
}
