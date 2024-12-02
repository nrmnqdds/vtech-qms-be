package config

import (
	"fmt"
	"os"
	"strconv"
)

const defaultPort = 1323

func GetServerString() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(defaultPort)
	}
	return fmt.Sprintf(":%v", port)
}
