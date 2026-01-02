package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const defaultMaxSkewSeconds = 60

type HMACConfig struct {
	Secret         string
	MaxSkewSeconds int64
}

func NewHMACAuth(config HMACConfig) gin.HandlerFunc {
	maxSkew := config.MaxSkewSeconds
	if maxSkew <= 0 {
		maxSkew = defaultMaxSkewSeconds
	}

	return func(c *gin.Context) {
		if config.Secret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "HMAC secret not configured"})
			c.Abort()
			return
		}

		// timestamp := c.GetHeader("X-Timestamp")
		signature := c.GetHeader("X-Signature")

		// if timestamp == "" || signature == "" {
		if signature == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing signature headers"})
			c.Abort()
			return
		}

		// ts, err := strconv.ParseInt(timestamp, 10, 64)
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid timestamp"})
		// 	c.Abort()
		// 	return
		// }

		// now := time.Now().Unix()
		// if delta := now - ts; delta > maxSkew || delta < -maxSkew {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "expired signature"})
		// 	c.Abort()
		// 	return
		// }

		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid body"})
			c.Abort()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		path := c.Request.URL.Path
		if rawQuery := c.Request.URL.RawQuery; rawQuery != "" {
			path = path + "?" + rawQuery
		}

		// base := fmt.Sprintf("%s\n%s\n%s\n%s", timestamp, c.Request.Method, path, string(bodyBytes))
		// base := fmt.Sprintf("%s\n%s\n%s\n%s", c.Request.Method, path, string(bodyBytes))
		base := strings.Join([]string{
			c.Request.Method,
			path,
			string(bodyBytes),
		}, "")
		expected := signHMAC(base, config.Secret)

		provided, err := hex.DecodeString(signature)
		if err != nil || !hmac.Equal(provided, expected) {
			fmt.Printf("BASE=%q\n", base)
			fmt.Printf("EXPECTED=%x\n", expected)
			fmt.Printf("SIG_HDR=%q\n", signature)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func signHMAC(base, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(base))
	return mac.Sum(nil)
}
