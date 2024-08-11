package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestLoggerMiddleware(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	expectedMethod := "POST"
	expectedURL := "/api/strong_password_steps"
	expectedRequestBody := `{"init_password":"aA1"}`
	expectedResponseBody := `{"num_of_steps":3}`
	expectedStatusCode := http.StatusOK

	mock.ExpectExec("INSERT INTO request_logs").
		WithArgs(expectedMethod, expectedURL, expectedRequestBody, expectedResponseBody, expectedStatusCode, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(LoggerMiddleware(db))

	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		var req PasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		steps := calculateSteps(req.Password)
		resp := PasswordResponse{NumOfSteps: steps}
		c.JSON(http.StatusOK, resp)
	})

	req, _ := http.NewRequest("POST", "/api/strong_password_steps", bytes.NewBuffer([]byte(expectedRequestBody)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, expectedStatusCode, w.Code)
	assert.JSONEq(t, expectedResponseBody, w.Body.String())

	assert.NoError(t, mock.ExpectationsWereMet())
}
