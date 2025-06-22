package main

import (
	"github.com/gin-gonic/gin"
	"github.com/john-henry-alfeche/inventory-app-api/config"
	"github.com/john-henry-alfeche/inventory-app-api/routes"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	log.Printf("🛠️ Gin mode: %s\n", gin.Mode())

	router := routes.SetupRoutes()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("✅ Server running on http://localhost:%s\n", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
