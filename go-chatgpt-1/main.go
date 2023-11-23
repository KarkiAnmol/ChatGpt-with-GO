package main

import (
	"log"
	"os"

	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}
	client, _ := gpt35.NewClient(apiKey)
	req := &gpt35.Request{
		Model: gpt35.ModelGpt35Turbo,
		Messages: []*gpt35.Message{
			{
				Role:    gpt35.RoleUser,
				Content: "how are you ",
			},
		},
	}

	resp, err := client.GetChat(req)
	if err != nil {
		panic(err)
	}

	println(resp.Choices[0].Message.Content)
}
