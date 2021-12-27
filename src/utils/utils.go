package utils

import (
	"os"
    "fmt"
)

func MustGet(arg string) string {
    env := os.Getenv(arg)
    if env == "" {
        var message = fmt.Sprintf("env not found: %s", arg)
        // panic(message)
        fmt.Println(message)
    }
    return env
}