package config

import (
    "fmt"
    "log"
    "os"
    "sync"

    "github.com/joho/godotenv"
)

type Config struct {
    AppPort      string
    MySQLHost    string
    MySQLPort    string
    MySQLUser    string
    MySQLPass    string
    MySQLDB      string
    MySQLParams  string
}

var (
    cfg  Config
    once sync.Once
)

func Load() (Config, error) {
    var loadErr error
    once.Do(func() {
        _ = godotenv.Load()
        cfg = Config{
            AppPort:     getEnv("APP_PORT", "8080"),
            MySQLHost:   getEnv("MYSQL_HOST", "127.0.0.1"),
            MySQLPort:   getEnv("MYSQL_PORT", "3306"),
            MySQLUser:   getEnv("MYSQL_USER", "root"),
            MySQLPass:   getEnv("MYSQL_PASSWORD", ""),
            MySQLDB:     getEnv("MYSQL_DB", "forum"),
            MySQLParams: getEnv("MYSQL_PARAMS", "parseTime=true&loc=Local"),
        }
        if cfg.MySQLUser == "" || cfg.MySQLDB == "" {
            loadErr = fmt.Errorf("missing required MySQL env vars")
        }
    })
    return cfg, loadErr
}

func MustLoad() Config {
    conf, err := Load()
    if err != nil {
        log.Fatalf("config error: %v", err)
    }
    return conf
}

func getEnv(key, fallback string) string {
    if val, ok := os.LookupEnv(key); ok && val != "" {
        return val
    }
    return fallback
}

