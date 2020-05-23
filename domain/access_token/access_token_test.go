package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstents(t *testing.T) {
	assert.EqualValues(t, 24, EXPIRATION_TIME, "expiration time should be 24")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be exoired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expireing 3 hours from now should not be exoired")
}
