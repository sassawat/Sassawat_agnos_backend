package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to database.
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

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

	// Start server with port 8080
	r.Run(":8080")
}
