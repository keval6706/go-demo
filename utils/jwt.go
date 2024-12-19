package utils

import (
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	// privateKey *rsa.PrivateKey
	// publicKey  *rsa.PublicKey
	privateKey, publicKey = jwt.MustLoadRSA("private.pem", "public.pem")

	signer   = jwt.NewSigner(jwt.RS256, privateKey, 15*time.Minute)
	verifier = jwt.NewVerifier(jwt.RS256, publicKey)
)

type TokenClaims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AuthHeaders struct {
	Authorization string `header:"Authorization,required"`
}

// GENERATE TOKEN
func GenerateToken(payload TokenClaims) ([]byte, error) {

	token, err := signer.Sign(payload)

	return token, err
}

// VERIFY THE TOKEN
func VerifyToken(token string) (TokenClaims, error) {

	verified, err := verifier.VerifyToken([]byte(token))
	var payload TokenClaims
	verified.Claims(&payload)

	return payload, err
}
