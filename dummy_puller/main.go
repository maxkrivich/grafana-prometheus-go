package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr = flag.String("address", ":8080", "The address to listen on for HTTP requests.")
)

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "dummy_puller",
			Name:      "counter",
			Help:      "Dumppy puller counter",
		})
	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "dummy_puller",
		Name:      "gauge",
		Help:      "Dummy puller gauge",
	})
	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "dummy_puller",
			Name:      "histogram",
			Help:      "Dummy puller histogram",
		})

	summary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace: "dummy_puller",
			Name:      "summary",
			Help:      "Dummy puller summary",
		})
)

func init() {
	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)
}
func recordMetrics() {
	for {
		counter.Add(rand.Float64() * 5)
		gauge.Add(rand.Float64()*15 - 5)
		histogram.Observe(rand.Float64() * 10)
		summary.Observe(rand.Float64() * 10)

		time.Sleep(time.Second * 3)
	}
}

func main() {
	flag.Parse()

	go recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
