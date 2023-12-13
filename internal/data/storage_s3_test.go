package data

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tj/assert"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	ctx := context.Background()
	data := getData()
	repo := NewS3UserRepo(data, getLogger())
	userId, _ := uuid.Parse("17559def-c514-41fd-b044-b59b81f2fb6b")
	s3User, err := repo.GetS3User(ctx, userId)
	assert.NoError(t, err)
	bucket, err := repo.CreateBucket(ctx, s3User, "123")
	assert.NoError(t, err)
	fmt.Println(bucket)
}
