package util

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
)

func GenerateContentTextOnly(query string) ([]genai.Part, error) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GEMINI_API))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(query))
	if err != nil {
		log.Fatal(err)
	}

	println(query)
	result := printResponse(resp)
	return result, nil
}
func printResponse(resp *genai.GenerateContentResponse) []genai.Part {
	var result []genai.Part
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
				result = append(result, part)
			}
		}
	}

	return result
}

func GenerateContentImagePrompt(imageData []byte, fileType string, query string) ([]genai.Part, error) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GEMINI_API))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx,
		genai.Text(query),
		genai.ImageData(fileType, imageData))
	if err != nil {
		return nil, err
	}

	return printResponse(resp), nil
}
