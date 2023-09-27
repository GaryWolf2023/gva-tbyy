package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/minio/minio-go/v7"
	"golang.org/x/net/context"
	"log"
	"os"
	"path/filepath"
)

func SaveFileToMinio(bucketName string, filePath string, fileName string) error {
	ctx := context.Background()
	// 创建这个bucket
	err := global.MINIO.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
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

	// 需要上传文件的基本信息
	contentType := "multipart/form-data"
	fPath := filepath.Join(filePath, fileName)
	fileInfo, err := os.Stat(fPath)
	if err == os.ErrNotExist {
		log.Printf("%s目标文件不存在", fPath)
	}
	f, err := os.Open(fPath)
	if err != nil {
		return err
	}
	uploadInfo, err := global.MINIO.PutObject(ctx, bucketName, fileName, f, fileInfo.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Printf("Successfully uploaded %s of size %d\n", fileName, uploadInfo.Size)
	return nil
}
