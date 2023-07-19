package main

import (
	"authserver/config"
	"authserver/internal/server"
	"authserver/internal/server/http/handler"
	"authserver/pkg/logging"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s", buildVersion)
	fmt.Println()
	fmt.Printf("Build date: %s", buildDate)
	fmt.Println()
	fmt.Printf("Build commit: %s", buildCommit)
	fmt.Println()

	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	ctx, cancel := context.WithCancel(context.Background())
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		<-quit
		cancel()
		return nil
	})

	logger.Println("config initializing")
	cfg := config.GetConfig()

	db, err := newDB(ctx, cfg.RedisUri)
	if err != nil {
		logger.Error(err)
		return
	}

	s := service.NewService(db)
	r := handler.NewHandler(cfg, *s, logger, cfg.Secret)
	srv := server.NewServer(logger, cfg.Address, r.Init())

	g.Go(func() error {
		return srv.Start()
	})

	g.Go(func() error {
		<-gCtx.Done()
		return srv.GetServer().Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
	defer cancel()
}

func newDB(ctx context.Context, addr string) (*redis.Client, *redis.Client) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err := RedisClient.Set(ctx, "test", "How to Refresh Access Tokens the Right Way in Golang", 0).Err()
	if err != nil {
		panic(err)
	}

	return RedisClient, nil
}
