package main

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	Method       string
	URL          string
	RequestBody  string
	ResponseBody string
	StatusCode   int
	CreatedAt    time.Time
}

type PasswordRequest struct {
	Password string `json:"init_password"`
}

type PasswordResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
