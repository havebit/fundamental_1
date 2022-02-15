package main

import (
	"context"
	"fmt"
	"hello/todo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := "host=localhost user=postgres password=mysecretpassword dbname=myapp port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()

	r.POST("/logins", loginHandler)

	protectRounter := r.Group("")

	todoHandler := todo.NewHandler(db)
	protectRounter.POST("/todos", todoHandler.NewTaskTodoHandler)
	protectRounter.GET("/todos", todoHandler.ListTaskTodoHandler)

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

type Credential struct {
	Account  string
	Password string
}

var mySigningKey = []byte("AllYourBase")

func authenMiddleware(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	//"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJwYWxsYXQiLCJleHAiOjE2NDQ5MTAzNTV9.THHGbJhNLMq522iBu72WZUV2vd0obkkf3hZtkUA3SdE"
	tokenString := strings.TrimPrefix(bearer, "Bearer ")
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func loginHandler(c *gin.Context) {
	var cred Credential
	if err := c.Bind(&cred); err != nil {
		return
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		Audience:  cred.Account,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}

func pingPongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
