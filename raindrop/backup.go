package raindrop

import (
	"context"
	"fmt"
	"github.com/Medialo/raindrop-go/models"
)

//https://api.raindrop.io/v1/backup

type BackupService service

func (s *BackupService) Create(ctx context.Context) (bool, error) {
	out := struct {
		Result  bool   `json:"result"`
		Message string `json:"message"`
	}{}
	req := s.client.newRequest("GET", "/backup", nil)
	if err := req.Do(ctx, &out); err != nil {
		return false, err
	}
	return true, nil
}

func (s *BackupService) List(ctx context.Context) (*models.Backups, error) {
	req := s.client.newRequest("GET", "/backups", nil)
	out := &models.Backups{}
	if err := req.Do(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *BackupService) DownloadAsCsv(ctx context.Context, id string, path string) error {
	err := s.client.DownloadToFile(ctx, fmt.Sprintf("/backup/%s.csv", id), path)
	return err
}

func (s *BackupService) DownloadAHtml(ctx context.Context, id string, path string) error {
	err := s.client.DownloadToFile(ctx, fmt.Sprintf("/backup/%s.html", id), path)
	return err
}

func (s *BackupService) DownloadLastBackupAsCsv(ctx context.Context, path string) error {
	backups, err := s.List(ctx)
	if err != nil {
		return err
	}
	if len(backups.Items) > 0 {
		return s.DownloadAsCsv(ctx, backups.Items[0].Id, path)
	}
	return fmt.Errorf("no backup found")
}

func (s *BackupService) DownloadLastBackupAsHtml(ctx context.Context, path string) error {
	backups, err := s.List(ctx)
	if err != nil {
		return err
	}
	if len(backups.Items) > 0 {
		return s.DownloadAHtml(ctx, backups.Items[0].Id, path)
	}
	return fmt.Errorf("no backup found")
}

//https://api.raindrop.io/v1/backup/65d29bfbd9429ff5db9a6274.html
