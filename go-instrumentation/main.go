package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(zerolog.NewConsoleWriter())

	log.Info().Msg("Starting Server!")

	r := resource.NewWithAttributes(
		// schema URL for semantic convenetion, this is will contain URL for semantic schema
		semconv.SchemaURL,
		// service name which is being instrumented
		semconv.ServiceName("go-instrumentation"),
		// service version
		semconv.ServiceVersion("0.0.1"),
	)

	// create console exporter
	// OpenTelemetry allows to create your own custom exporter using OTLP
	// Else we can use default exporter https://github.com/open-telemetry/opentelemetry-go/tree/v1.15.1/exporters
	exp, err := stdoutmetric.New()
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to Start: %v", err))
	}

	// reader is a wrapper for exporter
	// it allowes to define several properties to control behaviour of export
	// such as duration at which metric should be exported or timeout for export
	reader := sdkmetric.NewPeriodicReader(exp, sdkmetric.WithInterval(time.Duration(10000*time.Millisecond)))

	meterProvider := sdkmetric.NewMeterProvider(
		// resource created in early step
		sdkmetric.WithResource(r),
		// reader for the metrics. It can be OTEL collector
		sdkmetric.WithReader(reader),
	)

	defer func() {
		// meter provider should be shutdown before exiting application
		// it will flush all the pending telemetry data
		err = meterProvider.Shutdown(context.Background())
		if err != nil {
			log.Error().Msg(fmt.Sprintf("Failed to shutown: %v", err))
		}
	}()

	// Create Metric to group http metric
	// instrument to represent http metrics will be created using this meter
	httpMeter := meterProvider.Meter("http")
	// create duration instrument
	// it will represent duration for http request
	durationIns, err := httpMeter.Int64Histogram(
		"requests.duration",
	)

	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to Start: %v", err))
	}

	// create server mux
	mux := http.NewServeMux()
	// simple hello world api call
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	})

	// simple API which calls another third party api and return the response from it
	userHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		res, err := http.Get("https://random-data-api.com/api/users/random_user")
		if err != nil {
			log.Err(err).Msg("user request failed!")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		users, err := io.ReadAll(res.Body)
		if err != nil {
			log.Err(err).Msg("user request failed!")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(users)
	})

	// decorate function handlers to provide additional functionality
	mux.Handle("/hello", httpInstrumentationMiddleware(helloHandler, durationIns))
	mux.Handle("/user", httpInstrumentationMiddleware(userHandler, durationIns))

	// serve the http request
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to Start: %v", err))
	}
}

// It's a middleware function
// we will wrap the func handlers using this function
func httpInstrumentationMiddleware(next http.Handler, durationIns metric.Int64Histogram) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// start time for the API call
		startTime := time.Now()
		// call the server API request
		next.ServeHTTP(w, r)
		// finish time for the API call
		end := time.Now()
		// difference between the start and end
		diff := end.Sub(startTime).Milliseconds()
		// record the metric using instrument with attributes such as  path
		// attributes provide more info about any metric
		// attributes can be used for filter and grouping of metrics to make better sense out of metrics
		durationIns.Record(r.Context(), diff, metric.WithAttributes(attribute.String("path", r.URL.Path)))
	})
}
