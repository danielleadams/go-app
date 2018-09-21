package main

import (
	"os"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {

	t.Fatalf(strings.Join(os.Environ(), "\n"))
}
