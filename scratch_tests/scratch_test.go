package main

import (
	"fmt"
	"os"
	"testing"
)

func TestScratch(*testing.T) {
	fmt.Println("", os.Getenv("JOHAN_TOKEN"))
}
