package internal

import (
	"os"
	"time"

	"aidanwoods.dev/go-paseto"
)

// GeneratePasetoToken generates a PASETO token for the given user ID
// and returns the token and public key
func GeneratePasetoToken(userID string) (string, error) {
	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(24 * time.Hour * 30)) // 30 days
	token.SetIssuer("quddus")

	token.SetString("user-id", userID)

	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(os.Getenv("PASETO_SECRET_KEY"))
	if err != nil {
		return "", err
	}

	signed := token.V4Sign(secretKey, nil)

	return signed, nil
}

// DecodePasetoToken decodes the given PASETO token and returns the user ID
func DecodePasetoToken(token string) (string, error) {
	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(os.Getenv("PASETO_PUBLIC_KEY")) // this wil fail if given key in an invalid format
	if err != nil {
		return "", err
	}

	parser := paseto.NewParser()              // only used because this example token has expired, use NewParser() (which checks expiry by default)
	parser.AddRule(paseto.NotExpired())       // this will fail if the token has expired
	parser.AddRule(paseto.IssuedBy("quddus")) // this will fail if the token was not issued by "quddus"

	decodedToken, err := parser.ParseV4Public(publicKey, token, nil) // this will fail if parsing failes, cryptographic checks fail, or validation rules fail
	if err != nil {
		return "", err
	}

	userID, err := decodedToken.GetString("user-id")
	if err != nil {
		return "", err
	}

	return userID, nil
}
