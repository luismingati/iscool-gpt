package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const systemPrompt = `Você é um assistente educacional especializado em programação.
Responda APENAS perguntas relacionadas a:
- Linguagens de programação (Go, Python, JavaScript, Java, C++, Rust, etc.)
- Desenvolvimento de software e engenharia de software
- Algoritmos e estruturas de dados
- Boas práticas de código e design patterns
- Debugging e troubleshooting técnico
- Frameworks e bibliotecas de programação
- Ferramentas de desenvolvimento (Git, Docker, CI/CD, etc.)
- Bancos de dados e SQL
- APIs e desenvolvimento web/mobile

Para perguntas fora deste escopo, responda educadamente:
"Desculpe, eu só posso ajudar com questões relacionadas a programação e desenvolvimento de software."`

type Client struct {
	genaiClient *genai.Client
}

func NewClient(ctx context.Context, apiKey string) (*Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}

	return &Client{
		genaiClient: client,
	}, nil
}

func (c *Client) Close() error {
	return c.genaiClient.Close()
}

func (c *Client) GenerateResponse(ctx context.Context, userPrompt string) (string, error) {
	model := c.genaiClient.GenerativeModel("gemini-2.5-flash")

	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(systemPrompt)},
	}

	resp, err := model.GenerateContent(ctx, genai.Text(userPrompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no response candidates returned")
	}

	if len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content parts in response")
	}

	text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "", fmt.Errorf("response part is not text")
	}

	return string(text), nil
}
