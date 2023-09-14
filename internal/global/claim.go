package global

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	ExpiresTime = time.Hour * 24
)

type ComputeServerClaim struct {
	UserID string
	jwt.RegisteredClaims
}

//func FromContext(ctx context.Context) (token *ComputeServerClaim, ok bool) {
//	claim, ok := kratosJWT.FromContext(ctx)
//	if !ok {
//		return &ComputeServerClaim{}, ok
//	}
//	token, ok = claim.(*ComputeServerClaim)
//	return
//}

func FromContext(_ context.Context) (token *ComputeServerClaim, ok bool) {
	return &ComputeServerClaim{
		UserID: "a3546e51-8976-44ea-94a3-1c74ebda0118",
	}, true
}
