package main

import (
	"GophKeeper/keepserver/config"
	"GophKeeper/keepserver/internal/server"
	"GophKeeper/keepserver/internal/server/http/handler"
	"GophKeeper/keepserver/internal/service"
	"GophKeeper/keepserver/pkg/logging"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
	"os"
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

	logger.Println("config initializing")
	cfg := config.GetConfig()

	ctx, cancel := context.WithCancel(context.Background())
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		<-quit
		cancel()
		return nil
	})

	db, err := newDB(ctx, cfg.DatabaseDSN)
	if err != nil {
		logger.Error(err)
		return
	}

	s := service.NewService(db)
	r := handler.NewHandler(*s, logger)
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

func newDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
