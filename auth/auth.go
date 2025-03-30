package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(subj string) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss": "id.beyondbroadcast.com",
			"sub": subj,
			"aud": "123456", // some client ID
		})
}

func privateKey() *ecdsa.PrivateKey {
	k, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	return k
}
