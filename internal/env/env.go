package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	res, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return res
}

func GetInt(key string, fallback int) int {
	res, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	convertedRes, err := strconv.Atoi(res)
	if err != nil {
		return fallback
	}
	return convertedRes
}
