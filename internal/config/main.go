package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr *string
	Dsn  *string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func LoadConfig() *Config {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbString := fmt.Sprintf(`postgresql://%s:%s/%s`, dbHost, dbPort, dbName)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	addrString := fmt.Sprintf(`%s:%s`, host, port)

	addr := flag.String("addr", addrString, "HTTP network address")
	dsn := flag.String("dsn", dbString, "PostgreSQL data source name")
	flag.Parse()

	return &Config{
		Addr: addr,
		Dsn:  dsn,
	}
}
