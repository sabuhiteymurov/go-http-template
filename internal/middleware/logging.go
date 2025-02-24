package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func getStatusColor(status int) string {
	switch {
	case status >= 500:
		return Red
	case status >= 400:
		return Yellow
	case status >= 300:
		return Cyan
	case status >= 200:
		return Green
	default:
		return Blue
	}
}

func getClientInfo(userAgent string) (string, string) {
	ua := strings.ToLower(userAgent)

	var os string
	switch {
	case strings.Contains(ua, "windows"):
		os = "Windows"
	case strings.Contains(ua, "ios") || strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		os = "iOS"
	case strings.Contains(ua, "mac"):
		os = "MacOS"
	case strings.Contains(ua, "linux"):
		os = "Linux"
	case strings.Contains(ua, "android"):
		os = "Android"
	}

	var client string
	switch {
	case strings.Contains(ua, "postman"):
		client = "Postman"
	case strings.Contains(ua, "chrome"):
		client = "Chrome"
	case strings.Contains(ua, "firefox"):
		client = "Firefox"
	case strings.Contains(ua, "safari"):
		client = "Safari"
	case strings.Contains(ua, "edge") || strings.Contains(ua, "edg"):
		client = "Edge"
	}

	return os, client
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)
		os, client := getClientInfo(r.UserAgent())

		parts := []string{
			fmt.Sprintf("%s%d%s", getStatusColor(wrapped.statusCode), wrapped.statusCode, Reset),
			fmt.Sprintf("%s%s%s", Cyan, r.Method, Reset),
			fmt.Sprintf("%s%s%s", Blue, r.URL.Path, Reset),
		}

		if os != "" {
			parts = append(parts, fmt.Sprintf("%s%s%s", Yellow, os, Reset))
		}
		if client != "" {
			parts = append(parts, fmt.Sprintf("%s%s%s", Purple, client, Reset))
		}

		parts = append(parts, fmt.Sprintf("%s%.2fms%s", Green, duration.Seconds()*1000, Reset))

		logMessage := strings.Join(parts, " ")
		log.Println(logMessage)
	})
}
