package main

import (
	"context"
	"currencyApi/config"
	"currencyApi/currency"
	"currencyApi/redis"
	"currencyApi/server"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	goredis "github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func main() {
	cfg := cfg()

	redisClient := redisClient(ctx, cfg.Redis)
	defer redisClient.Close()

	currencyRepository := redis.NewCurrencyRepository(ctx, redisClient)
	currencyService := currency.NewCurrencyService(currencyRepository)

	srv := server.New(cfg.Server.Address, currencyService)
	srv.Run()

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)

	<- shutdown
	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Printf("server: Error while attempt shutdown: %s", err)
	}
}

func cfg() *config.Config {
	c, err := config.New()
	if err != nil {
		panic(err)
	}

	return c
}

func redisClient(ctx context.Context, cfg config.Redis) *goredis.Client {
	c := goredis.NewClient(&goredis.Options{
		Addr: cfg.Address,
		Password: cfg.Password,
		DB: cfg.Database,
	})

	return c
}