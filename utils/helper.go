package utils

import (
	"os"

	"github.com/sajanjswl/auth/auth"
	"github.com/sajanjswl/auth/config"
	schema "github.com/sajanjswl/auth/data"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAccessTokenAndRefreshtoken(email string) (string, string, error) {

	refreshTokenDetails := &schema.RefreshTokenDetails{
		UserId: email,
		Issuer: os.Getenv("TOKEN_ISSUER"),
	}

	refreshToken, err := auth.GetRefreshToken(refreshTokenDetails)
	if err != nil {
		log.Error(err)
		return "", "", status.Errorf(codes.Internal, config.InternalErrorLogin)
	}

	// // //generating access token
	accessTokenDetails := &schema.AccessTokenDetails{
		UserId: email,

		Issuer: refreshTokenDetails.Issuer,
	}

	accessToken, err := auth.GetAccessToken(accessTokenDetails)
	if err != nil {
		log.Error(err)
		return "", "", status.Errorf(codes.Internal, config.InternalErrorLogin)
	}

	return accessToken, refreshToken, nil

}
