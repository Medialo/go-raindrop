package raindrop

import (
	"context"
	"fmt"
	"github.com/Medialo/raindrop-go/models"
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

func (s *RaindropService) Create(ctx context.Context, raindrop *models.RaindropRequest) (*models.RaindropResponse, error) {
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

func (s *RaindropService) Update(ctx context.Context, id int, raindrop *models.RaindropResponse) (*models.RaindropResponse, error) {
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

//func (s *RaindropService) UploadFile(ctx context.Context, file io.Reader, filename string, collectionId int) (*models.RaindropResponse, error) {
//	out := struct {
//		Result bool             `json:"result"`
//		Item   *models.RaindropResponse `json:"item"`
//	}{}
//	req := s.client.NewFileRequest("PUT", "/raindrop/file", file, filename, collectionId)
//	if err := req.Do(ctx, &out); err != nil {
//		return nil, err
//	}
//	return out.Item, nil
//}
//
//func (s *RaindropService) UploadCover(ctx context.Context, id int, cover io.Reader, filename string) (*models.RaindropResponse, error) {
//	out := struct {
//		Result bool             `json:"result"`
//		Item   *models.RaindropResponse `json:"item"`
//	}{}
//	req := s.client.NewFileRequest("PUT", fmt.Sprintf("/raindrop/%d/cover", id), cover, filename, 0)
//	if err := req.Do(ctx, &out); err != nil {
//		return nil, err
//	}
//	return out.Item, nil
//}
//
//func (s *RaindropService) GetCache(ctx context.Context, id int) (string, error) {
//	req := s.client.newRequest("GET", fmt.Sprintf("/raindrop/%d/cache", id), nil)
//	location, err := req.DoGetLocation(ctx)
//	if err != nil {
//		return "", err
//	}
//	return location, nil
//}
//
//func (s *RaindropService) Suggest(ctx context.Context, link string) (*models.SuggestResult, error) {
//	out := struct {
//		Result bool                  `json:"result"`
//		Item   *models.SuggestResult `json:"item"`
//	}{}
//	req := s.client.newRequest("POST", "/raindrop/suggest", map[string]string{"link": link})
//	if err := req.Do(ctx, &out); err != nil {
//		return nil, err
//	}
//	return out.Item, nil
//}
//
//func (s *RaindropService) SuggestById(ctx context.Context, id int) (*models.SuggestResult, error) {
//	out := struct {
//		Result bool                  `json:"result"`
//		Item   *models.SuggestResult `json:"item"`
//	}{}
//	req := s.client.newRequest("GET", fmt.Sprintf("/raindrop/%d/suggest", id), nil)
//	if err := req.Do(ctx, &out); err != nil {
//		return nil, err
//	}
//	return out.Item, nil
//}
