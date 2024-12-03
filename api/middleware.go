package api

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/nrmnqdds/vtech-qms-be/internal"
)

func CheckAuthHeader(key string, _ echo.Context) (bool, error) {
	if key == "" {
		return false, nil
	}

	if key == os.Getenv("API_KEY") {
		return true, nil
	}

	_, err := internal.DecodePasetoToken(key)
  if err != nil {
    return false, nil
  }

	return err == nil, nil
}
