package common

import (
	"errors"
	"testing"
)

func TestLogInfo(*testing.T) {
	LogApiInfo(DISCORD, "TEST", "This is sample info")
}

func TestLogErrorAPIFatal(*testing.T) {
	err := errors.New("Test")
	LogErrorAPIFatal(SLACK, "logging_test.go", err)
}
