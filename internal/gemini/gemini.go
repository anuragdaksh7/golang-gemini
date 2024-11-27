package gemini

import (
	"context"
	"github.com/google/generative-ai-go/genai"
)

type GeminiResponseReq struct {
	Query string `json:"query"`
}

type GeminiResponseRes struct {
	Message []genai.Part `json:"message"`
}

type Service interface {
	GetResponse(c context.Context, req *GeminiResponseReq) (*GeminiResponseRes, error)
}
