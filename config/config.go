package config

import "os"

func GetDBHost() string {
	fallback := "127.0.0.1"
	host := os.Getenv("REDPLANET_DB_HOST")
	if len(host) == 0 {
		return fallback
	}
	return host
}

func GetVerifyToken() string {
	return os.Getenv("TOKEN")
}
