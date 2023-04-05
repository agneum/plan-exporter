// Package chatgpt provides a query plan exporter for the ChatGPT.
package chatgpt

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

// Visualizer constants.
const (
	VisualizerType = "chatgpt"
	pgVersion      = "14"
)

// nolint: lll
var cfg = map[string]string{
	"explain": "I have Postgres version %s, and the following is an EXPLAIN plan. Please describe it in a less than 200 words. Explain plan: %s",
	"opt":     "I have Postgres version %s, and the following is an EXPLAIN plan. Provide ideas on its optimization. Ideas should have short explanations (<100 words), and code snippets so I can copy-paste them and try. Explain plan: %s",
}

// ChatGPT defines a query plan exporter for the ChatGPT.
type ChatGPT struct {
	token string
	mode  string
}

// New creates a new ChatGPT exporter.
func New(mode string) *ChatGPT {
	token := os.Getenv("PE_CHATGPT_TOKEN")

	return &ChatGPT{
		token: token,
		mode:  mode,
	}
}

// Export exports command to ChatGPT.
func (c *ChatGPT) Export(plan string) (string, error) {
	client := openai.NewClient(c.token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(cfg[c.mode], pgVersion, plan),
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}

// Target describes ChatGPT mode.
func (c *ChatGPT) Target() string {
	return fmt.Sprintf("ChatGPT. Mode: %s", c.mode)
}
