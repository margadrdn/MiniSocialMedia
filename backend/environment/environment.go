package environment

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Gets the environment variable with the input name and returns as a string.
//
// This function loads the .env file.
//
// If error occurs, the error is propagated to the caller.
func GetEnvVar(varName string) (val string, err error) {
	if err = godotenv.Load(); err != nil {
		return
	}
	val = os.Getenv("MONGODB_URI")
	if val == "" {
		err = errors.New("Environment variable not set correctly.")
		return
	}
	return
}
