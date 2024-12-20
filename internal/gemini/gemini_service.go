package gemini

import (
	"context"
	"gemini-go/util"
	"time"
)

type service struct {
	timeout time.Duration
}

func NewService() Service {
	return &service{
		time.Duration(20) * time.Second,
	}
}

func (s *service) GetResponse(c context.Context, req *GeminiResponseReq) (*GeminiResponseRes, error) {
	println(req.Query)
	only, err := util.GenerateContentTextOnly(req.Query)
	if err != nil {
		return nil, err
	}

	res := &GeminiResponseRes{
		Message: only,
	}

	return res, nil
}

func (s *service) GetImageResponse(c context.Context, req *GeminiImageQueryReq) (*GeminiImageQueryRes, error) {
	println(req.Query)

	only, err := util.GenerateContentImagePrompt(req.File, req.FileType, req.Query)
	if err != nil {
		return nil, err
	}

	res := &GeminiImageQueryRes{
		Message: only,
	}
	return res, nil
}
