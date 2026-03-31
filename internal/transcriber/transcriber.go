package transcriber

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type Transcriber struct {
	client        openai.Client
	filepath      string
	ctx           context.Context
	transcription string
}

func NewTranscriber(ctx context.Context) *Transcriber {
	baseUrl := "https://bothub.chat/api/v2/openai/v1"
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env file")
	}
	token := os.Getenv("BOTHUB_API_TOKEN")
	client := openai.NewClient(option.WithAPIKey(token), option.WithBaseURL(baseUrl))
	filepath := "./tmp/mp3/tmp.mp3"

	return &Transcriber{client: client, filepath: filepath}
}

func (t *Transcriber) GetTranscription() string {
	file, err := os.Open(t.filepath)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	defer file.Close()

	response, err := t.client.Audio.Transcriptions.New(t.ctx, openai.AudioTranscriptionNewParams{
		Model: openai.AudioModelWhisper1,
		File:  file,
	})
	if err != nil {
		log.Fatal("Failed to create request to Bothub")
	}

	t.transcription = response.Text

	return t.transcription
}

func (t *Transcriber) GetSummary() string {
	if t.transcription == "" {
		t.transcription = t.GetTranscription()
	}

	request := fmt.Sprintf("Сделай краткое содержание данного текста:\n%s", t.transcription)
	options := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT5_4,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(request),
		},
	}

	response, err := t.client.Chat.Completions.New(t.ctx, options)
	if err != nil {
		log.Fatal("Failed to create request to Bothub")
	}

	summary := response.Choices[0].Message.Content

	return summary
}

func (t *Transcriber) GetPoints() string {
	if t.transcription == "" {
		t.transcription = t.GetTranscription()
	}

	request := fmt.Sprintf("Проанализируй текст, выдели основные смыслы, "+
		"на основе которых можно создать тригеррный контент. "+
		"Без примеров:\n%s", t.transcription)
	options := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT5_4,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(request),
		},
	}

	response, err := t.client.Chat.Completions.New(t.ctx, options)
	if err != nil {
		log.Fatal("Failed to create request to Bothub")
	}

	points := response.Choices[0].Message.Content

	return points
}
