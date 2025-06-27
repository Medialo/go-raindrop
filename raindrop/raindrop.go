package raindrop

import (
	"context"
	"fmt"
	"github.com/medialo/go-raindrop/models"
)

type RaindropService service

func (s *RaindropService) Get(ctx context.Context, id int) (*models.RaindropResponse, error) {
	out := struct {
		Item *models.RaindropResponse `json:"item"`
	}{}
	req := s.client.newRequest("GET", fmt.Sprintf("/raindrop/%d", id), nil)
	if err := req.Do(ctx, &out); err != nil {
		return nil, err
	}
	return out.Item, nil
}

func (s *RaindropService) Create(ctx context.Context, raindrop *models.RaindropCreate) (*models.RaindropResponse, error) {
	out := struct {
		Result bool                     `json:"result"`
		Item   *models.RaindropResponse `json:"item"`
	}{}
	req := s.client.newRequest("POST", "/raindrop", raindrop)
	if err := req.Do(ctx, &out); err != nil {
		return nil, err
	}
	return out.Item, nil
}

func (s *RaindropService) Update(ctx context.Context, id int, raindrop *models.RaindropUpdate) (*models.RaindropResponse, error) {
	out := struct {
		Result bool                     `json:"result"`
		Item   *models.RaindropResponse `json:"item"`
	}{}
	req := s.client.newRequest("PUT", fmt.Sprintf("/raindrop/%d", id), raindrop)
	if err := req.Do(ctx, &out); err != nil {
		return nil, err
	}
	return out.Item, nil
}

func (s *RaindropService) Remove(ctx context.Context, id int) error {
	out := struct {
		Result bool `json:"result"`
	}{}
	req := s.client.newRequest("DELETE", fmt.Sprintf("/raindrop/%d", id), nil)
	if err := req.Do(ctx, &out); err != nil {
		return err
	}
	return nil
}
