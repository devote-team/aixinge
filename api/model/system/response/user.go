package response

import (
	"aixinge/api/model/system"
)

type UserResponse struct {
	User system.User `json:"user"`
}

type LoginResponse struct {
	User         system.User `json:"user"`
	Token        string      `json:"token"`
	RefreshToken string      `json:"refreshToken"`
}
