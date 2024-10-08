package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID
}

// type AuthService struct {
// 	userRepo repo.User
// 	signKey  string
// }

// func NewAuthService(userRepo repo.User, signKey string) *AuthService {
// 	return &AuthService{
// 		userRepo: userRepo,
// 		signKey:  signKey,
// 	}
// }

type AuthService struct {
	// userRepo repo.User
	signKey string
}

func NewAuthService(signKey string) *AuthService {
	return &AuthService{
		// userRepo: userRepo,
		signKey: signKey,
	}
}

func (s *AuthService) ParseToken(accessToken string) (uuid.UUID, []string, error) {
	publicKey, err := parseKeycloakRSAPublicKey(s.signKey)
	if err != nil {
		panic(err)
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return the public key that is used to validate the token.
		return publicKey, nil
	})
	if err != nil {
		return uuid.Nil, nil, err
	}

	if !token.Valid {
		return uuid.Nil, nil, fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	userId := uuid.MustParse(claims["sub"].(string))
	userRoles := claims["realm_access"].(map[string]interface{})["roles"]
	var rolesStr []string
	for _, role := range userRoles.([]interface{}) {
		rolesStr = append(rolesStr, role.(string))
	}

	return userId, rolesStr, nil
}

func parseKeycloakRSAPublicKey(base64Encoded string) (*rsa.PublicKey, error) {
	buf, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		return nil, err
	}
	parsedKey, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, err
	}
	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if ok {
		return publicKey, nil
	}
	return nil, fmt.Errorf("unexpected key type %T", publicKey)
}
