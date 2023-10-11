package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sakupay-apps/config"
	"github.com/sakupay-apps/internal/model"
)

func CreateAccessToken(user *model.User) (string, error) {
	now := time.Now().UTC()
	end := now.Add(config.Cfg.TokenConfig.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Cfg.TokenConfig.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(config.Cfg.TokenConfig.JWTSigningMethod, claims)
	ss, err := token.SignedString(config.Cfg.TokenConfig.JWTSignatureKey)

	if err != nil {
		return "", fmt.Errorf("failed to create access token : %s", err.Error())
	}

	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != config.Cfg.TokenConfig.JWTSigningMethod {
			return nil, fmt.Errorf("invalid token string method")
		}
		return config.Cfg.TokenConfig.JWTSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid parse token : %s", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid || claims["iss"] != config.Cfg.TokenConfig.ApplicationName {
		return nil, fmt.Errorf("invalid token mapclaims")
	}

	return claims, nil
}
