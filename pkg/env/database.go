package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	StringConnection string
)

func LoadEnvs() {
	godotenv.Load()
	StringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
