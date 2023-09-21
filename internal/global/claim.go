package global

import (
	"context"
	kratosJWT "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

const (
	ExpiresTime = time.Hour * 24
)

type ComputeServerClaim struct {
	UserID string
	jwt.RegisteredClaims
}

func (c *ComputeServerClaim) GetUserId() uuid.UUID {
	id, err := uuid.Parse(c.UserID)
	if err != nil {
		return uuid.Nil
	}
	return id
}

// FromContext 从上下文获取用户信息
func FromContext(ctx context.Context) (*ComputeServerClaim, bool) {
	return getToken(ctx)
}

func getToken(ctx context.Context) (token *ComputeServerClaim, ok bool) {
	claim, ok := kratosJWT.FromContext(ctx)
	if !ok {
		return &ComputeServerClaim{}, ok
	}
	token, ok = claim.(*ComputeServerClaim)
	return
}

func getTokenMock(_ context.Context) (token *ComputeServerClaim, ok bool) {
	return &ComputeServerClaim{
		UserID: "a3546e51-8976-44ea-94a3-1c74ebda0118",
	}, true
}
