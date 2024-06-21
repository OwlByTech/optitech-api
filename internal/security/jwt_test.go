package security

import (
	"encoding/json"
	"optitech/internal/dto"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const Secret = "secret"

// You must provide the token type struct
type UserToken struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func TestSignToken(t *testing.T) {
	user := &UserToken{
		UserId: 1,
	}

	token, err := JWTSign(user, Secret)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// Verify return a map, so the trick is marshal and unmarshal to create a
	// new struct you should also validate the jwt
	payloadClaims, err := JWTVerify(token, Secret)
	assert.Nil(t, err)

	var userVerified UserToken

	// TODO: remove the used of json for more performance
	bytes, err := json.Marshal(payloadClaims.Claims)
	assert.Nil(t, err)

	err = json.Unmarshal(bytes, &userVerified)
	assert.Nil(t, err)

	err = dto.ValidateDTO(userVerified)
	assert.Nil(t, err)

	assert.Equal(t, user.UserId, userVerified.UserId)
}
