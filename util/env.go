package util

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func EnvOr(env, defValue string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		return defValue
	}

	return value
}
