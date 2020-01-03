package access_token

import (
	"github.com/Teslenk0/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}
//Web frontend - Client-Id: 123
//Android APP - Client-Id: 234


func (at *AccessToken) Validate() *errors.RestError{
	accessTokenId := strings.TrimSpace(at.AccessToken)
	if accessTokenId == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid access token user id")
	}

	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid access token client id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid access token expiration time")
	}


	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
