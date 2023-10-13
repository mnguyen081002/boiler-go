package service

import (
	"erp/api_errors"
	config "erp/config"
	"erp/constants"
	"erp/domain"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type jwtService struct {
	config *config.Config
	logger *zap.Logger
}

func NewJwtService(config *config.Config, logger *zap.Logger) domain.JwtService {
	return &jwtService{
		config: config,
		logger: logger,
	}
}

func (j *jwtService) GenerateToken(userID string, kid string, tokenType constants.TokenType, expiresIn int64) (string, error) {

	claims := domain.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "erp",
		},
		TokenType: string(tokenType),
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenObj.Header["kid"] = kid
	key, _ := os.ReadFile("./config/keys/private_rsa")

	fmt.Println("key", key)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return "", errors.Wrap(err, "Error parsing private key")
	}
	token, err := tokenObj.SignedString(privateKey)
	if err != nil {
		return "", errors.Wrap(err, "Error generating token")
	}

	return token, nil
}

func (j *jwtService) GenerateAuthTokens(userID string) (string, string, error) {

	j.logger.Debug("Generating auth tokens", zap.Any("ExpiresIn", j.config.Jwt.AccessTokenExpiresIn))

	accessToken, err := j.GenerateToken(userID, j.config.Jwt.Kid, constants.AccessToken, j.config.Jwt.AccessTokenExpiresIn)
	if err != nil {
		return "", "", errors.Wrap(err, "Error generating access token")
	}

	refreshToken, err := j.GenerateToken(userID, j.config.Jwt.Kid, constants.RefreshToken, j.config.Jwt.RefreshTokenExpiresIn)
	if err != nil {
		return "", "", errors.Wrap(err, "Error generating refreshToken token")
	}

	return accessToken, refreshToken, nil
}

func (j *jwtService) ValidateToken(token string, tokenType constants.TokenType) (*string, error) {
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New(api_errors.ErrTokenExpired)
	}

	if claims.Issuer != "erp" {
		return nil, errors.New(api_errors.ErrTokenInvalid)
	}

	if claims.Subject == "" {
		return nil, errors.New(api_errors.ErrTokenInvalid)
	}

	return &claims.Subject, nil
}
