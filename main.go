package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuanyu90221/ethereum_blockchain_services/configs"
	"github.com/yuanyu90221/ethereum_blockchain_services/ethjsonrpc"
	"github.com/yuanyu90221/ethereum_blockchain_services/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config := configs.GetEnvConfig()
	rpcConnect := ethjsonrpc.GetConnect()
	ginRouter := router.SetupRouter()
	ginRouter.StaticFile("/favicon.ico", "./favicon.svg")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.PORT),
		Handler: ginRouter,
	}

	log.Printf("%s listen on %d\n", config.APP_NAME, config.PORT)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer rpcConnect.Close()
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
