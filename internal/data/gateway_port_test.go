package data

import (
	"context"
	"github.com/google/uuid"
	"testing"
)

func TestGatewayPortRepo_Create(t *testing.T) {
	data := getData()
	fkGatewayId, _ := uuid.Parse("db3bec89-2eea-4ae0-a567-1b5f88d334f2")
	ctx := context.Background()
	for i := 11000; i < 15000; i++ {
		data.db.GatewayPort.Create().SetFkGatewayID(fkGatewayId).SetPort(int32(i)).SetIsPublic(false).SetIsUse(false).ExecX(ctx)
	}
}
