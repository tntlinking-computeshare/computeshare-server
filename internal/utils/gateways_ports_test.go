package utils

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"testing"
)

func TestAddGatewaysPorts(t *testing.T) {
	dsnStr := "user:password@tcp(host:port)/<database>?parseTime=True"
	client, err := ent.Open("mysql", dsnStr)
	if err != nil {
		fmt.Errorf("mysql 连接失败 %s", err.Error())
		panic(t)
	}
	defer client.Close()
	fkGatewayID := "db3b9c89-2eea-4ee0-a567-1f5f88d334f2"
	fkGatewayIDUUId, err := uuid.Parse(fkGatewayID)
	for i := 40001; i <= 50000; i++ {
		gatewaysPort, err := CreateGatewaysPort(context.Background(), client, fkGatewayIDUUId, int32(i), t)
		if err != nil {
			fmt.Errorf("CreateGatewaysPort 失败 %s", err.Error())
			panic(t)
		}
		fmt.Println(gatewaysPort)
	}
}

func CreateGatewaysPort(ctx context.Context, client *ent.Client, fkGatewayID uuid.UUID, port int32, t *testing.T) (*ent.GatewayPort, error) {
	isPublic := false
	if port > 41500 {
		isPublic = true
	}
	u, err := client.GatewayPort.
		Create().
		SetFkGatewayID(fkGatewayID).SetPort(port).SetIsPublic(isPublic).SetIsUse(false).Save(ctx)
	if err != nil {
		fmt.Errorf("failed creating GatewaysPort: %s", err.Error())
		panic(t)
	}
	return u, nil
}
