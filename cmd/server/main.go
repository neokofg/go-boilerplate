package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"go-boilerplate/prelude/database"
	"go-boilerplate/prelude/env"
	"go-boilerplate/prelude/logs"
	"go-boilerplate/prelude/server"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Количество HTTP запросов",
		},
		[]string{"method", "endpoint", "status"},
	)
)

func main() {
	env.InitDotenv()
	sugar := logs.InitZap()
	client := database.InitDb(sugar)
	defer client.Close()
	server.InitGin(sugar, client)
}
