package blog

import (
	"context"
	"github.com/rs/zerolog/log"
)

type Service struct {
	repo Repository
}

func ProvideService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetEntry(ctx context.Context, slug string) (*Entry, error) {
	post, err := s.repo.FindOne(ctx, slug)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Could not load entry from repository")
		return nil, ErrRepository
	}
	return post, nil
}

func (s Service) ListEntries(ctx context.Context) (*[]Entry, error) {
	var entries *[]Entry
	entries, err := s.repo.Find(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Could not load entries from repository")
		return nil, ErrRepository
	}
	return entries, nil
}
