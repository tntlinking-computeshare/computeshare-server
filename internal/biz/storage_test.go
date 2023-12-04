package biz

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestStoragecase_S3UserCreate(t *testing.T) {

	//command := "docker exec -it seaweedfs-master-1 sh -c 'echo \"s3.configure -access_key=any -secret_key=any -buckets=bucket1 -user=me -actions=Read,Write,List,Tagging,Admin -apply \" | weed shell'"

	command := "docker"
	args := []string{
		"exec",
		"seaweedfs-master-1",
		"sh",
		"-c",
		`echo "s3.configure -access_key=any -secret_key=any -buckets=bucket1,bucket2 -user=me -actions=Read,Write,List,Tagging,Admin -apply " | weed shell`,
	}

	cmd := exec.Command(command, args...)

	// 将命令的输出连接到标准输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		fmt.Printf("命令执行失败: %v\n", err)
	}
}
