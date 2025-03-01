package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	llm, err := openai.New()
	if err != nil {
		log.Fatalf("error creating langchain model: %v", err)
	}

	ctx := context.Background()

	completion, err := llm.Call(ctx, "Write a 'Hello, World!' program in Go.")
	if err != nil {
		log.Fatalf("error calling langchain model: %v", err)
	}

	fmt.Println(completion)
}
