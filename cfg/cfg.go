package cfg

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

var (
	SERVICE_USER_HOST   string = getEnv("SERVICE_USER_HOST", "0.0.0.0")
	SERVICE_USER_PORT   string = getEnv("SERVICE_USER_PORT", "50051")
	SERVICE_USER_SECRET string = mustGetEnv("SERVICE_USER_SECRET")
	MYSQL_USER          string = mustGetEnv("MYSQL_USER")
	MYSQL_PASS          string = mustGetEnv("MYSQL_PASS")
	MYSQL_NAME          string = mustGetEnv("MYSQL_NAME")
	MYSQL_HOST          string = mustGetEnv("MYSQL_HOST")
	MYSQL_PORT          string = mustGetEnv("MYSQL_PORT")
	DSN                 string = fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s",
		MYSQL_USER,
		MYSQL_PASS,
		MYSQL_HOST,
		MYSQL_PORT,
		MYSQL_NAME,
	)
)
