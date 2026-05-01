package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Client struct {
		Requests    int64
		LastRequest time.Time
	}
)

var (
	mutex   sync.Mutex
	clients = make(map[string]*Client)
)

func NewRateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mutex.Lock()
		client, ok := clients[ctx.ClientIP()]
		if !ok {
			client = &Client{
				Requests:    0,
				LastRequest: time.Now(),
			}
			clients[ctx.ClientIP()] = client
		}

		if time.Since(client.LastRequest) > time.Second {
			client.Requests = 0
		}

		if client.Requests > 5 {
			client.LastRequest = time.Now()
			mutex.Unlock()

			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		client.Requests++
		client.LastRequest = time.Now()
		mutex.Unlock()

		ctx.Next()
	}
}
