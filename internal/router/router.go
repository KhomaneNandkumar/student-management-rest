package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/student-management/initializers"
	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/gologger"
)

const WAIT_SECONDS = 5

func IntializeRouter() {
	c := initializers.GetConfig()
	router := gin.Default()
	InitializeRoutes(router)

	serverAddr := c.SERVERADDR
	if serverAddr == "" {
		serverAddr = "0.0.0.0:8080"
	}

	server := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			gologger.Info().Msgf(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gologger.Info().Msgf("Shutdown Server")
	ctx, cancel := context.WithTimeout(context.Background(), WAIT_SECONDS*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		gologger.Fatal().Msgf(err.Error())
	}
}
