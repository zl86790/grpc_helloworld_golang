package services

import (
	"context"
	"fmt"
)

type HiService struct {
}

func (hs HiService) GetHiResponed(ctx context.Context, request *HiRequest) (*HiResponse, error) {
	fmt.Println("Say:" + request.Say)
	return &HiResponse{Responed: "helloworld"}, nil
}
