package request

import (
	"aixinge/utils/snowflake"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type TokenClaims struct {
	UUID     uuid.UUID
	ID       snowflake.ID
	Username string
	NickName string
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	ID snowflake.ID
	jwt.RegisteredClaims
}
