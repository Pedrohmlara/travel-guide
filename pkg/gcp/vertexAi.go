package gcp

import (
	"cloud.google.com/go/vertexai/genai"
	"context"
	"fmt"
	"io"
	"log"
)

func GenerateMultimodalContent(w io.Writer, parts []genai.Part, projectID, location, modelName string, temperature float32) ([]byte, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)
	model.SetTemperature(temperature)

	res, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		return nil, fmt.Errorf("unable to generate contents: %v", err)
	}

	part := res.Candidates[0].Content.Parts[0]
	response, error := part.(genai.Text)
	if !error {
		log.Fatalf("unable to get Text: %v", part)
	}

	return []byte(response), nil
}
