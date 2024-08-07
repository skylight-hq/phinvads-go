package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gopkg.in/yaml.v2"
)

type application struct {
	logger *slog.Logger
	db     *sql.DB
}

func main() {
	env := getEnvironment()
	config := readConfig(env)

	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	//dsn := flag.String("dsn", "postgresql://localhost:5432/phinvads", "PostgreSQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(config["postgres"])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	app := &application{logger, db}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server", slog.String("addr", *addr))

	err = srv.ListenAndServeTLS("./tls/localhost.pem", "./tls/localhost-key.pem")
	logger.Error(err.Error())
	os.Exit(1)
}

const DEFAULT_ENV = "development"

func getEnvironment() string {
	e := os.Getenv("PHINVADS_GO_ENV")
	if len(e) == 0 {
		return DEFAULT_ENV
	} else {
		return e
	}
}

func readConfig(environment string) map[string]map[string]interface{} {
	config := make(map[string]map[string]interface{})

	yamlFile, err := os.ReadFile(fmt.Sprintf("config/%s.yml", environment))
	if err != nil {
		fmt.Printf("Error reading config file for environment '%s': #%v ", env, err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	return config
}

func openDB(config map[string]interface{}) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config["host"], config["port"], config["user"], config["password"], config["dbname"])
	flag.String("dsn", psqlInfo, "PostgreSQL data source name")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
