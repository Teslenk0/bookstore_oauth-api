package services

import (
	"github.com/Teslenk0/bookstore_oauth-api/src/domain/access_token"
	"github.com/Teslenk0/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(token access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestError
}

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(token access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestError
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestError) {

	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(token access_token.AccessToken) *errors.RestError{
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.Create(token)
}

func (s *service) UpdateExpirationTime(token access_token.AccessToken) *errors.RestError{
	if err := token.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(token)
}