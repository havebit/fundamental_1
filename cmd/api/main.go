package main

import (
	"context"
	"fmt"
	"hello/todo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

const defaultPort = ":8081"

var port = defaultPort

var builddate string

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("not found .env file")
	}

	if v := os.Getenv("PORT"); v != "" {
		port = v
	}

	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	r := gin.Default()
	r.GET("/ping", pingPongHandler)
	r.GET("/infoz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"builddate": builddate,
		})
	})

	limiter := rate.NewLimiter(5, 5)
	r.GET("/limitz", func(c *gin.Context) {
		if !limiter.Allow() {
			c.Status(http.StatusTooManyRequests)
			return
		}
		c.Status(http.StatusOK)
	})
	r.POST("/todos", todo.NewTaskTodoHandler)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		fmt.Println("Listening and serving HTTP on", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func pingPongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
