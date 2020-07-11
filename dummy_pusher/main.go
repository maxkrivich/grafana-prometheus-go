package main

import (
	"log"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	gatewayURL   = os.Getenv("PUSHGATEWAY_URL")
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "pusher_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func pushMetrics() {
	opsProcessed.Inc()
	if err := push.New(gatewayURL, "pusher").
		Collector(opsProcessed).
		Grouping("ops", "ticks").
		Push(); err != nil {
		log.Printf("Could not push completion time to Pushgateway: %s\n", err.Error())
	}
}

func main() {
	for {
		log.Printf("Pushing metrics at %v\n", time.Now().Unix())
		pushMetrics()
		time.Sleep(5 * time.Second)
	}
}
