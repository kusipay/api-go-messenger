package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mefellows/vesper"
)

func Handler(ctx context.Context) (string, error) {
	fmt.Println("Hello", "world")

	something := map[string]string{
		"hello": "world",
		"foo":   "bar",
		"baz":   "qux",
	}

	bts, err := json.MarshalIndent(something, "", "  ")
	if err == nil {
		fmt.Println("something |", string(bts))
	}

	return "Hello, World!", errors.New("Pemrimto")
}

func main() {
	v := vesper.New(Handler)

	v.Start()
}
