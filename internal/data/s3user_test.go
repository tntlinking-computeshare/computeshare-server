package data

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/samber/lo"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func getData() *Data {

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", 1,
		"service.name", "test",
		"service.version", "Testv1",
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	dataConf := &conf.Data{
		Database: &conf.Data_Database{
			Driver: "mysql",
			Source: "root:Aline123456@tcp(computeshare.newtouch.com:31252)/computeshare?charset=utf8&parseTime=true",
		},
		Redis: &conf.Data_Redis{
			Addr: "127.0.0.1:6379",
		},
	}
	data, _, err := NewData(dataConf, logger)
	if err != nil {
		return nil
	}

	return data
}

func getLogger() log.Logger {
	return log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", 1,
		"service.name", "test",
		"service.version", "Testv1",
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
}
func TestS3UserCreate(t *testing.T) {
	data := getData()
	userId, _ := uuid.Parse("17559def-c514-41fd-b044-b59b81f2fb6b")
	ctx := context.Background()

	s3user, err := data.db.S3User.Create().SetFkUserID(userId).SetAccessKey("any").SetSecretKey("any").Save(ctx)

	if err != nil {
		panic(err)
	}

	s3Bucket, err := data.db.S3Bucket.Create().SetBucket("some-bucket1").SetS3User(s3user).Save(ctx)
	if err != nil {
		return
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(s3Bucket.ID)

}

func TestS3UserGet(t *testing.T) {
	data := getData()
	ctx := context.Background()
	id, _ := uuid.Parse("e0c86b96-4f90-4d57-b9da-1347f4d2d5c9")
	s3User, err := data.db.S3User.Get(ctx, id)
	if err != nil {
		return
	}

	bucketList, err := s3User.QueryBuckets().All(ctx)
	if err != nil {
		return
	}
	buckets := lo.Reduce(bucketList, func(agg string, item *ent.S3Bucket, index int) string {
		if agg == "" {
			return item.Bucket
		}
		return strings.Join([]string{agg, item.Bucket}, ",")
	}, "")
	fmt.Println(buckets)
}

func TestS3User_Flush(t *testing.T) {

	data := getData()
	ctx := context.Background()
	id, _ := uuid.Parse("e0c86b96-4f90-4d57-b9da-1347f4d2d5c9")
	s3User, err := data.db.S3User.Get(ctx, id)
	bucketList, err := s3User.QueryBuckets().All(ctx)
	if err != nil {
		return
	}
	buckets := lo.Reduce(bucketList, func(agg string, item *ent.S3Bucket, index int) string {
		if agg == "" {
			return item.Bucket
		}
		return strings.Join([]string{agg, item.Bucket}, ",")
	}, "")
	fmt.Println(buckets)

	command := "docker"
	args := []string{
		"exec",
		"seaweedfs-master-1",
		"sh",
		"-c",
		fmt.Sprintf(`echo "s3.configure -access_key=%s -secret_key=%s -buckets=%s -user=%s -actions=Read,Write,List,Tagging,Admin -apply " | weed shell`,
			s3User.AccessKey, s3User.SecretKey, buckets, s3User.FkUserID),
	}

	fmt.Println(strings.Join(args, " "))

	cmd := exec.Command(command, args...)

	// 将命令的输出连接到标准输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err = cmd.Run()
	if err != nil {
		fmt.Printf("命令执行失败: %v\n", err)
	}
}

func TestS3User_DockerApi(t *testing.T) {
	// 创建Docker客户端
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Printf("无法创建Docker客户端: %v\n", err)
		return
	}
	data := getData()
	ctx := context.Background()
	id, _ := uuid.Parse("e0c86b96-4f90-4d57-b9da-1347f4d2d5c9")

	s3User, err := data.db.S3User.Get(ctx, id)
	bucketList, err := s3User.QueryBuckets().All(ctx)
	if err != nil {
		return
	}
	buckets := lo.Reduce(bucketList, func(agg string, item *ent.S3Bucket, index int) string {
		if agg == "" {
			return item.Bucket
		}
		return strings.Join([]string{agg, item.Bucket}, ",")
	}, "")
	fmt.Println(buckets)

	// 构建命令
	cmd := fmt.Sprintf(`echo "s3.configure -access_key=%s -secret_key=%s -buckets=%s -user=%s -actions=Read,Write,List,Tagging,Admin -apply " | weed shell`,
		s3User.AccessKey, s3User.SecretKey, buckets, s3User.FkUserID)

	// 准备执行命令
	cmdArgs := []string{"sh", "-c", cmd}

	// 创建容器执行命令的配置
	execConfig := types.ExecConfig{
		Cmd:          cmdArgs,
		AttachStdout: true,
		AttachStderr: true,
	}

	// 创建容器执行命令的请求
	execResp, err := cli.ContainerExecCreate(context.Background(), "seaweedfs-master-1", execConfig)
	if err != nil {
		fmt.Printf("无法创建容器执行命令: %v\n", err)
		return
	}

	// 执行命令
	resp, err := cli.ContainerExecAttach(context.Background(), execResp.ID, types.ExecStartCheck{})
	if err != nil {
		fmt.Printf("无法执行容器命令: %v\n", err)
		return
	}
	defer resp.Close()

	// 输出命令的标准输出和标准错误
	_, _ = io.Copy(os.Stdout, resp.Reader)

	// 等待命令完成
	status, err := cli.ContainerExecInspect(context.Background(), execResp.ID)
	if err != nil {
		fmt.Printf("无法检查命令状态: %v\n", err)
		return
	}

	if status.ExitCode != 0 {
		fmt.Printf("命令执行失败，退出码: %d\n", status.ExitCode)
		return
	}

	fmt.Println("命令执行成功")
}

func TestS3User_Format(t *testing.T) {
	str := fmt.Sprintf(`echo "s3.configure -access_key=%s -secret_key=any -buckets=bucket1,bucket2 -user=me -actions=Read,Write,List,Tagging,Admin -apply " | weed shell`, "d")
	fmt.Println(str)
}
