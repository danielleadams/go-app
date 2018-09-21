package main

import (
	"fmt"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	t.Fatalf("Foo!")
}
