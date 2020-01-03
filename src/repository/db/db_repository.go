package db

import (
	"fmt"
	"github.com/Teslenk0/bookstore_oauth-api/src/clients/cassandra"
	"github.com/Teslenk0/bookstore_oauth-api/src/domain/access_token"
	"github.com/Teslenk0/bookstore_oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken           = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken        = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryUpdateExpiresAccessToken = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(token access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestError
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	session := cassandra.GetSession()

	if session.Closed() {
		return nil, errors.NewInternalServerError("internal server error: session has been closed")
	}

	defer session.Close()

	var result access_token.AccessToken

	if err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {

		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}

		return nil, errors.NewInternalServerError(fmt.Sprintf("internal server error:%s ", err.Error()))
	}

	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestError {
	session := cassandra.GetSession()
	if session.Closed() {
		return errors.NewInternalServerError("internal server error: session has been closed")
	}

	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestError {
	session := cassandra.GetSession()
	if session.Closed() {
		return errors.NewInternalServerError("internal server error: session has been closed")
	}

	defer session.Close()

	if err := session.Query(queryUpdateExpiresAccessToken,
		token.Expires,
		token.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
