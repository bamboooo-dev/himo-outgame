package grpcmiddleware

import (
	"context"
	"os"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type key int

// StringKey は Context に 値を埋め込んだ際の key
const StringKey key = iota

// AuthClaim は JWT に埋め込む Claim
type AuthClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Authenticate はメタデータに埋め込まれた JWT を Parse して認証認可を行う関数
func Authenticate(ctx context.Context) (context.Context, error) {
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	claim := AuthClaim{}
	_, err = jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, StringKey, claim.UserID), nil
}
