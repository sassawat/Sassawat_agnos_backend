package main

import (
	"bytes"
	"database/sql"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		reqBody := string(bodyBytes)

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		respBody := &bytes.Buffer{}
		writer := &bodyWriter{body: respBody, ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next()

		logEntry := LogEntry{
			Method:       c.Request.Method,
			URL:          c.Request.URL.String(),
			RequestBody:  reqBody,
			ResponseBody: respBody.String(),
			StatusCode:   c.Writer.Status(),
			CreatedAt:    time.Now(),
		}

		if err := logRequest(db, logEntry); err != nil {
			log.Println("Failed to log request:", err)
		}
	}
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func logRequest(db *sql.DB, entry LogEntry) error {
	_, err := db.Exec(
		`INSERT INTO request_logs (method, url, request_body, response_body, status_code, created_at) 
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		entry.Method, entry.URL, entry.RequestBody, entry.ResponseBody, entry.StatusCode, entry.CreatedAt,
	)
	return err
}
