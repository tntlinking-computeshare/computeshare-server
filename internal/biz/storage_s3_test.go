package biz

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
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
func TestKey(t *testing.T) {
	prefix := "d"
	var keys []string
	keys = append(keys, "a.txt")
	keys = append(keys, "b.txt")
	keys = append(keys, "c.txt")
	keys = append(keys, "d/a.txt")
	keys = append(keys, "d/d.txt")
	keys = append(keys, "w/d.txt")
	keys = append(keys, "c/d/a.txt")
	keys = append(keys, "c/d/r.txt")
	keys = append(keys, "d/d/t.txt")
	keys = append(keys, "d/d/d/d.txt")
	var file []string
	for _, key := range keys {
		if prefix == "" {
			splitN := strings.SplitN(key, "/", 2)
			if len(splitN) > 1 && !contains(file, splitN[0]+"/") {
				file = append(file, splitN[0]+"/")
			} else if len(splitN) == 1 {
				file = append(file, splitN[0])
			}
		} else {
			dir := prefix + "/"
			if key[:len(dir)] == dir {
				splitN := strings.SplitN(key[len(dir):], "/", 2)
				if len(splitN) > 1 && !contains(file, splitN[0]+"/") {
					file = append(file, splitN[0]+"/")
				} else if len(splitN) == 1 {
					file = append(file, splitN[0])
				}
			}
		}
	}
	fmt.Println(file)
}

func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
