package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIp(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func getRateLimiting(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	client, exists := clients[ip]
	if !exists {
		limiter := rate.NewLimiter(5, 15) // 5 request/sec, bursts 15
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient

		//log.Printf("A clients[%s] - {limiter: %+v, lastSeen: %s}", ip, limiter, newClient.lastSeen)
		return limiter
	}

	//log.Printf("A clients exist[%s] - {limiter: %+v, lastSeen: %s}", ip, client.limiter, client.lastSeen)
	client.lastSeen = time.Now()
	return client.limiter
}

func CleanupClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		//log.Printf("===> clients: %+v", clients)
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIp(ctx)
		//log.Println("ip:", ip) // ::1 <=> 127.0.0.1

		limiter := getRateLimiting(ip)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many requests",
				"message": "Bạn đã gửi quá nhiều request. Hãy thử lại sau",
			})
			return
		}

		ctx.Next()
	}
}
