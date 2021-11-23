package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	userPublicDelivery "golang-lab/delivery/http/public"
	"golang-lab/repository"
	"golang-lab/usecase"

	"github.com/labstack/echo/v4"
)

var jsonFile *os.File

func init() {
	var err error
	jsonFile, err = os.Open("save.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened save.json")
}

func main() {
	e := echo.New()
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"succuss": true})
	})

	uc := usecase.NewUsecase(
		repository.NewRepository(jsonFile),
	)

	v1 := e.Group("/v1")

	// log.Println("start public ... ")
	userPublicDelivery.NewHTTP(v1, uc)

	// v1.Use(middlewares.AuthMiddleWare())
	// v1.Use(middlewares.Validation(postgresDB))

	// log.Println("start private ... ")
	// userPrivateDelivery.NewHTTP(v1, uc)

	serveGracefulShutdown(e)
}

func serveGracefulShutdown(e *echo.Echo) {
	go func() {
		if err := e.Start(":8080"); err != nil {
			fmt.Println("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gracefulShutdownTimeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
