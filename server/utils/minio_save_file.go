package utils

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/minio/minio-go/v7"
	"golang.org/x/net/context"
	"log"
)

type MinioOperate struct{}

func (m *MinioOperate) SaveFileToMinio(bucketName string, filePath string, fileName string) (object, err error) {
	fmt.Println(bucketName, filePath, fileName)
	ctx := context.Background()
	// 创建这个bucket
	err1 := global.MINIO.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err1 != nil {
		// 检测这个bucket是否已经存在
		exists, errBucketExists := global.MINIO.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	// 上传文件
	contentType := "application/html"
	uploadInfo, err2 := global.MINIO.FPutObject(ctx, bucketName, filePath, fileName, minio.PutObjectOptions{ContentType: contentType})
	fmt.Println("----------------", uploadInfo)
	if err2 != nil {
		log.Println("上传文件失败", err2)
		return nil, err2
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", fileName, uploadInfo.Size)
	return object, nil
}

func (m *MinioOperate) GetFileFromMinio(bucketName string, filePath string, fileName string) {
	ctx := context.Background()
	err := global.MINIO.FGetObject(ctx, bucketName, fileName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Println("下载错误: ", err)
	}
}
