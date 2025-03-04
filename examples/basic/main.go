package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// Set explicit the model to use, an alternative to setting
	// the OPENAI_MODEL environment variable in the .env file
	option := openai.WithModel("gpt-4o-mini")

	// Create a new langchain model
	llm, err := openai.New(option)
	if err != nil {
		log.Fatalf("error creating langchain model: %v", err)
	}

	// Creates a new background context, which is used for managing timeouts
	// and cancellations in API requests..
	ctx := context.Background()

	// Requests a completion for the given prompt.
	completion, err := llm.Call(ctx, "Write a 'Hello, World!' program in Go.")
	if err != nil {
		log.Fatalf("error calling langchain model: %v", err)
	}

	fmt.Println(completion)
}
