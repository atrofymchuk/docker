package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type config struct {
	PostgresUri   string `env:"POSTGRES_URI" envDefault:"postgres://root:pass@127.0.0.1/postgres"`
	ListenAddress string `env:"LISTEN_ADDRESS" envDefault:":7000"`
}

var (
	db          *sql.DB
	errorsCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gocalc_errors_count",
			Help: "Gocalc Errors Count Per Type",
		},
		[]string{"type"},
	)

	requestsCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "gocalc_requests_count",
			Help: "Gocalc Requests Count",
		})
)

func main() {
	var err error

	// Initing prometheus
	prometheus.MustRegister(errorsCount)
	prometheus.MustRegister(requestsCount)

	// Getting env
	cfg := config{}
	if err = env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	// Connecting to database
	db, err = sql.Open("postgres", cfg.PostgresUri)
	if err != nil {
		log.Fatalf("Can't connect to postgresql: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Can't ping database: %v", err)
	}

	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(cfg.ListenAddress, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestsCount.Inc()

	keys, ok := r.URL.Query()["q"]
	if !ok || len(keys[0]) < 1 {
		errorsCount.WithLabelValues("missing").Inc()
		log.Println("Url Param 'q' is missing")
		http.Error(w, "Bad Request", 400)
		return
	}
	q := keys[0]
	log.Println("Got query: ", q)

	var result string
	sqlStatement := fmt.Sprintf("SELECT (%s)::numeric", q)
	row := db.QueryRow(sqlStatement)
	err := row.Scan(&result)

	if err != nil {
		log.Println("Error from db: %s", err)
		errorsCount.WithLabelValues("db").Inc()
		http.Error(w, "Internal Server Error", 500)
		return
	}

	fmt.Fprintf(w, "query %s; result %s", q, result)
}
