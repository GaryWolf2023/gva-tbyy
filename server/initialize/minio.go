package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinioClient() *minio.Client {
	// 基本的配置信息
	endpoint := "192.168.0.2:49153"
	accessKeyID := "admin123"
	secretAccessKey := "pass123456"

	// 初始化一个minio客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		fmt.Printf("初始化MinioClient错误：%s", err.Error())
		global.GVA_LOG.Error("初始化MinioClient错误")
		return nil
	}
	fmt.Println("minio启动成功")
	return minioClient
}
