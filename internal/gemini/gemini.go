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

type GeminiImageQueryReq struct {
	Query    string `json:"query"`
	File     []byte `json:"file"`
	FileType string `json:"file_type"`
}

type GeminiImageQueryRes struct {
	Message []genai.Part `json:"message"`
}

type Service interface {
	GetResponse(c context.Context, req *GeminiResponseReq) (*GeminiResponseRes, error)
	GetImageResponse(c context.Context, req *GeminiImageQueryReq) (*GeminiImageQueryRes, error)
}
