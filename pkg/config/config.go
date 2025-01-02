package config

import (
    "os"
    "log"
)

type Config struct {
    Port     string
    Database string
}

func LoadConfig() (*Config, error) {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default port
    }

    database := os.Getenv("DATABASE_URL")
    if database == "" {
        log.Fatal("DATABASE_URL must be set")
    }

    return &Config{
        Port:     port,
        Database: database,
    }, nil
}