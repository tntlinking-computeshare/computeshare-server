package biz

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"testing"
)

func TestUpload(t *testing.T) {
	// 配置您的SeaweedFS-S3信息
	accessKey := "someKey"
	secretKey := "someSecret"
	region := "your-region"
	endpoint := "http://127.0.0.1:8333" // 直接使用 IP 地址w
	bucket := "mohaijiang"

	// 配置AWS S3客户端
	config := &aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	}

	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	s3Client := s3.New(sess)

	_, err = s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return
	}

	// 示例：列出存储桶中的对象
	resp, err := s3Client.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	if err != nil {
		panic(err)
	}

	fmt.Println("Objects in bucket:")
	for _, item := range resp.Contents {
		fmt.Println(*item.Key)
	}
}
