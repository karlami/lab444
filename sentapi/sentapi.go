package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func getContent(pathFile string) (content string) {

	fileExtension := filepath.Ext(pathFile)

	if fileExtension == "pdf" {
		content = readPdf(pathFile)
	} else if fileExtension == "txt" {
		content = readtxt(pathFile)
	} else {
		content = readDocx(pathFile)
	}
	fmt.Println("File extension ", fileExtension)

	return content
}

func sentapi(pathFile string) (sentimentResult string, scoreResult int) {

	content := getContent(pathFile)

	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: content,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}

	fmt.Printf("Text: %v\n", string(content))
	scoreResult = int(sentiment.DocumentSentiment.Score)
	if sentiment.DocumentSentiment.Score >= 0 {
		sentimentResult = "Positive"
		fmt.Println("Sentiment: positive")
	} else {
		sentimentResult = "Positive"
		fmt.Println("Sentiment: negative")
	}
	return
}
