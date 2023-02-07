package shared

import (
	"fmt"
	"os"
	"strings"
)

// GetEnv method that reads the environment variables needed in the project.
//
// Note: if an environment variable is not found, a panic will occur.
func GetEnv(key string) string {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		panic(fmt.Sprintf("missing environment variable '%s'", key))
	}
	return value
}
