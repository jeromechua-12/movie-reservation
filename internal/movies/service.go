package movies

import (
	"context"
)

type Service struct {
	repo *Repository
}

func newService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (svc *Service) createMovie(ctx context.Context, movie Movie) (Movie, error) {
	movie, err := svc.repo.insertMovie(ctx, movie)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}
