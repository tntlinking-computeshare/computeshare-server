package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/tj/assert"
	"os"
	"testing"
	"time"
)

func GetDomainBindingUseCase() *DomainBindingUseCase {

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", 1,
		"service.name", "test",
		"service.version", "Testv1",
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	uc, err := NewDomainBindingUseCase(nil, logger)
	if err != nil {
		panic(err)
	}
	return uc
}

func TestDomainBindingUseCase_CreateDomainBinding(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	uc := GetDomainBindingUseCase()
	err := uc.CreateDomainBinding(ctx)

	assert.NoError(t, err)
}
