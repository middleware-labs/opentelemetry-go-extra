package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	runtimemetrics "go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	ctx := context.Background()
	configureOpentelemetry()

	meter := otel.GetMeterProvider().Meter("example")
	counter, err := meter.Int64Counter(
		"test.my_counter",
		metric.WithDescription("Just a test counter"),
	)
	if err != nil {
		panic(err)
	}

	for {
		n := rand.Intn(1000)
		time.Sleep(time.Duration(n) * time.Millisecond)

		counter.Add(ctx, 1)
	}
}

func configureOpentelemetry() {
	if err := runtimemetrics.Start(); err != nil {
		panic(err)
	}
	_ = configureMetrics()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("listenening on http://localhost:8088/metrics")

	go func() {
		_ = http.ListenAndServe(":8088", nil)
	}()
}

func configureMetrics() *prometheus.Exporter {
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}
	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))

	otel.SetMeterProvider(provider)

	return exporter
}
