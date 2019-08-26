package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nEdAy/Shepherd/model"
	"github.com/nEdAy/Shepherd/pkg/config"
	"github.com/nEdAy/Shepherd/pkg/logger"
	"github.com/nEdAy/Shepherd/pkg/redis"
	"github.com/nEdAy/Shepherd/router"
	"github.com/rs/zerolog/log"
)

func init() {
	logger.Setup()
	config.Setup()
	model.Setup()
	redis.Setup()
	router.Setup()
}

func main() {
	initGin()
	startGinServer()
}

func initGin() {
	gin.SetMode(config.App.RunMode)
	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()
	// Logging to a file.
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
}

func startGinServer() {
	// Listen and Server in 127.0.0.1:8000
	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler:        router.Router,
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if config.Server.Protocol == "http" {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal().Msgf("http listen: %s\n", err)
			}
		} else {
			if err := server.ListenAndServeTLS(config.Path.CertFilePath, config.Path.KeyFilePath); err != nil && err != http.ErrServerClosed {
				log.Fatal().Msgf("https listen: %s\n", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server Shutdown:", err)
	}
	log.Info().Msg("Server exiting")
}
