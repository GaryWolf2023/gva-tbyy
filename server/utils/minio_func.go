package utils

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	testpayloadone "github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne"
	"github.com/minio/minio-go/v7"
	"golang.org/x/net/context"
	"io"
	"log"
)

type MinioOperate struct{}

func (m *MinioOperate) SaveFileToMinio(bucketName string, objectName string, fileName string, file io.Reader, fileSize int64) (object testpayloadone.MinioSaveInfo, err error) {
	ctx := context.Background()
	fmt.Println(bucketName, objectName, fileName)
	var minioObject testpayloadone.MinioSaveInfo
	bucketInfos, err := global.MINIO.ListBuckets(ctx)
	if err != nil {
		fmt.Println("List Buckets err：", err.Error())
		return minioObject, err
	}
	for index, bucketInfo := range bucketInfos {
		fmt.Printf("List Bucket No {%d}----filename{%s}-----createTime{%s}\n", index+1, bucketInfo.Name, bucketInfo.CreationDate.Format("2006-01-02 15:04:05"))
	}
	// 检测是否有了bucket
	isExist, err0 := global.MINIO.BucketExists(context.Background(), bucketName)
	fmt.Println(isExist)
	if err0 != nil {
		fmt.Printf("Check %s err：%s\n", bucketName, err0.Error())
		global.GVA_LOG.Error(fmt.Sprintf("%s", err0.Error()))
		return minioObject, err0
	}
	if isExist {
		// 检测到有
		fmt.Printf("We already own %s\n", bucketName)
	} else {
		// 检测到没有，然后创建
		err1 := global.MINIO.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true})
		// 创建失败
		if err1 != nil {
			fmt.Println("创建BucketName失败")
			log.Printf("Failed created %s\n", bucketName)
			global.GVA_LOG.Error("创建")
			return minioObject, err1
		}
		//创建成功
		log.Printf("Successfully created %s\n", bucketName)
	}
	// 上传文件
	contentType := "multipart/form-data"
	fmt.Println(bucketName, objectName)
	fmt.Println(file)
	uploadInfo, err2 := global.MINIO.PutObject(ctx, bucketName, objectName, file, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err2 != nil {
		log.Println("上传文件失败", err2)
		return minioObject, err2
	}
	fmt.Printf("%s\n%s\n%s", uploadInfo, uploadInfo.Bucket, uploadInfo.ETag, uploadInfo.Key)
	minioObject.Bucket = uploadInfo.Bucket
	minioObject.Key = uploadInfo.Key
	log.Println("Successfully uploaded %s of size %d\n", fileName, uploadInfo.Size)
	return minioObject, nil
}

// RemoveFile
// 删除文件
func (m *MinioOperate) RemoveFile(BucketName string, objectName string) (error error) {
	opts := minio.RemoveObjectOptions{}
	err := global.MINIO.RemoveObject(context.Background(), BucketName, objectName, opts)
	if err != nil {
		global.GVA_LOG.Error("删除文件失败")
		return err
	}
	return nil
}

// 暂时没写完，还不知道怎么写后面再看，暂时没有功能需要这个功能
func (m *MinioOperate) RemoveFiles(bucketName string, searchData string) (err error) {
	objectsCh := make(chan minio.ObjectInfo)
	// 注意一般不要自己来构造，直接选择从bucket中查询，查询到的对象放入objectsCh
	for object := range global.MINIO.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Prefix:    searchData,
		Recursive: true,
	}) {
		if object.Key == "头像.jpg" {
			objectsCh <- object
		}
	}
	defer close(objectsCh)
	// 删除
	for rErr := range global.MINIO.RemoveObjects(context.Background(), bucketName, objectsCh, minio.RemoveObjectsOptions{}) {
		fmt.Println("Delete err:", rErr.Err.Error())
	}
	return nil
}

// 获取object
func (m *MinioOperate) GetFileFromMinio(bucketName string, objectName string) (*minio.Object, error) {
	ctx := context.Background()
	data, err := global.MINIO.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	fmt.Println(data)
	if err != nil {
		global.GVA_LOG.Error("获取失败")
		return nil, err
	}
	return data, nil
}
