package config

import (
	"fmt"
	"os"
)

type Config struct {
	Server Server
	DB     DataBase
}

type Server struct {
	Address string
	Port    string
}

type DataBase struct {
	DBUser     string
	DBPassword string
	DBAddress  string
	DBPort     string
	DBName     string
}

func Load() Config {
	return Config{
		Server{
			Address: getEnv("Adress", ""),
			Port:    getEnv("Port", "8080"),
		},
		DataBase{
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASSWORD", "latte"),
			DBAddress:  getEnv("DB_ADDRESS", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBName:     getEnv("DB_NAME", "frappuccino"),
		},
	}
}

func (d *DataBase) MakeConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.DBAddress, d.DBPort, d.DBUser, d.DBPassword, d.DBName,
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
