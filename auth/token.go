package auth

import (
	"os"
	"strconv"

	schema "github.com/dezhab-service/data"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type AccessTokenClaims struct {
	UserID string `json:"uid"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	UserID string `json:"uid"`
	jwt.StandardClaims
}

func GetAccessToken(tokenDetails *schema.AccessTokenDetails) (string, error) {
	signingKey := []byte(os.Getenv("TOKEN_KEY"))

	now := jwt.TimeFunc().Unix()
	accessTokenAliveTime, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_ALIVE_TIME"), 10, 64)

	if err != nil {
		return "", err
	}
	expireTime := now + accessTokenAliveTime

	claims := AccessTokenClaims{
		UserID: tokenDetails.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    tokenDetails.Issuer,
			IssuedAt:  now,
			NotBefore: now,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	return ss, err
}

func GetRefreshToken(tokenDetails *schema.RefreshTokenDetails) (string, error) {
	signingKey := []byte(os.Getenv("TOKEN_KEY"))

	now := jwt.TimeFunc().Unix()
	accessTokenAliveTime, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_ALIVE_TIME"), 10, 64)
	if err != nil {
		return "", err
	}
	expireTime := now + accessTokenAliveTime

	claims := RefreshTokenClaims{
		UserID: tokenDetails.UserId,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    tokenDetails.Issuer,
			IssuedAt:  now,
			NotBefore: now,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	return ss, err
}

func VerifyAccessToken(tokenString string) (*AccessTokenClaims, error) {
	tokenKey := os.Getenv("TOKEN_KEY")
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Println("i was in verify access token")
	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Error(err)
		return nil, err
	}
}

func VerifyRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	tokenKey := os.Getenv("TOKEN_KEY")
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshTokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Error(err)
		return nil, err
	}
}
