package common

import (
	"errors"
	"testing"
)

type testStruct struct {
	x int
	y int
}

func TestLogInfo(*testing.T) {
	LogInfoAPI(DISCORD, "TEST", "This is sample info")
}

func TestLogErrorAPIFatal(*testing.T) {
	err := errors.New("Test")
	LogErrorAPIFatal(SLACK, "logging_test.go", err)
}

func TestLogDebug(t *testing.T) {
	fakestruct := testStruct{5, 5}
	LogDebugAPI(TEST, "logging_test.go", "", fakestruct)
}
