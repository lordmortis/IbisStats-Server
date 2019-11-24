package auth

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func GinMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
		redisClient := ctx.MustGet("redisConnection").(*redis.Client)
		redisPrefix := ctx.MustGet("redisPrefix").(string)

		header := ctx.GetHeader("Authorization")
		if len(header) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 && headerParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		secretString, err := ctx.Cookie("sessionSecret")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		secret, err := base64.StdEncoding.DecodeString(secretString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		sessionID := headerParts[1]

		data, err := redisClient.Get(redisPrefix + "session:" + sessionID).Result()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		redisSession := redisSessionData{}
		err = json.Unmarshal([]byte(data), &redisSession)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "internal server error"})
			return
		}

		if subtle.ConstantTimeCompare(redisSession.Secret, secret) != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		expiryTime := time.Now().Add(expiryDuration)
		redisClient.Set(redisPrefix + "session:" + sessionID, data, expiryDuration)

		maxAge := int(math.Round(expiryDuration.Seconds()))
		domain := strings.Split(ctx.Request.Host, ":")[0]
		ctx.SetCookie("sessionSecret", secretString, maxAge, "/", domain, false, true)

		ctx.Set("AuthModelType", redisSession.Model)
		ctx.Set("AuthModelID", redisSession.ModelID)
		ctx.Set("SessionID", sessionID)
		ctx.Set("ExpiryTime", expiryTime)

		_ = dbCon
	}
}