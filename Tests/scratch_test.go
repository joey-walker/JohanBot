package main

import (
	"fmt"
	"testing"
	"os"
)

func TestScratch(*testing.T) {
	fmt.Println("",os.Getenv("JOHAN_TOKEN"))
}
