package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Teslenk0/bookstore_oauth-api/src/domain/users"
	"github.com/Teslenk0/bookstore_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://api.bookstore.com:8082",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestError) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user: restclient error")
	}

	if response.StatusCode > 299 {
		apiErr, err := errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, errors.NewInternalServerError(fmt.Sprintf("invalid error interface when trying to login user %s", err))
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response: json parsing error")
	}
	return &user, nil
}