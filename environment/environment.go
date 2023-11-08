package environment

import "os"

// GetEnvironment gets the name of the current environment
func GetEnviroment() string {
	return os.Getenv("ENV")
}
