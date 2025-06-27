package raindrop

import (
	"context"
	"fmt"
	"github.com/medialo/go-raindrop/models"
)

type RaindropsService service

type ListOptions struct {
	Page       int    `json:"page,omitempty"`
	PerPage    int    `json:"perpage,omitempty"`
	Collection string `json:"collection,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Search     string `json:"search,omitempty"`
	Sort       string `json:"sort,omitempty"`
	Nested     bool   `json:"nested,omitempty"`
}

func (s *RaindropsService) List(ctx context.Context, collectionId int, opts *ListOptions) ([]*models.RaindropResponse, error) {
	out := struct {
		Items []*models.RaindropResponse `json:"items"`
	}{}
	u := fmt.Sprintf("/raindrops/%d", collectionId)
	req := s.client.newRequest("GET", u, opts)
	if opts != nil {
		if opts.Page > 0 {
			req.Param("page", fmt.Sprint(opts.Page))
		}
		if opts.PerPage > 0 {
			req.Param("perpage", fmt.Sprint(opts.PerPage))
		}
		if opts.Collection != "" {
			req.Param("collection", opts.Collection)
		}
		if opts.Tag != "" {
			req.Param("tag", opts.Tag)
		}
		if opts.Search != "" {
			req.Param("search", opts.Search)
		}
		if opts.Sort != "" {
			req.Param("sort", opts.Sort)
		}
		if opts.Nested {
			req.Param("nested", "true")
		}
	}
	if err := req.Do(ctx, &out); err != nil {
		return nil, err
	}
	return out.Items, nil
}

func (s *RaindropsService) CreateMany(ctx context.Context, raindrops []*models.RaindropCreate) ([]*models.RaindropResponse, error) {
	if len(raindrops) > 100 {
		return nil, fmt.Errorf("maximum 100 raindrops can be created at once")
	}

	reqBody := struct {
		Items []*models.RaindropCreate `json:"items"`
	}{
		Items: raindrops,
	}

	out := struct {
		Items []*models.RaindropResponse `json:"items"`
	}{}

	req := s.client.newRequest("POST", "/raindrops", reqBody)
	if err := req.Do(ctx, &out); err != nil {
		return nil, err
	}

	return out.Items, nil
}

type UpdateManyOptions struct {
	Ids        []int    `json:"ids,omitempty"`
	Important  *bool    `json:"important,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Media      []string `json:"media,omitempty"`
	Cover      string   `json:"cover,omitempty"`
	Collection *struct {
		ID int `json:"$id"`
	} `json:"collection,omitempty"`
}

func (s *RaindropsService) UpdateMany(ctx context.Context, collectionId int, opts *UpdateManyOptions) ([]*models.RaindropResponse, error) {
	out := struct {
		Items []*models.RaindropResponse `json:"items"`
	}{}

	u := fmt.Sprintf("/raindrops/%d", collectionId)
	req := s.client.newRequest("PUT", u, opts)

	if err := req.Do(ctx, &out); err != nil {
		return nil, err
	}

	return out.Items, nil
}

type RemoveManyOptions struct {
	Search string `json:"search,omitempty"`
	Ids    []int  `json:"ids,omitempty"`
	Nested bool   `json:"nested,omitempty"`
}

func (s *RaindropsService) RemoveMany(ctx context.Context, collectionId int, opts *RemoveManyOptions) (int, error) {
	out := struct {
		Result   bool `json:"result"`
		Modified int  `json:"modified"`
	}{}

	u := fmt.Sprintf("/raindrops/%d", collectionId)
	req := s.client.newRequest("DELETE", u, opts)

	if err := req.Do(ctx, &out); err != nil {
		return 0, err
	}

	return out.Modified, nil
}
