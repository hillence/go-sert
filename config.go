package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	BotToken   string
	AdminID    int64
	CardNumber string
	CardBank   string
	CardName   string
	AppPort    string
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Printf("env file not loaded: %v", err)
	}

	BotToken = mustEnv("BOT_TOKEN")
	AdminID = mustInt64("ADMIN_ID")
	CardNumber = mustEnv("CARD_NUMBER")
	CardBank = getEnv("CARD_BANK", "Bank")
	CardName = getEnv("CARD_NAME", "Card Holder")
	AppPort = getEnv("PORT", "8080")
}

func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return val
}

func mustInt64(key string) int64 {
	val := mustEnv(key)
	parsed, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Fatalf("invalid value for %s: %v", key, err)
	}
	return parsed
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
