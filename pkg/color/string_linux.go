//go:build linux
// +build linux

package color

import (
	"fmt"
	"math/rand"
	"strconv"
)

var _ = RandomColor()

// RandomColor generates a random color.
func RandomColor() string {
	return fmt.Sprintf("#%s", strconv.FormatInt(int64(rand.Intn(16777216)), 16))
}

// Yellow ...
func Yellow(msg string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", msg)
}

// Red ...
func Red(msg string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", msg)
}

// Redf ...
func Redf(msg string, arg interface{}) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m %+v\n", msg, arg)
}

// Blue ...
func Blue(msg string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", msg)
}

// Green ...
func Green(msg string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", msg)
}

// Greenf ...
func Greenf(msg string, arg interface{}) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m %+v\n", msg, arg)
}
