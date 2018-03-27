package main

import "os"

func GetDBHost() string {
	fallback := "localhost"
	host := os.Getenv("REDPLANET_DB_HOST")
	if len(host) == 0 {
		return fallback
	}
	return host
}

func GetVerifyToken() string {
	return os.Getenv("TOKEN")
}