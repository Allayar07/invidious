package service

import "invidious/internal/repository"

type Service struct {
	DB *DbService
}

func NewService(repo *repository.VideoRepository) *Service {
	return &Service{
		DB: NewDbService(repo),
	}
}
