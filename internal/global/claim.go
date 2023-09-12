package global

import (
	"context"
	kratosJWT "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	ExpiresTime = time.Hour * 24
)

type ComputeServerClaim struct {
	ID string
	jwt.RegisteredClaims
}

func FromContext(ctx context.Context) (token *ComputeServerClaim, ok bool) {
	claim, ok := kratosJWT.FromContext(ctx)
	if !ok {
		return &ComputeServerClaim{}, ok
	}
	token, ok = claim.(*ComputeServerClaim)
	return
}
