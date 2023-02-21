package pac

import "os"

// Env returns the value of the environment variable named by the key,
// panic fallback if not found
func Env(key, errorMessage string) string {
	env_var := os.Getenv(key)

	if env_var == "" {
		panic(errorMessage)
	}

	return env_var
}
