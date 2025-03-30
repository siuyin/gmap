package auth

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenPrivateKey(t *testing.T) {
	k := privateKey()
	der, err := x509.MarshalPKCS8PrivateKey(k)
	if err != nil {
		t.Fatal(err)
	}

	b := pem.EncodeToMemory(&pem.Block{Type: "ECDSA PRIVATE KEY", Bytes: der})

	blk, rest := pem.Decode(b)
	if len(rest) > 0 {
		t.Fatal("rest is not empty")
	}
	if blk.Type != "ECDSA PRIVATE KEY" {
		t.Fatal("wrong type")
	}
	if !bytes.Equal(blk.Bytes, der) {
		t.Error("wrong bytes")
	}

	m, err := x509.ParsePKCS8PrivateKey(blk.Bytes)
	if err != nil {
		t.Fatal(err)
	}

	if !m.(*ecdsa.PrivateKey).PublicKey.Equal(k.Public()) {
		t.Errorf("keys are not equal")
	}
}

func TestJWTGenAndParse(t *testing.T) {
	k := privateKey()
	tk := Auth("siuyin")
	s, err := tk.SignedString(k)
	if err != nil {
		t.Error(err)
	}

	tk2, err := jwt.Parse(s, func(tk *jwt.Token) (interface{}, error) {
		return k.Public(), nil
	}, jwt.WithValidMethods([]string{"ES256"}))
	if !tk2.Valid {
		t.Error("invalid token")
	}
	if sub, _ := tk2.Claims.GetSubject(); sub != "siuyin" {
		t.Errorf("wrong subject: %s", sub)
	}
	aud, _ := tk2.Claims.GetAudience()
	if v := aud[0]; v != "123456" {
		t.Errorf("wrong audience: %s", v)
	}
}
