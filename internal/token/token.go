package token

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/pkg/jwt"
	jwtgo "github.com/dgrijalva/jwt-go"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func generateJWT(exp time.Duration, secret string) (string, error) {
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	claims := jwtgo.MapClaims{
		"id":  id,
		"iat": time.Now().Unix(),
		"exp": exp,
	}

	return jwt.Generate(claims, secret)
}
